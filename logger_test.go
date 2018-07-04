package slf4go

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestPanicIsCausedWhenCallingGetLoggerWithoutSetLoggerFactory(t *testing.T) {
	Convey("Given we didn't called SetLoggerFactory() before", t, func() {
		Convey("When calling GetLogger()", func() {
			Convey("Then a panic is caused", func() {
				defer func() {
					if r := recover(); r == nil {
						t.Error("The expected panic was not caused!")
					}
				}()
				GetLogger("test")
			})
		})
	})
}

func TestSetLoggerFactoryDoesNotAcceptNilArgument(t *testing.T) {
	Convey("Given we are about to initialize the slf4go", t, func() {
		Convey("When we are calling SetLoggerFactory with nil argument", func() {
			Convey("Then a panic is caused", func() {
				defer func() {
					if r := recover(); r == nil {
						t.Error("The expected panic is not caused!")
					}
				}()
				SetLoggerFactory(nil)
			})
		})
	})
}

// now test some log adapter functionality by using a mock.
// we implement only what we need to fulfill the interfaces.

type mockLoggerFactory struct{}

type mockLogger struct {
	LoggerAdaptor
}

func (*mockLogger) Trace(args ...interface{}) {
}

func (*mockLogger) Tracef(format string, args ...interface{}) {
}

func (*mockLogger) Debug(args ...interface{}) {
}

func (*mockLogger) Debugf(format string, args ...interface{}) {
}

func (*mockLogger) Info(args ...interface{}) {
}

func (*mockLogger) Infof(format string, args ...interface{}) {
}

func (*mockLogger) Warn(args ...interface{}) {
}

func (*mockLogger) Warnf(format string, args ...interface{}) {
}

func (*mockLogger) Error(args ...interface{}) {
}

func (*mockLogger) Errorf(format string, args ...interface{}) {
}

func (*mockLogger) Fatal(args ...interface{}) {
}

func (*mockLogger) Fatalf(format string, args ...interface{}) {
}

func (*mockLogger) Panic(args ...interface{}) {
}

func (*mockLogger) Panicf(format string, args ...interface{}) {
}

func (m *mockLoggerFactory) GetLogger(name string) Logger {
	l := new(mockLogger)
	l.SetName(name)
	l.SetLevel(LevelDebug)
	return l
}

func TestLoggerAdaptor_TRACE(t *testing.T) {
	Convey("Given a new logger factory", t, func() {

		factory := new(mockLoggerFactory)
		SetLoggerFactory(factory)

		Convey("When logger with level TRACE is created", func() {

			logger := GetLogger("test")
			logger.SetLevel(LevelTrace)

			Convey("Then IsTraceEnabled() must be 'true'", func() {
				So(logger.IsTraceEnabled(), ShouldBeTrue)
			})
			Convey("Then IsDebugEnabled() must be 'true'", func() {
				So(logger.IsDebugEnabled(), ShouldBeTrue)
			})
			Convey("Then IsInfoEnabled() must be 'true'", func() {
				So(logger.IsInfoEnabled(), ShouldBeTrue)
			})
			Convey("Then IsWarnEnabled() must be 'true'", func() {
				So(logger.IsWarnEnabled(), ShouldBeTrue)
			})
			Convey("Then IsErrorEnabled() must be 'true'", func() {
				So(logger.IsErrorEnabled(), ShouldBeTrue)
			})
			Convey("Then IsFatalEnabled() must be 'true'", func() {
				So(logger.IsFatalEnabled(), ShouldBeTrue)
			})
			Convey("Then IsPanicEnabled() must be 'true'", func() {
				So(logger.IsPanicEnabled(), ShouldBeTrue)
			})
		})
	})
}
func TestLoggerAdaptor_DEBUG(t *testing.T) {
	Convey("Given a new logger factory", t, func() {

		factory := new(mockLoggerFactory)
		SetLoggerFactory(factory)

		Convey("When logger with level DEBUG is created", func() {

			logger := GetLogger("test")
			logger.SetLevel(LevelDebug)

			Convey("Then IsTraceEnabled() must be 'false'", func() {
				So(logger.IsTraceEnabled(), ShouldBeFalse)
			})
			Convey("Then IsDebugEnabled() must be 'true'", func() {
				So(logger.IsDebugEnabled(), ShouldBeTrue)
			})
			Convey("Then IsInfoEnabled() must be 'true'", func() {
				So(logger.IsInfoEnabled(), ShouldBeTrue)
			})
			Convey("Then IsWarnEnabled() must be 'true'", func() {
				So(logger.IsWarnEnabled(), ShouldBeTrue)
			})
			Convey("Then IsErrorEnabled() must be 'true'", func() {
				So(logger.IsErrorEnabled(), ShouldBeTrue)
			})
			Convey("Then IsFatalEnabled() must be 'true'", func() {
				So(logger.IsFatalEnabled(), ShouldBeTrue)
			})
			Convey("Then IsPanicEnabled() must be 'true'", func() {
				So(logger.IsPanicEnabled(), ShouldBeTrue)
			})
		})
	})
}
func TestLoggerAdaptor_INFO(t *testing.T) {
	Convey("Given a new logger factory", t, func() {

		factory := new(mockLoggerFactory)
		SetLoggerFactory(factory)

		Convey("When logger with level INFO is created", func() {

			logger := GetLogger("test")
			logger.SetLevel(LevelInfo)

			Convey("Then IsTraceEnabled() must be 'false'", func() {
				So(logger.IsTraceEnabled(), ShouldBeFalse)
			})
			Convey("Then IsDebugEnabled() must be 'false'", func() {
				So(logger.IsDebugEnabled(), ShouldBeFalse)
			})
			Convey("Then IsInfoEnabled() must be 'true'", func() {
				So(logger.IsInfoEnabled(), ShouldBeTrue)
			})
			Convey("Then IsWarnEnabled() must be 'true'", func() {
				So(logger.IsWarnEnabled(), ShouldBeTrue)
			})
			Convey("Then IsErrorEnabled() must be 'true'", func() {
				So(logger.IsErrorEnabled(), ShouldBeTrue)
			})
			Convey("Then IsFatalEnabled() must be 'true'", func() {
				So(logger.IsFatalEnabled(), ShouldBeTrue)
			})
			Convey("Then IsPanicEnabled() must be 'true'", func() {
				So(logger.IsPanicEnabled(), ShouldBeTrue)
			})
		})
	})
}
func TestLoggerAdaptor_WARN(t *testing.T) {
	Convey("Given a new logger factory", t, func() {

		factory := new(mockLoggerFactory)
		SetLoggerFactory(factory)

		Convey("When logger with level WARN is created", func() {

			logger := GetLogger("test")
			logger.SetLevel(LevelWarn)

			Convey("Then IsTraceEnabled() must be 'false'", func() {
				So(logger.IsTraceEnabled(), ShouldBeFalse)
			})
			Convey("Then IsDebugEnabled() must be 'false'", func() {
				So(logger.IsDebugEnabled(), ShouldBeFalse)
			})
			Convey("Then IsInfoEnabled() must be 'false'", func() {
				So(logger.IsInfoEnabled(), ShouldBeFalse)
			})
			Convey("Then IsWarnEnabled() must be 'true'", func() {
				So(logger.IsWarnEnabled(), ShouldBeTrue)
			})
			Convey("Then IsErrorEnabled() must be 'true'", func() {
				So(logger.IsErrorEnabled(), ShouldBeTrue)
			})
			Convey("Then IsFatalEnabled() must be 'true'", func() {
				So(logger.IsFatalEnabled(), ShouldBeTrue)
			})
			Convey("Then IsPanicEnabled() must be 'true'", func() {
				So(logger.IsPanicEnabled(), ShouldBeTrue)
			})
		})
	})
}
func TestLoggerAdaptor_ERROR(t *testing.T) {
	Convey("Given a new logger factory", t, func() {

		factory := new(mockLoggerFactory)
		SetLoggerFactory(factory)

		Convey("When logger with level ERROR is created", func() {

			logger := GetLogger("test")
			logger.SetLevel(LevelError)

			Convey("Then IsTraceEnabled() must be 'false'", func() {
				So(logger.IsTraceEnabled(), ShouldBeFalse)
			})
			Convey("Then IsDebugEnabled() must be 'false'", func() {
				So(logger.IsDebugEnabled(), ShouldBeFalse)
			})
			Convey("Then IsInfoEnabled() must be 'false'", func() {
				So(logger.IsInfoEnabled(), ShouldBeFalse)
			})
			Convey("Then IsWarnEnabled() must be 'false'", func() {
				So(logger.IsWarnEnabled(), ShouldBeFalse)
			})
			Convey("Then IsErrorEnabled() must be 'true'", func() {
				So(logger.IsErrorEnabled(), ShouldBeTrue)
			})
			Convey("Then IsFatalEnabled() must be 'true'", func() {
				So(logger.IsFatalEnabled(), ShouldBeTrue)
			})
			Convey("Then IsPanicEnabled() must be 'true'", func() {
				So(logger.IsPanicEnabled(), ShouldBeTrue)
			})
		})
	})
}
func TestLoggerAdaptor_FATAL(t *testing.T) {
	Convey("Given a new logger factory", t, func() {

		factory := new(mockLoggerFactory)
		SetLoggerFactory(factory)

		Convey("When logger with level FATAL is created", func() {

			logger := GetLogger("test")
			logger.SetLevel(LevelFatal)

			Convey("Then IsTraceEnabled() must be 'false'", func() {
				So(logger.IsTraceEnabled(), ShouldBeFalse)
			})
			Convey("Then IsDebugEnabled() must be 'false'", func() {
				So(logger.IsDebugEnabled(), ShouldBeFalse)
			})
			Convey("Then IsInfoEnabled() must be 'false'", func() {
				So(logger.IsInfoEnabled(), ShouldBeFalse)
			})
			Convey("Then IsWarnEnabled() must be 'false'", func() {
				So(logger.IsWarnEnabled(), ShouldBeFalse)
			})
			Convey("Then IsErrorEnabled() must be 'false'", func() {
				So(logger.IsErrorEnabled(), ShouldBeFalse)
			})
			Convey("Then IsFatalEnabled() must be 'true'", func() {
				So(logger.IsFatalEnabled(), ShouldBeTrue)
			})
			Convey("Then IsPanicEnabled() must be 'true'", func() {
				So(logger.IsPanicEnabled(), ShouldBeTrue)
			})
		})
	})
}
func TestLoggerAdaptor_PANIC(t *testing.T) {
	Convey("Given a new logger factory", t, func() {

		factory := new(mockLoggerFactory)
		SetLoggerFactory(factory)

		Convey("When logger with level PANIC is created", func() {

			logger := GetLogger("test")
			logger.SetLevel(LevelPanic)

			Convey("Then IsTraceEnabled() must be 'false'", func() {
				So(logger.IsTraceEnabled(), ShouldBeFalse)
			})
			Convey("Then IsDebugEnabled() must be 'false'", func() {
				So(logger.IsDebugEnabled(), ShouldBeFalse)
			})
			Convey("Then IsInfoEnabled() must be 'false'", func() {
				So(logger.IsInfoEnabled(), ShouldBeFalse)
			})
			Convey("Then IsWarnEnabled() must be 'false'", func() {
				So(logger.IsWarnEnabled(), ShouldBeFalse)
			})
			Convey("Then IsErrorEnabled() must be 'false'", func() {
				So(logger.IsErrorEnabled(), ShouldBeFalse)
			})
			Convey("Then IsFatalEnabled() must be 'false'", func() {
				So(logger.IsFatalEnabled(), ShouldBeFalse)
			})
			Convey("Then IsPanicEnabled() must be 'true'", func() {
				So(logger.IsPanicEnabled(), ShouldBeTrue)
			})
		})
	})
}
