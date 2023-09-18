package logger

import (
	"bufio"
	"fmt"
	"os"
)

type FileExporter struct {
	filePath string
}

func NewFileExporter(filePath string) *FileExporter {
	return &FileExporter{
		filePath: filePath,
	}
}

func (f *FileExporter) Export(message Message) {
	file, err := os.OpenFile(f.filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("file exporter load file failed:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	writer.WriteString(string(message))
	writer.Flush()
}
