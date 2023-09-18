package structs

import (
	"encoding/json"
	"os"
)

type JSONParser struct{}

func NewJSONParser() *JSONParser {
	return &JSONParser{}
}

func (j *JSONParser) Parse(fileName string) ([]*Patient, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	patients := make([]*Patient, 0)
	if err = json.NewDecoder(file).Decode(&patients); err != nil {
		return nil, err
	}

	return patients, nil
}
