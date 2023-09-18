package logger

import (
	"encoding/json"
	"fmt"
	"os"
)

type JSONParser struct{}

func NewJSONParser() *JSONParser {
	return &JSONParser{}
}

func (j *JSONParser) Parse(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("JSONParser parse file failed: %s", err)
	}
	defer file.Close()

	var loggersFromJSONFile LoggersFromJSONFile
	if err = json.NewDecoder(file).Decode(&loggersFromJSONFile); err != nil {
		return fmt.Errorf("JSONParser decode file failed: %s", err)
	}

	root := loggersFromJSONFile.Loggers
	rootLogger := NewRootLogger(
		WithLevelThreshold(generateLevel(root.LevelThreshold)),
		WithExporter(root.Exporter.generateExporter()),
		WithLayout(generateLayout(root.Layout)),
	)

	appGame := root.AppGame
	gameLogger, err := NewNormalLogger(
		"app.game",
		rootLogger,
		WithLevelThreshold(generateLevel(appGame.LevelThreshold)),
		WithExporter(appGame.Exporter.generateExporter()),
	)
	if err != nil {
		return fmt.Errorf("JSONParser new game Logger failed: %s", err)
	}

	appGameAi := appGame.AppGameAi
	aiLogger, err := NewNormalLogger(
		"app.game.ai",
		gameLogger,
		WithLevelThreshold(generateLevel(appGameAi.LevelThreshold)),
	)
	if err != nil {
		return fmt.Errorf("JSONParser new ai Logger failed: %s", err)
	}

	DeclareLoggers(rootLogger, gameLogger, aiLogger)
	return nil
}

type LoggersFromJSONFile struct {
	Loggers struct {
		LoggerOptionsForJSONParser
		AppGame struct {
			LoggerOptionsForJSONParser
			AppGameAi struct {
				LoggerOptionsForJSONParser
			} `json:"app.game.ai"`
		} `json:"app.game"`
	} `json:"loggers"`
}

type LoggerOptionsForJSONParser struct {
	LevelThreshold string                `json:"levelThreshold"`
	Exporter       ExporterForJSONParser `json:"exporter"`
	Layout         string                `json:"layout"`
}

type ExporterForJSONParser struct {
	Type     string                  `json:"type"`
	FileName string                  `json:"fileName"`
	Children []ExporterForJSONParser `json:"children"`
}

func (e ExporterForJSONParser) generateExporter() Exporter {
	if len(e.Children) == 0 {
		switch e.Type {
		case "console":
			return NewConsoleExporter()
		case "file":
			return NewFileExporter(e.FileName)
		default:
			fmt.Printf("type %s is NOT supported\n", e.Type)
			return nil
		}
	}

	exporters := make([]Exporter, 0)
	for _, child := range e.Children {
		exporters = append(exporters, child.generateExporter())
	}

	return NewCompositeExporter(exporters...)
}

func generateLevel(stringOfLevel string) Level {
	switch stringOfLevel {
	case TRACE.String():
		return TRACE
	case INFO.String():
		return INFO
	case DEBUG.String():
		return DEBUG
	case WARN.String():
		return WARN
	case ERROR.String():
		return ERROR
	default:
		fmt.Printf("string of level %s is NOT supported\n", stringOfLevel)
		return Level(-1)
	}
}

func generateLayout(nameOfLayout string) Layout {
	switch nameOfLayout {
	case "standard":
		return NewStandardLayout()
	default:
		fmt.Printf("name of layout %s is NOT supported\n", nameOfLayout)
		return nil
	}
}
