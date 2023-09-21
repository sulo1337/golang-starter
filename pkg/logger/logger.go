package logger

type Logger interface {
	GetLevel() int

	Infof(string, ...interface{})
	Noticef(string, ...interface{})
	Errorf(string, ...interface{})
	Warnf(string, ...interface{})
	Debugf(string, ...interface{})
	Panicf(string, ...interface{})
	Fatalf(string, ...interface{})

	Info(...interface{})
	Notice(...interface{})
	Error(...interface{})
	Warn(...interface{})
	Debug(...interface{})
	Panic(...interface{})
	Fatal(...interface{})
}

type LoggingOptions struct {
	Level      int
	ShowTime   bool
	ShowCaller bool
}

var (
	Fatal  = 4
	Panic  = 3
	Error  = 2
	Warn   = 1
	Info   = 0
	Notice = 0
	Debug  = -1
)
