package structs

import "errors"

type Medicine struct {
	Name string
}

func NewMedicine(name string) (*Medicine, error) {
	m := &Medicine{}

	if err := m.setName(name); err != nil {
		return nil, err
	}

	return m, nil
}

func NewMedicines(namesOfMedicines []string) ([]*Medicine, error) {
	medicines := make([]*Medicine, 0)

	for _, nameOfMedicine := range namesOfMedicines {
		medicine, err := NewMedicine(nameOfMedicine)
		if err != nil {
			return nil, err
		}

		medicines = append(medicines, medicine)
	}

	return medicines, nil
}

func (m *Medicine) setName(name string) error {
	if !LenOfStringShouldBeInRange(name, 8, 30) {
		return errors.New("len of medicine name should be in range 8-30")
	}

	m.Name = name
	return nil
}
