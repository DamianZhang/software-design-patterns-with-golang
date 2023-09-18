package structs

import (
	"encoding/csv"
	"fmt"
	"os"
)

type CSVStorer struct{}

func NewCSVStorer() *CSVStorer {
	return &CSVStorer{}
}

func (c *CSVStorer) String() string {
	return "CSV"
}

func (c *CSVStorer) MatchFileFormat(fileFormat string) bool {
	return c.String() == fileFormat
}

func (c *CSVStorer) StorePrescriptions(prescriptions []*Prescription, fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("store CSV file format prescriptions failed:", err)
		return
	}
	defer file.Close()

	stringOfPrescriptions := make([][]string, len(prescriptions))
	for i, prescription := range prescriptions {
		stringOfPrescription := make([]string, 0)
		stringOfPrescription = append(stringOfPrescription, prescription.Name, prescription.PotentialDisease, prescription.Usage)

		stringOfMedicines := ""
		for _, medicine := range prescription.Medicines {
			stringOfMedicines += medicine.Name + ","
		}
		stringOfPrescription = append(stringOfPrescription, stringOfMedicines)

		stringOfPrescriptions[i] = stringOfPrescription
	}

	csvWriter := csv.NewWriter(file)
	csvWriter.WriteAll(stringOfPrescriptions)
	csvWriter.Flush()

	fmt.Println("store CSV file format prescriptions successfully")
}
