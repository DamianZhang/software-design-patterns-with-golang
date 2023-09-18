package structs

type Attractive struct{}

func NewAttractive() *Attractive {
	return &Attractive{}
}

func (a *Attractive) MatchSymptoms(prescriptionDiagnosticer *PrescriptionDiagnosticer, prescriptionDemand *PrescriptionDemand) bool {
	var (
		patient   = prescriptionDemand.Patient()
		age       = patient.Age
		gender    = patient.Gender
		symptoms  = prescriptionDemand.Symptoms()
		hasSneeze = prescriptionDiagnosticer.HasSymptom(symptoms, "Sneeze")
	)

	return age == 18 && gender == Female && hasSneeze
}

func (a *Attractive) PrescriptionTreatment(prescriptionDiagnosticer *PrescriptionDiagnosticer, prescriptionDemand *PrescriptionDemand) (*Prescription, error) {
	var (
		name             = "青春抑制劑"
		potentialDisease = "有人想你了（專業學名：Attractive）"
		usage            = "把假鬢角黏在臉的兩側，讓自己異性緣差一點，自然就不會有人想妳了。"
		namesOfMedicines = []string{"假的鬢角", "臭臭的味道"}
	)

	return prescriptionDiagnosticer.TemplateOfPrescriptionTreatment(prescriptionDemand, name, potentialDisease, usage, namesOfMedicines)
}

func (a *Attractive) String() string {
	return "Attractive"
}

func (a *Attractive) MatchNameOfDelegatorOfPrescriptionDiagnosticer(nameOfDelegatorOfPrescriptionDiagnosticer string) bool {
	return a.String() == nameOfDelegatorOfPrescriptionDiagnosticer
}

func (a *Attractive) GenerateDelegatorOfPrescriptionDiagnosticer() DelegatorOfPrescriptionDiagnosticer {
	return a
}
