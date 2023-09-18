package structs

import (
	"fmt"
	"time"
)

type Prescriber struct {
	prescriptionDemandQueue  chan *PrescriptionDemand
	prescriptionDiagnosticer *PrescriptionDiagnosticer
}

func NewPrescriber() *Prescriber {
	return &Prescriber{
		prescriptionDemandQueue:  make(chan *PrescriptionDemand, 100),
		prescriptionDiagnosticer: NewPrescriptionDiagnosticer(),
	}
}

func (p *Prescriber) AddPrescriptionDemand(prescriptionDemand *PrescriptionDemand) {
	p.prescriptionDemandQueue <- prescriptionDemand
}

func (p *Prescriber) HandlePrescriptionDemand(prescriptionsQueue chan []*Prescription) {
	go func() {
		for {
			select {
			case prescriptionDemand := <-p.prescriptionDemandQueue:
				fmt.Println("【Get prescriptionDemand】")

				go p.prescriptionDiagnosticer.HandlePrescriptionDemand(prescriptionDemand, prescriptionsQueue)
			default:
				time.Sleep(1 * time.Second)
			}
		}
	}()
}

func (p *Prescriber) PrescriptionDiagnosticer() *PrescriptionDiagnosticer {
	return p.prescriptionDiagnosticer
}
