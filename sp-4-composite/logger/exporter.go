package logger

type Exporter interface {
	Export(message Message)
}
