package slf4go

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"time"
)

const (
	l_TRACE    = "TRACE"
	l_DEBUG    = "DEBUG"
	l_INFO     = "INFO"
	l_WARN     = "WARN"
	l_ERROR    = "ERROR"
	l_FATAL    = "FATAL"
	l_PANIC    = "PANIC"
	call_depth = 2
)

//------------------------------------------------------------------------------------------------------------
// simple l that use log package
type loggerAdaptorNative struct {
	LoggerAdaptor
	tf   string
	flag int
}

// it should be private
func newNativeLogger(name string) *loggerAdaptorNative {
	logger := new(loggerAdaptorNative)
	logger.name = name
	logger.level = LEVEL_DEBUG
	logger.tf = "2006-01-02 15:04:05.999"
	logger.flag = log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile
	return logger
}

func (l *loggerAdaptorNative) Trace(args ...interface{}) {
	if l.level <= LEVEL_TRACE {
		str := fmt.Sprint(args...)
		l.output(call_depth, l_TRACE, str)
	}
}

func (l *loggerAdaptorNative) Tracef(format string, args ...interface{}) {
	if l.level <= LEVEL_TRACE {
		str := fmt.Sprintf(format, args...)
		l.output(call_depth, l_TRACE, str)
	}
}

func (l *loggerAdaptorNative) Debug(args ...interface{}) {
	if l.level <= LEVEL_DEBUG {
		str := fmt.Sprint(args...)
		l.output(call_depth, l_DEBUG, str)
	}
}

func (l *loggerAdaptorNative) Debugf(format string, args ...interface{}) {
	if l.level <= LEVEL_DEBUG {
		str := fmt.Sprintf(format, args...)
		l.output(call_depth, l_DEBUG, str)
	}
}

func (l *loggerAdaptorNative) Info(args ...interface{}) {
	if l.level <= LEVEL_INFO {
		str := fmt.Sprint(args...)
		l.output(call_depth, l_INFO, str)
	}
}

func (l *loggerAdaptorNative) Infof(format string, args ...interface{}) {
	if l.level <= LEVEL_INFO {
		str := fmt.Sprintf(format, args...)
		l.output(call_depth, l_INFO, str)
	}
}

func (l *loggerAdaptorNative) Warn(args ...interface{}) {
	if l.level <= LEVEL_WARN {
		str := fmt.Sprint(args...)
		l.output(call_depth, l_WARN, str)
	}
}

func (l *loggerAdaptorNative) Warnf(format string, args ...interface{}) {
	if l.level <= LEVEL_WARN {
		str := fmt.Sprintf(format, args...)
		l.output(call_depth, l_WARN, str)
	}
}

func (l *loggerAdaptorNative) Error(args ...interface{}) {
	if l.level <= LEVEL_ERROR {
		str := fmt.Sprint(args...)
		l.output(call_depth, l_ERROR, str)
	}
}

func (l *loggerAdaptorNative) Errorf(format string, args ...interface{}) {
	if l.level <= LEVEL_ERROR {
		str := fmt.Sprintf(format, args...)
		l.output(call_depth, l_ERROR, str)
	}
}

func (l *loggerAdaptorNative) Fatal(args ...interface{}) {
	if l.level <= LEVEL_FATAL {
		str := fmt.Sprint(args...)
		l.output(call_depth, l_FATAL, str)
		os.Exit(1)
	}
}

func (l *loggerAdaptorNative) Fatalf(format string, args ...interface{}) {
	if l.level <= LEVEL_FATAL {
		str := fmt.Sprintf(format, args...)
		l.output(call_depth, l_FATAL, str)
		os.Exit(1)
	}
}

func (l *loggerAdaptorNative) Panic(args ...interface{}) {
	str := fmt.Sprint(args...)
	l.output(call_depth, l_PANIC, str)
	panic("panic!")
}

func (l *loggerAdaptorNative) Panicf(format string, args ...interface{}) {
	str := fmt.Sprintf(format, args...)
	l.output(call_depth, l_PANIC, str)
	panic(str)
}

func (l *loggerAdaptorNative) output(calldepth int, level, s string) error {
	var file string
	var line int
	var ts = time.Now().Format(l.tf)
	if l.flag&(log.Lshortfile|log.Llongfile) != 0 {
		var ok bool
		_, file, line, ok = runtime.Caller(calldepth)
		if !ok {
			file = "???"
			line = 0
		}
		lastIndex := strings.LastIndex(file, "/")
		if lastIndex > 0 {
			file = file[lastIndex+1:]
		}
	}
	result := fmt.Sprintf("%-29s [%-5s] %s:%d %s\n", ts, level, file, line, s)
	_, err := Writer.Write([]byte(result))
	return err
}

//------------------------------------------------------------------------------------------------------------
// factory
type nativeLoggerFactory struct {
}

func newNativeLoggerFactory() LoggerFactory {
	factory := &nativeLoggerFactory{}
	return factory
}

func (factory *nativeLoggerFactory) GetLogger(name string) Logger {
	return newNativeLogger(name)
}
