package structs

import "time"

type CaseOfPatient struct {
	Symptoms      []*Symptom
	Prescriptions []*Prescription
	CaseTime      time.Time
}

func NewCaseOfPatient(symptoms []*Symptom, prescriptions []*Prescription) *CaseOfPatient {
	return &CaseOfPatient{
		Symptoms:      symptoms,
		Prescriptions: prescriptions,
		CaseTime:      time.Now(),
	}
}
