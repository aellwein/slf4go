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
