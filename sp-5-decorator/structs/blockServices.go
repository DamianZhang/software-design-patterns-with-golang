package structs

import "fmt"

type BlockServices []*Service

func NewBlockServices(filePath string) (*BlockServices, error) {
	blockServices, err := LoadServicesFromPlainText(filePath)
	if err != nil {
		return nil, fmt.Errorf("load block services from plain text failed: %s", err)
	}

	b := BlockServices(blockServices)
	return &b, nil
}
