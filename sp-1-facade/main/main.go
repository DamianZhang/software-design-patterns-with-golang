package main

import (
	"fmt"
	"sp-1-facade/structs"
)

func main() {
	fmt.Println("prescriber system is starting...")

	prescriberSystemFacade := structs.NewPrescriberSystemFacade()
	prescriberSystemFacade.LoadPatientsFromJSONFile("./patients.json")
	prescriberSystemFacade.LoadDelegatorsOfPrescriptionDiagnosticerFromPlainText("./delegators_of_prescription_diagnosticer.txt")
	prescriberSystemFacade.DiagnosisAndHandlePrescriptions()
}
