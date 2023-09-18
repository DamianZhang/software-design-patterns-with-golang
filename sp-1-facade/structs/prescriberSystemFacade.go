package structs

import (
	"fmt"
	"strings"
	"time"
)

type PrescriberSystemFacade struct {
	fileStorer            *FileStorer
	db                    *PatientDatabase
	prescriber            *Prescriber
	prescriptionsQueue    chan []*Prescription
	prescriptionsHandlers []PrescriptionsHandler
}

func NewPrescriberSystemFacade() *PrescriberSystemFacade {
	p := &PrescriberSystemFacade{
		fileStorer:            NewFileStorer(),
		db:                    NewPatientDatabase(),
		prescriber:            NewPrescriber(),
		prescriptionsQueue:    make(chan []*Prescription, 100),
		prescriptionsHandlers: make([]PrescriptionsHandler, 0),
	}

	p.AddPrescriptionsHandler(p.fileStorer)
	p.AddPrescriptionsHandler(p.db)
	return p
}

func (p *PrescriberSystemFacade) LoadPatientsFromJSONFile(fileName string) {
	jsonParser := NewJSONParser()
	patients, err := jsonParser.Parse(fileName)
	if err != nil {
		fmt.Println("load patients from JSON file failed:", err)
		return
	}

	p.db.SetPatients(patients)

	fmt.Println("load patients from JSON file successfully")
}

func (p *PrescriberSystemFacade) LoadDelegatorsOfPrescriptionDiagnosticerFromPlainText(fileName string) {
	plainTextParser := NewPlainTextParser()
	delegatorsOfPrescriptionDiagnosticer, err := plainTextParser.Parse(fileName)
	if err != nil {
		fmt.Println("load delegators of prescription diagnosticer from plain text failed:", err)
		return
	}

	for _, delegatorOfPrescriptionDiagnosticer := range delegatorsOfPrescriptionDiagnosticer {
		p.prescriber.PrescriptionDiagnosticer().AddDelegatorOfPrescriptionDiagnosticer(delegatorOfPrescriptionDiagnosticer)
	}

	fmt.Println("load delegators of prescription diagnosticer from plain text successfully")
}

func (p *PrescriberSystemFacade) DiagnosisAndHandlePrescriptions() {
	p.prescriber.HandlePrescriptionDemand(p.prescriptionsQueue)
	p.handlePrescriptions()

	for {
		p.db.PrintPatients()

		patientId := p.scanPatientId()
		symptoms := p.scanSymptoms()
		fileName := p.scanFileName()
		fileFormat := p.scanFileFormat()
		p.diagnosis(patientId, symptoms, fileName, fileFormat)
	}
}

func (p *PrescriberSystemFacade) scanPatientId() string {
	var patientId string

	for {
		fmt.Println("please input patient ID:")
		fmt.Scanf("%s", &patientId)

		if IsLegalPatientId(patientId) {
			return patientId
		} else {
			fmt.Println("please enter the CORRECT patient ID format, starting with capital letters")
		}
	}
}

func (p *PrescriberSystemFacade) scanSymptoms() []*Symptom {
	var symptoms string

	fmt.Println("please input symptoms(separated by \",\"):")
	fmt.Println("example: Sneeze,Headache,Cough")
	fmt.Scanf("%s", &symptoms)

	namesOfSymptoms := strings.Split(symptoms, ",")
	return NewSymptoms(namesOfSymptoms)
}

func (p *PrescriberSystemFacade) scanFileName() string {
	var fileName string

	fmt.Println("please input file name:")
	fmt.Scanf("%s", &fileName)

	return fileName
}

func (p *PrescriberSystemFacade) scanFileFormat() string {
	var (
		fileFormat                  string
		stringOfSupportedFileFormat = p.stringOfSupportedFileFormat()
	)

	for {
		fmt.Printf("please input file format (%s):\n", stringOfSupportedFileFormat)
		fmt.Scanf("%s", &fileFormat)

		if p.isSupportedFileFormat(fileFormat) {
			return fileFormat
		} else {
			fmt.Printf("please input LEGAL file format (%s)\n", stringOfSupportedFileFormat)
		}
	}
}

func (p *PrescriberSystemFacade) stringOfSupportedFileFormat() (stringOfSupportedFileFormat string) {
	for _, delegatorOfFileStorer := range p.fileStorer.DelegatorsOfFileStorer() {
		stringOfSupportedFileFormat += delegatorOfFileStorer.String() + "/"
	}

	return stringOfSupportedFileFormat
}

func (p *PrescriberSystemFacade) isSupportedFileFormat(fileFormat string) bool {
	for _, delegatorOfFileStorer := range p.fileStorer.DelegatorsOfFileStorer() {
		if fileFormat == delegatorOfFileStorer.String() {
			return true
		}
	}

	return false
}

func (p *PrescriberSystemFacade) diagnosis(patientId string, symptoms []*Symptom, fileName, fileFormat string) {
	patient := p.db.SearchPatientByPatientId(patientId)
	if patient == nil {
		fmt.Printf("%s could NOT find patient\n", patientId)
		return
	}

	prescriptionDemand := NewPrescriptionDemand(patient, symptoms, fileName, fileFormat)
	p.prescriber.AddPrescriptionDemand(prescriptionDemand)
}

func (p *PrescriberSystemFacade) handlePrescriptions() {
	go func() {
		for {
			select {
			case prescriptions := <-p.prescriptionsQueue:
				fmt.Println("【Get prescriptions】")

				for _, prescriptionsHandler := range p.prescriptionsHandlers {
					go prescriptionsHandler.HandlePrescriptions(prescriptions)
				}
			default:
				time.Sleep(1 * time.Second)
			}
		}
	}()
}

func (p *PrescriberSystemFacade) AddPrescriptionsHandler(prescriptionsHandler PrescriptionsHandler) {
	p.prescriptionsHandlers = append(p.prescriptionsHandlers, prescriptionsHandler)
}
