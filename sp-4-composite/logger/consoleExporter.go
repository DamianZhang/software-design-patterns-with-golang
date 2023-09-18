package logger

import "fmt"

type ConsoleExporter struct{}

func NewConsoleExporter() *ConsoleExporter {
	return &ConsoleExporter{}
}

func (c *ConsoleExporter) Export(message Message) {
	fmt.Print(message)
}
