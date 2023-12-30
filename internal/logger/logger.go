package logger

import (
	"dunlap/internal/database"
	"fmt"
	"os"
	"strings"
	"time"
)

var globalLogger = NewLogger(INFO, &ConsoleLogOutput{}, time.RFC3339, true)

func (c *ConsoleLogOutput) WriteDatabaseEntry(message string) error {
	return nil
}

func (c *ConsoleLogOutput) Close() error {
	return nil
}

func NewCompositeLogOutput(outputs ...LogOutput) *CompositeLogOutput {
	return &CompositeLogOutput{outputs: outputs}
}

func (c *CompositeLogOutput) WriteDatabaseEntry(message string) error {
	c.mu.RLock()
	defer c.mu.RUnlock()

	for _, output := range c.outputs {
		if err := output.WriteDatabaseEntry(message); err != nil {
			return err
		}
	}
	return nil
}

func (c *CompositeLogOutput) Close() error {
	var err error
	for _, output := range c.outputs {
		if e := output.Close(); e != nil {
			err = e
		}
	}
	return err
}

func NewLogger(level Level, output LogOutput, timestampFormat string, printLogs bool) *Logger {
	return &Logger{
		output:          output,
		level:           level,
		timestampFormat: timestampFormat,
		printLogs:       printLogs,
	}
}

func (l *Logger) log(level Level, format string, v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if level < l.level {
		return
	}

	functioninfo, line := GetFunctionNameAndLine()
	functionName := strings.Replace(functioninfo, ".", " -> ", -1)
	lineLocation := fmt.Sprintf("line %d", line)

	message := fmt.Sprintf(format, v...)

	timestamp := time.Now().Format(l.timestampFormat)

	levelColor, ok := levelColors[level]

	if !ok {
		levelColor = colorReset
	}

	logEntry := fmt.Sprintf("%s[%s]%s | %s%s%s | %s @ %s | %s\n",
		timeColor, timestamp, colorReset, levelColor, level.String(),
		colorReset, functionName, lineLocation, message)

	if l.printLogs {
		fmt.Print(logEntry)
	}

	if l.output != nil {
		if err := l.output.WriteDatabaseEntry(logEntry); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to write log: %v\n", err)
		}
	}

	if level == PANIC {
		panic(fmt.Sprintf(format, v...))
	}
}

func InitializeLogger(config LoggerConfig) {
	consoleOutput := &ConsoleLogOutput{}
	globalLogger = NewLogger(config.LogLevel, consoleOutput, time.RFC3339, config.PrintLogs)

	if config.MongoURI != "" && config.DatabaseName != "" && config.CollectionName != "" {
		mongoDBOutput, err := database.NewMongoDBLogOutput(config.MongoURI, config.DatabaseName, config.CollectionName)
		if err != nil {
			fmt.Println("Error connecting to database for logs:", err)
			return
		}

		compositeOutput := NewCompositeLogOutput(consoleOutput, mongoDBOutput)
		globalLogger.output = compositeOutput
	}
}
