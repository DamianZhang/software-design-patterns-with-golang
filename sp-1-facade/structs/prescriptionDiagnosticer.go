package structs

import (
	"fmt"
	"time"
)

type PrescriptionDiagnosticer struct {
	delegatorsOfPrescriptionDiagnosticer []DelegatorOfPrescriptionDiagnosticer
}

type DelegatorOfPrescriptionDiagnosticer interface {
	MatchSymptoms(prescriptionDiagnosticer *PrescriptionDiagnosticer, prescriptionDemand *PrescriptionDemand) bool
	PrescriptionTreatment(prescriptionDiagnosticer *PrescriptionDiagnosticer, prescriptionDemand *PrescriptionDemand) (*Prescription, error)
}

func NewPrescriptionDiagnosticer() *PrescriptionDiagnosticer {
	return &PrescriptionDiagnosticer{delegatorsOfPrescriptionDiagnosticer: make([]DelegatorOfPrescriptionDiagnosticer, 0)}
}

func (p *PrescriptionDiagnosticer) HandlePrescriptionDemand(prescriptionDemand *PrescriptionDemand, prescriptionsQueue chan<- []*Prescription) {
	time.Sleep(3 * time.Second)
	prescriptions := make([]*Prescription, 0)

	for _, delegatorOfPrescriptionDiagnosticer := range p.delegatorsOfPrescriptionDiagnosticer {
		if delegatorOfPrescriptionDiagnosticer.MatchSymptoms(p, prescriptionDemand) {
			prescription, err := delegatorOfPrescriptionDiagnosticer.PrescriptionTreatment(p, prescriptionDemand)
			if err != nil {
				fmt.Println(err)
				return
			}

			prescriptions = append(prescriptions, prescription)
		}
	}

	if len(prescriptions) == 0 {
		prescription, err := p.generatePlaceboPrescription(prescriptionDemand)
		if err != nil {
			fmt.Println(err)
			return
		}

		prescriptions = append(prescriptions, prescription)
	}

	prescriptionsQueue <- prescriptions
}

func (p *PrescriptionDiagnosticer) generatePlaceboPrescription(prescriptionDemand *PrescriptionDemand) (*Prescription, error) {
	var (
		name             = "無法診斷"
		potentialDisease = "無法診斷"
		usage            = "溫水配維他命，無法診斷也安心。"
		namesOfMedicines = []string{"維他命安慰劑"}
	)

	return p.TemplateOfPrescriptionTreatment(prescriptionDemand, name, potentialDisease, usage, namesOfMedicines)
}

func (p *PrescriptionDiagnosticer) AddDelegatorOfPrescriptionDiagnosticer(delegatorOfPrescriptionDiagnosticer DelegatorOfPrescriptionDiagnosticer) {
	p.delegatorsOfPrescriptionDiagnosticer = append(p.delegatorsOfPrescriptionDiagnosticer, delegatorOfPrescriptionDiagnosticer)
}

func (p *PrescriptionDiagnosticer) HasSymptom(symptoms []*Symptom, nameOfSymptom string) bool {
	for _, symptom := range symptoms {
		if symptom.Name == nameOfSymptom {
			return true
		}
	}

	return false
}

func (p *PrescriptionDiagnosticer) TemplateOfPrescriptionTreatment(prescriptionDemand *PrescriptionDemand, name, potentialDisease, usage string, namesOfMedicines []string) (*Prescription, error) {
	medicines, err := NewMedicines(namesOfMedicines)
	if err != nil {
		return nil, err
	}

	return NewPrescription(prescriptionDemand, name, potentialDisease, usage, medicines)
}
