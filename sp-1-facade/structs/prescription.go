package structs

import "errors"

type Prescription struct {
	prescriptionDemand *PrescriptionDemand
	Name               string
	PotentialDisease   string
	Usage              string
	Medicines          []*Medicine
}

func NewPrescription(prescriptionDemand *PrescriptionDemand, name, potentialDisease, usage string, medicines []*Medicine) (*Prescription, error) {
	p := &Prescription{Medicines: make([]*Medicine, 0)}

	if err := p.setName(name); err != nil {
		return nil, err
	}

	if err := p.setPotentialDisease(potentialDisease); err != nil {
		return nil, err
	}

	if err := p.setUsage(usage); err != nil {
		return nil, err
	}

	p.prescriptionDemand = prescriptionDemand
	p.Medicines = medicines
	return p, nil
}

func (p *Prescription) PrescriptionDemand() *PrescriptionDemand {
	return p.prescriptionDemand
}

func (p *Prescription) setName(name string) error {
	if !LenOfStringShouldBeInRange(name, 4, 30) {
		return errors.New("len of prescription name should be in range 4-30")
	}

	p.Name = name
	return nil
}

func (p *Prescription) setPotentialDisease(potentialDisease string) error {
	if !LenOfStringShouldBeInRange(potentialDisease, 8, 100) {
		return errors.New("len of prescription potential disease should be in range 8-100")
	}

	p.PotentialDisease = potentialDisease
	return nil
}

func (p *Prescription) setUsage(usage string) error {
	if !LenOfStringShouldBeInRange(usage, 0, 1000) {
		return errors.New("len of prescription usage should be in range 0-1000")
	}

	p.Usage = usage
	return nil
}
