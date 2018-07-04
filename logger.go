package slf4go

type LogLevel int

// log levels defined by slf4go
const (
	LEVEL_TRACE LogLevel = iota
	LEVEL_DEBUG
	LEVEL_INFO
	LEVEL_WARN
	LEVEL_ERROR
	LEVEL_FATAL
	LEVEL_PANIC
)

// Logger interface
type Logger interface {
	// Get the name of l, which was used for `GetLogger`
	GetName() string

	// Setup l's level.
	SetLevel(l LogLevel)

	IsTraceEnabled() bool
	IsDebugEnabled() bool
	IsInfoEnabled() bool
	IsWarnEnabled() bool
	IsErrorEnabled() bool
	IsFatalEnabled() bool
	IsPanicEnabled() bool

	Trace(args ...interface{})
	Tracef(format string, args ...interface{})
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Panic(args ...interface{})
	Panicf(format string, args ...interface{})
}

// factory interface, which can be implemented by an adaptor.
type LoggerFactory interface {
	GetLogger(name string) Logger
}
