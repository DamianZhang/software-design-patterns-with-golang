package structs

import (
	"errors"
)

type Patient struct {
	Id             string
	Name           string
	Gender         Gender
	Age            int
	Height         float64
	Weight         float64
	CasesOfPatient []*CaseOfPatient
}

func NewPatient(id, name string, gender Gender, age int, height, weight float64) (*Patient, error) {
	p := &Patient{CasesOfPatient: make([]*CaseOfPatient, 0)}

	if err := p.setId(id); err != nil {
		return nil, err
	}

	if err := p.setName(name); err != nil {
		return nil, err
	}

	if err := p.setAge(age); err != nil {
		return nil, err
	}

	if err := p.setHeight(height); err != nil {
		return nil, err
	}

	if err := p.setWeight(weight); err != nil {
		return nil, err
	}

	p.Gender = gender
	return p, nil
}

func (p *Patient) setId(id string) error {
	if !IsLegalPatientId(id) {
		return errors.New("please enter the CORRECT patient ID format, starting with capital letters")
	}

	p.Id = id
	return nil
}

func (p *Patient) setName(name string) error {
	if !LenOfStringShouldBeInRange(name, 1, 30) {
		return errors.New("len of patient name should be in range 1-30")
	}

	p.Name = name
	return nil
}

func (p *Patient) setAge(age int) error {
	if !IntShouldBeInRange(age, 1, 180) {
		return errors.New("patient age should be in range 1-180")
	}

	p.Age = age
	return nil
}

func (p *Patient) setHeight(height float64) error {
	if !Float64ShouldBeInRange(height, 1, 500) {
		return errors.New("patient height should be in range 1-500")
	}

	p.Height = height
	return nil
}

func (p *Patient) setWeight(weight float64) error {
	if !Float64ShouldBeInRange(weight, 1, 500) {
		return errors.New("patient weight should be in range 1-500")
	}

	p.Weight = weight
	return nil
}

func (p *Patient) AddCaseOfPatient(caseOfPatient *CaseOfPatient) {
	p.CasesOfPatient = append(p.CasesOfPatient, caseOfPatient)
}
