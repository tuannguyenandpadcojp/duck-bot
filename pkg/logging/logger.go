package logging

var defaultLogger Logger

func init() {
	defaultLogger = NewNoopLogger()
}

func SetDefaultLogger(logger Logger) {
	if logger == nil {
		return
	}
	defaultLogger = logger
}

type Logger interface {
	Debug(msg string, tags ...any)

	Info(msg string, tags ...any)

	Warn(msg string, tags ...any)

	Error(msg string, tags ...any)

	Fatal(msg string, tags ...any)

	Debugf(template string, args ...interface{})

	Infof(template string, args ...interface{})

	Warnf(template string, args ...interface{})

	Errorf(template string, args ...interface{})

	Fatalf(template string, args ...interface{})

	With(tags ...any) Logger
}

func Debug(msg string, tags ...any) {
	defaultLogger.Debug(msg, tags...)
}

func Info(msg string, tags ...any) {
	defaultLogger.Info(msg, tags...)
}

func Warn(msg string, tags ...any) {
	defaultLogger.Warn(msg, tags...)
}

func Error(msg string, tags ...any) {
	defaultLogger.Error(msg, tags...)
}

func Fatal(msg string, tags ...any) {
	defaultLogger.Fatal(msg, tags...)
}

func Debugf(template string, args ...interface{}) {
	defaultLogger.Debugf(template, args...)
}

func Infof(template string, args ...interface{}) {
	defaultLogger.Infof(template, args...)
}

func Warnf(template string, args ...interface{}) {
	defaultLogger.Warnf(template, args...)
}

func Errorf(template string, args ...interface{}) {
	defaultLogger.Errorf(template, args...)
}

func Fatalf(template string, args ...interface{}) {
	defaultLogger.Fatalf(template, args...)
}
