package structs

import (
	"bufio"
	"os"
)

type PlainTextParser struct {
	delegatorsOfPlainTextParser []DelegatorOfPlainTextParser
}

type DelegatorOfPlainTextParser interface {
	String() string
	MatchNameOfDelegatorOfPrescriptionDiagnosticer(nameOfDelegatorOfPrescriptionDiagnosticer string) bool
	GenerateDelegatorOfPrescriptionDiagnosticer() DelegatorOfPrescriptionDiagnosticer
}

func NewPlainTextParser() *PlainTextParser {
	p := &PlainTextParser{delegatorsOfPlainTextParser: make([]DelegatorOfPlainTextParser, 0)}

	p.AddDelegatorOfPlainTextParser(NewCOVIDNineteen())
	p.AddDelegatorOfPlainTextParser(NewAttractive())
	p.AddDelegatorOfPlainTextParser(NewSleepApneaSyndrome())
	return p
}

func (p *PlainTextParser) Parse(fileName string) ([]DelegatorOfPrescriptionDiagnosticer, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	delegatorsOfPrescriptionDiagnosticer := make([]DelegatorOfPrescriptionDiagnosticer, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		for _, delegatorOfPlainTextParser := range p.delegatorsOfPlainTextParser {
			if delegatorOfPlainTextParser.MatchNameOfDelegatorOfPrescriptionDiagnosticer(scanner.Text()) {
				delegatorOfPrescriptionDiagnosticer := delegatorOfPlainTextParser.GenerateDelegatorOfPrescriptionDiagnosticer()
				delegatorsOfPrescriptionDiagnosticer = append(delegatorsOfPrescriptionDiagnosticer, delegatorOfPrescriptionDiagnosticer)
			}
		}
	}

	return delegatorsOfPrescriptionDiagnosticer, nil
}

func (p *PlainTextParser) AddDelegatorOfPlainTextParser(delegatorOfPlainTextParser DelegatorOfPlainTextParser) {
	p.delegatorsOfPlainTextParser = append(p.delegatorsOfPlainTextParser, delegatorOfPlainTextParser)
}
