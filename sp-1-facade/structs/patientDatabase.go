package structs

import (
	"fmt"
	"sync"
)

type PatientDatabase struct {
	patients []*Patient
	rwMutex  sync.RWMutex
}

func NewPatientDatabase() *PatientDatabase {
	return &PatientDatabase{patients: make([]*Patient, 0)}
}

func (db *PatientDatabase) HandlePrescriptions(prescriptions []*Prescription) {
	var (
		prescriptionDemand = prescriptions[0].PrescriptionDemand()
		patientId          = prescriptionDemand.Patient().Id
		symptoms           = prescriptionDemand.Symptoms()
		caseOfPatient      = NewCaseOfPatient(symptoms, prescriptions)
	)

	db.AddCaseOfPatientByPatientId(patientId, caseOfPatient)
}

func (db *PatientDatabase) AddCaseOfPatientByPatientId(patientId string, caseOfPatient *CaseOfPatient) {
	db.rwMutex.Lock()
	defer db.rwMutex.Unlock()

	for _, patient := range db.patients {
		if patient.Id == patientId {
			patient.AddCaseOfPatient(caseOfPatient)
			break
		}
	}

	// // 用下面這種方式會造成 dead lock
	// patient := db.SearchPatientByPatientId(patientId)
	// if patient != nil {
	// 	patient.AddCaseOfPatient(caseOfPatient)
	// }
}

func (db *PatientDatabase) SearchPatientByPatientId(patientId string) *Patient {
	db.rwMutex.RLock()
	defer db.rwMutex.RUnlock()

	for _, patient := range db.patients {
		if patient.Id == patientId {
			return patient
		}
	}

	return nil
}

func (db *PatientDatabase) PrintPatients() {
	db.rwMutex.RLock()
	defer db.rwMutex.RUnlock()

	for i, patient := range db.patients {
		fmt.Printf("patient %d: %v\n", i, patient)
	}
}

func (db *PatientDatabase) SetPatients(patients []*Patient) {
	db.rwMutex.Lock()
	defer db.rwMutex.Unlock()

	db.patients = patients
}
