package logger

type LoggerManager interface {
	Info(msg string, keysAndValues ...interface{})
	Error(msg string, keysAndValues ...interface{})
	Debug(msg string, keysAndValues ...interface{})
	Warn(msg string, keysAndValues ...interface{})
}

type CombinedLogger struct {
	consoleLogger LoggerManager
	fileLogger    LoggerManager
}

func NewCombinedLogger(consoleLogger LoggerManager, fileLogger LoggerManager) *CombinedLogger {
	return &CombinedLogger{
		consoleLogger: consoleLogger,
		fileLogger:    fileLogger,
	}
}

func (l *CombinedLogger) log(level LogLevel, msg string, keysAndValues ...interface{}) {
	switch level {
	case Debug:
		l.consoleLogger.Debug(msg, keysAndValues...)
		l.fileLogger.Debug(msg, keysAndValues...)
	case Info:
		l.consoleLogger.Info(msg, keysAndValues...)
		l.fileLogger.Info(msg, keysAndValues...)
	case Warn:
		l.consoleLogger.Warn(msg, keysAndValues...)
		l.fileLogger.Warn(msg, keysAndValues...)
	case Error:
		l.consoleLogger.Error(msg, keysAndValues...)
		l.fileLogger.Error(msg, keysAndValues...)
	}
}

func (l *CombinedLogger) Info(msg string, keysAndValues ...interface{}) {
	l.log(Info, msg, keysAndValues...)
}

func (l *CombinedLogger) Error(msg string, keysAndValues ...interface{}) {
	l.log(Error, msg, keysAndValues...)
}

func (l *CombinedLogger) Debug(msg string, keysAndValues ...interface{}) {
	l.log(Debug, msg, keysAndValues...)
}

func (l *CombinedLogger) Warn(msg string, keysAndValues ...interface{}) {
	l.log(Warn, msg, keysAndValues...)
}
