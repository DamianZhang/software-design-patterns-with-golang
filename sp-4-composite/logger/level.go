package logger

type Level int

const (
	TRACE Level = iota
	INFO
	DEBUG
	WARN
	ERROR
)

var stringsOfLevel = [5]string{"TRACE", "INFO", "DEBUG", "WARN", "ERROR"}

func (l Level) String() string {
	return stringsOfLevel[l]
}

func (l Level) IsMoreThanOrEqualTo(level Level) bool {
	return l >= level
}
