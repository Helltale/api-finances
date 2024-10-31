package logger

import (
	"log/slog"
	"os"
	"sync"
)

type LogLevel int

const (
	Debug LogLevel = iota
	Info
	Warn
	Error
	Fatal
)

type FileLogger struct {
	*slog.Logger
	file *os.File
	mu   sync.Mutex
}

func NewFLogger(filePath string) (*FileLogger, error) {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	return &FileLogger{
		Logger: slog.New(slog.NewTextHandler(file, nil)),
		file:   file,
	}, nil
}

func (l *FileLogger) log(level LogLevel, msg string, keysAndValues ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()

	switch level {
	case Debug:
		l.Logger.Debug(msg, keysAndValues...)
	case Info:
		l.Logger.Info(msg, keysAndValues...)
	case Warn:
		l.Logger.Warn(msg, keysAndValues...)
	case Error:
		l.Logger.Error(msg, keysAndValues...)
	case Fatal:
		l.Logger.Error(msg, keysAndValues...)
		os.Exit(1)
	}
}

func (l *FileLogger) Info(msg string, keysAndValues ...interface{}) {
	l.log(Info, msg, keysAndValues...)
}

func (l *FileLogger) Error(msg string, keysAndValues ...interface{}) {
	l.log(Error, msg, keysAndValues...)
}

func (l *FileLogger) Debug(msg string, keysAndValues ...interface{}) {
	l.log(Debug, msg, keysAndValues...)
}

func (l *FileLogger) Warn(msg string, keysAndValues ...interface{}) {
	l.log(Warn, msg, keysAndValues...)
}

func (l *FileLogger) Close() {
	if l.file != nil {
		l.file.Close()
	}
}
