package structs

import (
	"encoding/json"
	"fmt"
	"os"
)

type JSONStorer struct{}

func NewJSONStorer() *JSONStorer {
	return &JSONStorer{}
}

func (j *JSONStorer) String() string {
	return "JSON"
}

func (j *JSONStorer) MatchFileFormat(fileFormat string) bool {
	return j.String() == fileFormat
}

func (j *JSONStorer) StorePrescriptions(prescriptions []*Prescription, fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("store JSON file format prescriptions failed:", err)
		return
	}
	defer file.Close()

	if err = json.NewEncoder(file).Encode(&prescriptions); err != nil {
		fmt.Println("store JSON file format prescriptions failed:", err)
		return
	}

	fmt.Println("store JSON file format prescriptions successfully")
}
