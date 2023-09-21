package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type JsonLogger struct {
	logger  *zap.SugaredLogger
	options LoggingOptions
}

func NewJsonLogger(options LoggingOptions) *JsonLogger {
	cfg := zap.Config{
		Level:       zap.NewAtomicLevelAt(zapcore.Level(options.Level)),
		Development: false,
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "ts",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	if !options.ShowTime {
		cfg.EncoderConfig.TimeKey = ""
	}

	if !options.ShowCaller {
		cfg.EncoderConfig.CallerKey = ""
	}

	logger, _ := cfg.Build()
	return &JsonLogger{logger: logger.Sugar()}
}

func (z JsonLogger) GetLevel() int {
	return z.options.Level
}

func (z JsonLogger) Infof(format string, args ...interface{}) {
	z.logger.Infof(format, args...)
}

func (z JsonLogger) Noticef(format string, args ...interface{}) {
	z.logger.Infof(format, args...)
}

func (z JsonLogger) Errorf(format string, args ...interface{}) {
	z.logger.Errorf(format, args...)
}

func (z JsonLogger) Warnf(format string, args ...interface{}) {
	z.logger.Warnf(format, args...)
}

func (z JsonLogger) Debugf(format string, args ...interface{}) {
	z.logger.Debugf(format, args...)
}

func (z JsonLogger) Panicf(format string, args ...interface{}) {
	z.logger.Panicf(format, args...)
}

func (z JsonLogger) Fatalf(format string, args ...interface{}) {
	z.logger.Fatalf(format, args...)
}

func (z JsonLogger) Info(args ...interface{}) {
	z.logger.Info(args...)
}

func (z JsonLogger) Notice(args ...interface{}) {
	z.logger.Info(args...)
}

func (z JsonLogger) Error(args ...interface{}) {
	z.logger.Error(args...)
}

func (z JsonLogger) Warn(args ...interface{}) {
	z.logger.Warn(args...)
}

func (z JsonLogger) Debug(args ...interface{}) {
	z.logger.Debug(args...)
}

func (z JsonLogger) Panic(args ...interface{}) {
	z.logger.Panic(args...)
}

func (z JsonLogger) Fatal(args ...interface{}) {
	z.logger.Fatal(args...)
}
