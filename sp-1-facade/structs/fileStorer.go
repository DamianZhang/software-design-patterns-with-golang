package structs

type FileStorer struct {
	delegatorsOfFileStorer []DelegatorOfFileStorer
}

type DelegatorOfFileStorer interface {
	String() string
	MatchFileFormat(fileFormat string) bool
	StorePrescriptions(prescriptions []*Prescription, fileName string)
}

func NewFileStorer() *FileStorer {
	f := &FileStorer{delegatorsOfFileStorer: make([]DelegatorOfFileStorer, 0)}

	f.AddDelegatorOfFileStorer(NewJSONStorer())
	f.AddDelegatorOfFileStorer(NewCSVStorer())
	return f
}

func (f *FileStorer) HandlePrescriptions(prescriptions []*Prescription) {
	var (
		prescriptionDemand = prescriptions[0].PrescriptionDemand()
		fileName           = prescriptionDemand.FileName()
		fileFormat         = prescriptionDemand.FileFormat()
	)

	for _, delegatorOfFileStorer := range f.delegatorsOfFileStorer {
		if delegatorOfFileStorer.MatchFileFormat(fileFormat) {
			delegatorOfFileStorer.StorePrescriptions(prescriptions, fileName)
		}
	}
}

func (f *FileStorer) DelegatorsOfFileStorer() []DelegatorOfFileStorer {
	return f.delegatorsOfFileStorer
}

func (f *FileStorer) AddDelegatorOfFileStorer(delegatorOfFileStorer DelegatorOfFileStorer) {
	f.delegatorsOfFileStorer = append(f.delegatorsOfFileStorer, delegatorOfFileStorer)
}
