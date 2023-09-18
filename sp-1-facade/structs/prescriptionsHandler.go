package structs

type PrescriptionsHandler interface {
	HandlePrescriptions(prescriptions []*Prescription)
}
