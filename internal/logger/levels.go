package logger

import "os"

type Level int

const (
	DEBUG Level = iota + 1
	INFO
	WARNING
	ERROR
	FATAL
	PANIC
)

var levelStrings = []string{
	"DEBUG", "INFO", "WARNING", "ERROR", "FATAL", "PANIC",
}

func (l Level) String() string {
	if l < DEBUG || l > PANIC {
		return "UNKNOWN"
	}
	return levelStrings[l-1]
}

func Info(format string, v ...interface{}) {
	globalLogger.log(INFO, format, v...)
}

func Debug(format string, v ...interface{}) {
	globalLogger.log(DEBUG, format, v...)
}

func Warning(format string, v ...interface{}) {
	globalLogger.log(WARNING, format, v...)
}

func Error(format string, v ...interface{}) {
	globalLogger.log(ERROR, format, v...)
}

func Fatal(format string, v ...interface{}) {
	globalLogger.log(FATAL, format, v...)
	os.Exit(1)
}

func Panic(format string, v ...interface{}) {
	globalLogger.log(PANIC, format, v...)
}
