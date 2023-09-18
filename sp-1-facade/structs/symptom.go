package structs

type Symptom struct {
	Name string
}

func NewSymptom(name string) *Symptom {
	return &Symptom{Name: name}
}

func NewSymptoms(namesOfSymptoms []string) []*Symptom {
	symptoms := make([]*Symptom, 0)

	for _, nameOfSymptom := range namesOfSymptoms {
		symptoms = append(symptoms, NewSymptom(nameOfSymptom))
	}

	return symptoms
}
