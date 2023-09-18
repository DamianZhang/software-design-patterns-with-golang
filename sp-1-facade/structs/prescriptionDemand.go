package structs

type PrescriptionDemand struct {
	patient    *Patient
	symptoms   []*Symptom
	fileName   string
	fileFormat string
}

func NewPrescriptionDemand(patient *Patient, symptoms []*Symptom, fileName, fileFormat string) *PrescriptionDemand {
	return &PrescriptionDemand{
		patient:    patient,
		symptoms:   symptoms,
		fileName:   fileName,
		fileFormat: fileFormat,
	}
}

func (p *PrescriptionDemand) Patient() *Patient {
	return p.patient
}

func (p *PrescriptionDemand) Symptoms() []*Symptom {
	return p.symptoms
}

func (p *PrescriptionDemand) FileName() string {
	return p.fileName
}

func (p *PrescriptionDemand) FileFormat() string {
	return p.fileFormat
}
