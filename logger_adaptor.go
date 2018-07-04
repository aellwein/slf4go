package slf4go

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
	return a.level <= LEVEL_TRACE
}

func (a *LoggerAdaptor) IsDebugEnabled() bool {
	return a.level <= LEVEL_DEBUG
}

func (a *LoggerAdaptor) IsInfoEnabled() bool {
	return a.level <= LEVEL_INFO
}

func (a *LoggerAdaptor) IsWarnEnabled() bool {
	return a.level <= LEVEL_WARN
}

func (a *LoggerAdaptor) IsErrorEnabled() bool {
	return a.level <= LEVEL_ERROR
}

func (a *LoggerAdaptor) IsFatalEnabled() bool {
	return a.level <= LEVEL_FATAL
}

func (a *LoggerAdaptor) IsPanicEnabled() bool {
	return a.level <= LEVEL_PANIC
}
