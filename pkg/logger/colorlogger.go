package logger

import "github.com/nats-io/nats-server/v2/logger"

type ColorLogger struct {
	logger  *logger.Logger
	options LoggingOptions
}

func NewColorLogger(options LoggingOptions) *ColorLogger {
	stdLogger := logger.NewStdLogger(
		options.ShowTime,
		options.Level == Debug,
		options.Level == Debug,
		true,
		true,
	)
	return &ColorLogger{logger: stdLogger, options: options}
}

func (n ColorLogger) GetLevel() int {
	return n.options.Level
}

func (n ColorLogger) Infof(format string, args ...interface{}) {
	n.logger.Noticef(format, args...)
}

func (n ColorLogger) Noticef(format string, args ...interface{}) {
	n.logger.Noticef(format, args...)
}

func (n ColorLogger) Errorf(format string, args ...interface{}) {
	n.logger.Errorf(format, args...)
}

func (n ColorLogger) Warnf(format string, args ...interface{}) {
	n.logger.Warnf(format, args...)
}

func (n ColorLogger) Debugf(format string, args ...interface{}) {
	n.logger.Debugf(format, args...)
}

func (n ColorLogger) Panicf(format string, args ...interface{}) {
	n.logger.Fatalf(format, args...)
}

func (n ColorLogger) Fatalf(format string, args ...interface{}) {
	n.logger.Fatalf(format, args...)
}

func (n ColorLogger) Info(args ...interface{}) {
	n.logger.Noticef("%s", args...)
}

func (n ColorLogger) Notice(args ...interface{}) {
	n.logger.Noticef("%s", args...)
}

func (n ColorLogger) Error(args ...interface{}) {
	n.logger.Errorf("%s", args...)
}

func (n ColorLogger) Warn(args ...interface{}) {
	n.logger.Warnf("%s", args...)
}

func (n ColorLogger) Debug(args ...interface{}) {
	n.logger.Debugf("%s", args...)
}

func (n ColorLogger) Panic(args ...interface{}) {
	n.logger.Fatalf("%s", args...)
}

func (n ColorLogger) Fatal(args ...interface{}) {
	n.logger.Fatalf("%s", args...)
}
