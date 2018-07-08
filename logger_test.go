package slf4go

import (
	"errors"
	"fmt"
	"github.com/smartystreets/assertions"
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

type mockLoggerFactory struct {
	defaultLogLevel LogLevel
}

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

func (m *mockLoggerFactory) SetDefaultLogLevel(level LogLevel) {
	m.defaultLogLevel = level
}

func (m *mockLoggerFactory) SetLoggingParameters(params LoggingParameters) error {
	if _, ok := params["test"]; ok {
		return errors.New("provoked error")
	}
	return nil
}

func newMockLoggerFactory() mockLoggerFactory {
	m := mockLoggerFactory{}
	m.SetDefaultLogLevel(LevelTrace)
	return m
}

type level struct {
	checkFunc func() bool
	funcName  string
}

func TestLoggerAdaptor_AllLevels(t *testing.T) {
	Convey("Given a new logger factory", t, func() {

		factory := newMockLoggerFactory()
		SetLoggerFactory(&factory)

		for _, lvl := range AllLevels {
			Convey(fmt.Sprintf("When logger with level %v is created", lvl), func() {

				logger := GetLogger("test")
				logger.SetLevel(lvl)
				l := map[LogLevel]level{
					LevelTrace: {
						checkFunc: logger.IsTraceEnabled,
						funcName:  "IsTraceEnabled",
					},
					LevelDebug: {
						checkFunc: logger.IsDebugEnabled,
						funcName:  "IsDebugEnabled",
					},
					LevelInfo: {
						checkFunc: logger.IsInfoEnabled,
						funcName:  "IsInfoEnabled",
					},
					LevelWarn: {
						checkFunc: logger.IsWarnEnabled,
						funcName:  "IsWarnEnabled",
					},
					LevelError: {
						checkFunc: logger.IsErrorEnabled,
						funcName:  "IsErrorEnabled",
					},
					LevelFatal: {
						checkFunc: logger.IsFatalEnabled,
						funcName:  "IsFatalEnabled",
					},
					LevelPanic: {
						checkFunc: logger.IsPanicEnabled,
						funcName:  "IsPanicEnabled",
					},
				}

				for k, v := range l {
					var mustBe = k >= lvl
					Convey(fmt.Sprintf("Then %s() must be '%v'", v.funcName, mustBe), func() {
						if mustBe {
							So(v.checkFunc(), ShouldBeTrue)
						} else {
							So(v.checkFunc(), ShouldBeFalse)
						}
					})
				}
			})
		}
	})
}

func TestLoggerAdaptor_GetName(t *testing.T) {
	Convey("Given a new logger factory", t, func() {
		factory := new(mockLoggerFactory)
		SetLoggerFactory(factory)

		Convey("When a new named logger 'test' is created", func() {
			logger := GetLogger("test")

			Convey("Then GetName() returns 'test'", func() {
				So(logger.GetName(), ShouldEqual, "test")
			})
		})
	})
}

func TestLoggerAdaptor_GetLevel(t *testing.T) {
	Convey("Given a new logger factory", t, func() {
		factory := new(mockLoggerFactory)
		SetLoggerFactory(factory)

		Convey("When a new named logger with level 'Info' is created", func() {
			logger := GetLogger("test")
			logger.SetLevel(LevelInfo)

			Convey("Then GetLevel() returns 'Info'", func() {
				So(logger.GetLevel(), ShouldEqual, LevelInfo)
			})
		})
	})
}

func TestLoggerFactory_SetLoggingParameters(t *testing.T) {
	Convey("Given a new logger factory", t, func() {
		factory := new(mockLoggerFactory)
		SetLoggerFactory(factory)

		Convey("When SetLoggingParameters() with 'test' parameter is called", func() {
			err := GetLoggerFactory().SetLoggingParameters(LoggingParameters{"test": "bla"})

			Convey("Then an error is returned", func() {
				So(err.Error(), ShouldEqual, "provoked error")
			})
		})
	})
}

func TestUnknownLogLevel_String_ShouldPanic(t *testing.T) {
	assertions.ShouldPanic(func() { fmt.Print(LogLevel(42).String()) })
}
