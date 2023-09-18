package structs

type SleepApneaSyndrome struct{}

func NewSleepApneaSyndrome() *SleepApneaSyndrome {
	return &SleepApneaSyndrome{}
}

func (s *SleepApneaSyndrome) MatchSymptoms(prescriptionDiagnosticer *PrescriptionDiagnosticer, prescriptionDemand *PrescriptionDemand) bool {
	var (
		patient  = prescriptionDemand.Patient()
		height   = patient.Height
		weight   = patient.Weight
		BMI      = CalculateBMI(height, weight)
		symptoms = prescriptionDemand.Symptoms()
		hasSnore = prescriptionDiagnosticer.HasSymptom(symptoms, "Snore")
	)

	return BMI > 26 && hasSnore
}

func (s *SleepApneaSyndrome) PrescriptionTreatment(prescriptionDiagnosticer *PrescriptionDiagnosticer, prescriptionDemand *PrescriptionDemand) (*Prescription, error) {
	var (
		name             = "打呼抑制劑"
		potentialDisease = "睡眠呼吸中止症（專業學名：SleepApneaSyndrome）"
		usage            = "睡覺時，撕下兩塊膠帶，將兩塊膠帶交錯黏在關閉的嘴巴上，就不會打呼了。"
		namesOfMedicines = []string{"一捲膠帶"}
	)

	return prescriptionDiagnosticer.TemplateOfPrescriptionTreatment(prescriptionDemand, name, potentialDisease, usage, namesOfMedicines)
}

func (s *SleepApneaSyndrome) String() string {
	return "SleepApneaSyndrome"
}

func (s *SleepApneaSyndrome) MatchNameOfDelegatorOfPrescriptionDiagnosticer(nameOfDelegatorOfPrescriptionDiagnosticer string) bool {
	return s.String() == nameOfDelegatorOfPrescriptionDiagnosticer
}

func (s *SleepApneaSyndrome) GenerateDelegatorOfPrescriptionDiagnosticer() DelegatorOfPrescriptionDiagnosticer {
	return s
}
