package logger

import "sync"

type LoggerConfig struct {
	LogLevel       Level
	PrintLogs      bool
	MongoURI       string
	DatabaseName   string
	CollectionName string
}

type LogOutput interface {
	WriteDatabaseEntry(message string) error
	Close() error
}

type ConsoleLogOutput struct{}

type CompositeLogOutput struct {
	mu      sync.RWMutex
	outputs []LogOutput
}

type Logger struct {
	mu              sync.Mutex
	output          LogOutput
	level           Level
	timestampFormat string
	printLogs       bool
}
