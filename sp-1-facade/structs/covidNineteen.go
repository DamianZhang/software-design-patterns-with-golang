package structs

type COVIDNineteen struct{}

func NewCOVIDNineteen() *COVIDNineteen {
	return &COVIDNineteen{}
}

func (c *COVIDNineteen) MatchSymptoms(prescriptionDiagnosticer *PrescriptionDiagnosticer, prescriptionDemand *PrescriptionDemand) bool {
	var (
		symptoms    = prescriptionDemand.Symptoms()
		hasSneeze   = prescriptionDiagnosticer.HasSymptom(symptoms, "Sneeze")
		hasHeadache = prescriptionDiagnosticer.HasSymptom(symptoms, "Headache")
		hasCough    = prescriptionDiagnosticer.HasSymptom(symptoms, "Cough")
	)

	return hasSneeze && hasHeadache && hasCough
}

func (c *COVIDNineteen) PrescriptionTreatment(prescriptionDiagnosticer *PrescriptionDiagnosticer, prescriptionDemand *PrescriptionDemand) (*Prescription, error) {
	var (
		name             = "清冠一號"
		potentialDisease = "新冠肺炎（專業學名：COVID-19）"
		usage            = "將相關藥材裝入茶包裡，使用 500 ML 溫、熱水沖泡悶煮 1~3 分鐘後即可飲用。"
		namesOfMedicines = []string{"清冠一號"}
	)

	return prescriptionDiagnosticer.TemplateOfPrescriptionTreatment(prescriptionDemand, name, potentialDisease, usage, namesOfMedicines)
}

func (c *COVIDNineteen) String() string {
	return "COVID-19"
}

func (c *COVIDNineteen) MatchNameOfDelegatorOfPrescriptionDiagnosticer(nameOfDelegatorOfPrescriptionDiagnosticer string) bool {
	return c.String() == nameOfDelegatorOfPrescriptionDiagnosticer
}

func (c *COVIDNineteen) GenerateDelegatorOfPrescriptionDiagnosticer() DelegatorOfPrescriptionDiagnosticer {
	return c
}
