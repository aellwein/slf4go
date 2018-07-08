package slf4go

import (
	"io"
	"os"
)

type LogLevel int

// log levels defined by slf4go
const (
	LevelTrace LogLevel = iota
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelPanic
)

var (
	// the one and only global log factory per application, must be first set using SetLoggerFactory() method.
	theLoggerFactory LoggerFactory

	// A default Writer used for log output, usually log implementations use own one.
	Writer io.Writer = os.Stderr

	// all log levels as array
	AllLevels = []LogLevel{LevelTrace, LevelDebug, LevelInfo, LevelWarn, LevelError, LevelFatal, LevelPanic}
)

// Logger interface
type Logger interface {
	// Get the name of this logger.
	GetName() string

	// Setup log level
	SetLevel(l LogLevel)

	// Get log level
	GetLevel() LogLevel

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

type LoggingParameters map[string]interface{}

// factory interface, which can be implemented by an adaptor.
type LoggerFactory interface {
	// This method is called in order to get a logger instance.
	GetLogger(name string) Logger

	// A set of arbitrary parameters can be passed to underlying logger adaptor.
	// It's up to a logger adaptor implementation to validate the parameters and return appropriate error,
	// or nil, if no error was occurred.
	SetLoggingParameters(params LoggingParameters) error

	// Sets default log level for all loggers created by this factory.
	SetDefaultLogLevel(level LogLevel)
}

type LoggerAdaptor struct {
	name  string
	level LogLevel
}

func (a *LoggerAdaptor) SetName(name string) {
	a.name = name
}

func (a *LoggerAdaptor) GetName() string {
	return a.name
}

func (a *LoggerAdaptor) GetLevel() LogLevel {
	return a.level
}

func (a *LoggerAdaptor) SetLevel(l LogLevel) {
	a.level = l
}

func (a *LoggerAdaptor) IsTraceEnabled() bool {
	return a.level <= LevelTrace
}

func (a *LoggerAdaptor) IsDebugEnabled() bool {
	return a.level <= LevelDebug
}

func (a *LoggerAdaptor) IsInfoEnabled() bool {
	return a.level <= LevelInfo
}

func (a *LoggerAdaptor) IsWarnEnabled() bool {
	return a.level <= LevelWarn
}

func (a *LoggerAdaptor) IsErrorEnabled() bool {
	return a.level <= LevelError
}

func (a *LoggerAdaptor) IsFatalEnabled() bool {
	return a.level <= LevelFatal
}

func (a *LoggerAdaptor) IsPanicEnabled() bool {
	return a.level <= LevelPanic
}

// The LoggerFactory is the interface meant to be provided by slf4go implementors.
func SetLoggerFactory(factory LoggerFactory) {
	if factory == nil {
		panic("LoggerFactory can't be nil")
	}
	theLoggerFactory = factory
}

// get logger from our logger factory. It will panic, if no logger factory is set yet.
func GetLogger(name string) Logger {
	if theLoggerFactory == nil {
		panic("LoggerFactory was not set! Please ensure you use SetLoggerFactory() before you get a logger instance")
	}
	return theLoggerFactory.GetLogger(name)
}

// get logger factory
func GetLoggerFactory() LoggerFactory {
	return theLoggerFactory
}

// Stringify LogLevel
func (l LogLevel) String() string {
	switch l {
	case LevelTrace:
		return "TRACE"
	case LevelDebug:
		return "DEBUG"
	case LevelInfo:
		return "INFO"
	case LevelWarn:
		return "WARN"
	case LevelError:
		return "ERROR"
	case LevelFatal:
		return "FATAL"
	case LevelPanic:
		return "PANIC"
	}
	panic("no match for log level")
}
