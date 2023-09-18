package structs

import "fmt"

type AvailableServices []*Service

func NewAvailableServices(filePath string) (*AvailableServices, error) {
	availableServices, err := LoadServicesFromPlainText(filePath)
	if err != nil {
		return nil, fmt.Errorf("load available services from plain text failed: %s", err)
	}

	a := AvailableServices(availableServices)
	return &a, nil
}

func (a *AvailableServices) GetAvailableServiceByHostOfURL(hostOfURL string) *Service {
	for _, availableService := range *(a) {
		if a.isHostOfURLInAvailableServices(hostOfURL, availableService.DomainName()) || a.isHostOfURLInAvailableServices(hostOfURL, availableService.IPs()...) {
			return availableService
		}
	}

	return nil
}

func (a *AvailableServices) isHostOfURLInAvailableServices(hostOfURL string, availableServices ...string) bool {
	for _, availableService := range availableServices {
		if availableService == hostOfURL {
			return true
		}
	}

	return false
}
