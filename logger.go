package slf4go

import (
	"io"
	"os"
)

// Loglevel as used by slf4go
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

	// Writer is a default Writer used for log output, usually log implementations use own one.
	Writer io.Writer = os.Stderr

	// AllLevels contain all log levels as array
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

	// IsTraceEnabled returns true, if level TRACE is enabled.
	IsTraceEnabled() bool

	// IsDebugEnabled returns true, if level DEBUG is enabled.
	IsDebugEnabled() bool

	// IsInfoEnabled returns true, if level INFO is enabled.
	IsInfoEnabled() bool

	// IsWarnEnabled returns true, if level WARN is enabled.
	IsWarnEnabled() bool

	// IsErrorEnabled returns true, if level ERROR is enabled.
	IsErrorEnabled() bool

	// IsFatalEnabled returns true, if level FATAL is enabled.
	IsFatalEnabled() bool

	// IsPanicEnabled returns true, if level PANIC is enabled.
	IsPanicEnabled() bool

	// Trace logs a message on log level TRACE
	Trace(args ...interface{})

	// Tracef logs a formatted message on log level TRACE
	Tracef(format string, args ...interface{})

	// Debug logs a message on log level DEBUG
	Debug(args ...interface{})

	// Tracef logs a formatted message on log level TRACE
	Debugf(format string, args ...interface{})

	// Info logs a message on log level INFO
	Info(args ...interface{})

	// Infof logs a formatted message on log level INFO
	Infof(format string, args ...interface{})

	// Warn logs a message on log level WARN
	Warn(args ...interface{})

	// Warnf logs a formatted message on log level WARN
	Warnf(format string, args ...interface{})

	// Error logs a message on log level ERROR
	Error(args ...interface{})

	// Errorf logs a formatted message on log level ERROR
	Errorf(format string, args ...interface{})

	// Fatal logs a message on log level FATAL. Program exits afterwards, no defer functions are called.
	Fatal(args ...interface{})

	// Fatalf logs a formatted message on log level FATAL. Program exits afterwards, no defer functions are called.
	Fatalf(format string, args ...interface{})

	// Panic logs a message on log level PANIC. Panic is caused (can be recovered), defer functions are called.
	Panic(args ...interface{})

	// Panicf logs a formatted message on log level PANIC. Panic is caused (can be recovered), defer functions are called.
	Panicf(format string, args ...interface{})
}

// LoggingParameters is a map, which contains parameters passed to the logging adaptor implementation.
type LoggingParameters map[string]interface{}

// LoggerFactory is the factory interface, which is to be implemented by an adaptor.
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

// LoggerAdaptor pre-implements some functions of Logger
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
