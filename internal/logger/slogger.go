package logger

import (
	"log/slog"
	"os"
)

type Slogger struct {
	*slog.Logger
}

func NewSLogger() *Slogger {
	return &Slogger{
		slog.New(slog.NewTextHandler(os.Stdout, nil)),
	}
}

func (l *Slogger) log(level LogLevel, msg string, keysAndValues ...interface{}) {
	switch level {
	case Debug:
		l.Logger.Debug(msg, keysAndValues...)
	case Info:
		l.Logger.Info(msg, keysAndValues...)
	case Warn:
		l.Logger.Warn(msg, keysAndValues...)
	case Error:
		l.Logger.Error(msg, keysAndValues...)
	}
}

func (l *Slogger) Info(msg string, keysAndValues ...interface{}) {
	l.log(Info, msg, keysAndValues...)
}

func (l *Slogger) Error(msg string, keysAndValues ...interface{}) {
	l.log(Error, msg, keysAndValues...)
}

func (l *Slogger) Debug(msg string, keysAndValues ...interface{}) {
	l.log(Debug, msg, keysAndValues...)
}

func (l *Slogger) Warn(msg string, keysAndValues ...interface{}) {
	l.log(Warn, msg, keysAndValues...)
}
