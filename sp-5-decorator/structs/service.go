package structs

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Service struct {
	domainName string
	ips        []string
}

func NewService(domainName string, ips []string) *Service {
	return &Service{
		domainName: domainName,
		ips:        ips,
	}
}

func (s *Service) DomainName() string {
	return s.domainName
}

func (s *Service) IPs() []string {
	return s.ips
}

func (s *Service) GetAvailableIP() (availableIP string) {
	for _, ip := range s.ips {
		if ip != "" {
			return ip
		}
	}

	return ""
}

func (s *Service) GetAvailableIPs() (availableIPs string) {
	for _, ip := range s.ips {
		if ip != "" {
			availableIPs += ip + ","
		}
	}

	return strings.TrimRight(availableIPs, ",")
}

func (s *Service) AddIP(ip string) {
	for i, ipOfService := range s.ips {
		if ipOfService == "" {
			s.ips[i] = ip
			return
		}
	}

	s.ips = append(s.ips, ip)
}

func (s *Service) DeleteIP(ip string) {
	for i, ipOfService := range s.ips {
		if ipOfService == ip {
			s.ips[i] = ""
			return
		}
	}
}

func LoadServicesFromPlainText(filePath string) ([]*Service, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("load services from plain text failed: %s", err)
	}
	defer file.Close()

	var (
		scanner  = bufio.NewScanner(file)
		services = make([]*Service, 0)
	)
	for scanner.Scan() {
		var (
			infoOfService = strings.Split(scanner.Text(), ": ")
			domainName    = infoOfService[0]
			ips           = strings.Split(infoOfService[1], ", ")
		)
		services = append(services, NewService(domainName, ips))
	}

	return services, nil
}
