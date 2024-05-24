package logging

import (
	"os"

	"github.com/rs/zerolog"
)

var _ Logger = (*ZLogger)(nil)

type ZLogger struct {
	*zerolog.Logger
}

func NewZLogger() *ZLogger {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	return &ZLogger{
		&logger,
	}
}

func (l *ZLogger) Debug(msg string, tags ...any) {
	l.Logger.Debug().Fields(tags).Msg(msg)
}

func (l *ZLogger) Info(msg string, tags ...any) {
	l.Logger.Info().Fields(tags).Msg(msg)
}

func (l *ZLogger) Warn(msg string, tags ...any) {
	l.Logger.Warn().Fields(tags).Msg(msg)
}

func (l *ZLogger) Error(msg string, tags ...any) {
	l.Logger.Error().Fields(tags).Msg(msg)
}

func (l *ZLogger) Fatal(msg string, tags ...any) {
	l.Logger.Fatal().Fields(tags).Msg(msg)
}

func (l *ZLogger) Debugf(template string, args ...interface{}) {
	l.Logger.Debug().Msgf(template, args...)
}

func (l *ZLogger) Infof(template string, args ...interface{}) {
	l.Logger.Info().Msgf(template, args...)
}

func (l *ZLogger) Warnf(template string, args ...interface{}) {
	l.Logger.Warn().Msgf(template, args...)
}

func (l *ZLogger) Errorf(template string, args ...interface{}) {
	l.Logger.Error().Msgf(template, args...)
}

func (l *ZLogger) Fatalf(template string, args ...interface{}) {
	l.Logger.Fatal().Msgf(template, args...)
}

func (l *ZLogger) With(tags ...any) Logger {
	newLogger := l.Logger.With().Fields(tags).Logger()
	return &ZLogger{
		Logger: &newLogger,
	}
}
