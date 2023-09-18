package logger

import (
	"errors"
	"fmt"
)

var (
	NameOfRoot    = "Root"
	GlobalLoggers = make([]*Logger, 0)
)

type Logger struct {
	name     string
	parent   *Logger
	children []*Logger
	LoggerOptions
}

type LoggerOptions struct {
	levelThreshold Level
	exporter       Exporter
	layout         Layout
}

type SetLoggerOptionsFunc func(loggerOptions *LoggerOptions)

func NewRootLogger(setLoggerOptionsFuncs ...SetLoggerOptionsFunc) *Logger {
	loggerOptions := defaultLoggerOptions()
	for _, setLoggerOptionsFunc := range setLoggerOptionsFuncs {
		setLoggerOptionsFunc(&loggerOptions)
	}

	return &Logger{
		name:          NameOfRoot,
		parent:        nil,
		children:      make([]*Logger, 0),
		LoggerOptions: loggerOptions,
	}
}

func NewNormalLogger(name string, parent *Logger, setLoggerOptionsFuncs ...SetLoggerOptionsFunc) (*Logger, error) {
	for _, logger := range GlobalLoggers {
		if logger.name == name {
			return nil, errors.New("normal logger should have unique name")
		}
	}

	if parent == nil {
		return nil, errors.New("normal logger should have parent logger")
	}

	// 如果 Logger.LoggerOptions 改用 pointer 型態的話
	// 所有 GlobalLoggers 的 LoggerOptions 都會指向同一個記憶體位址
	// 顯然 GlobalLoggers 應該要各自擁有不同的 LoggerOptions
	// 所以 Logger.LoggerOptions 不採用 pointer 型態
	loggerOptions := parent.LoggerOptions
	for _, setLoggerOptionsFunc := range setLoggerOptionsFuncs {
		setLoggerOptionsFunc(&loggerOptions)
	}

	normalLogger := &Logger{
		name:          name,
		parent:        parent,
		children:      make([]*Logger, 0),
		LoggerOptions: loggerOptions,
	}
	parent.AddChild(normalLogger)

	return normalLogger, nil
}

func (l *Logger) Trace(message Message) {
	l.templateForHandlingMessage(TRACE, message)
}

func (l *Logger) Info(message Message) {
	l.templateForHandlingMessage(INFO, message)
}

func (l *Logger) Debug(message Message) {
	l.templateForHandlingMessage(DEBUG, message)
}

func (l *Logger) Warn(message Message) {
	l.templateForHandlingMessage(WARN, message)
}

func (l *Logger) Error(message Message) {
	l.templateForHandlingMessage(ERROR, message)
}

func (l *Logger) templateForHandlingMessage(levelOfMessage Level, message Message) {
	if levelOfMessage.IsMoreThanOrEqualTo(l.levelThreshold) {
		formatedMessage := l.layout.ModifyMessageFormat(levelOfMessage, l.name, message)
		l.exporter.Export(formatedMessage)
	}
}

func (l *Logger) AddChild(child *Logger) {
	l.children = append(l.children, child)
}

func defaultLoggerOptions() LoggerOptions {
	return LoggerOptions{
		levelThreshold: DEBUG,
		exporter:       NewConsoleExporter(),
		layout:         NewStandardLayout(),
	}
}

func WithLevelThreshold(levelThreshold Level) SetLoggerOptionsFunc {
	return func(loggerOptions *LoggerOptions) {
		loggerOptions.levelThreshold = levelThreshold
	}
}

func WithExporter(exporter Exporter) SetLoggerOptionsFunc {
	return func(loggerOptions *LoggerOptions) {
		loggerOptions.exporter = exporter
	}
}

func WithLayout(layout Layout) SetLoggerOptionsFunc {
	return func(loggerOptions *LoggerOptions) {
		loggerOptions.layout = layout
	}
}

func DeclareLoggers(loggers ...*Logger) {
	GlobalLoggers = append(GlobalLoggers, loggers...)
}

func GetLogger(name string) (*Logger, error) {
	for _, logger := range GlobalLoggers {
		if logger.name == name {
			return logger, nil
		}
	}

	return nil, fmt.Errorf("could NOT find logger %s in GlobalLoggers", name)
}
