package wifi

import (
	"fmt"
	"net"

	"github.com/mdlayher/wifi"
)

type Interface interface {
	Interfaces() ([]*wifi.Interface, error)
}

type Service struct {
	WiFi Interface
}

func New(wifi Interface) Service {
	return Service{WiFi: wifi}
}

func (s Service) GetAddresses() ([]net.HardwareAddr, error) {
	interfaces, err := s.WiFi.Interfaces()
	if err != nil {
		return nil, fmt.Errorf("failed to get interfaces: %w", err)
	}

	addrs := make([]net.HardwareAddr, 0, len(interfaces))
	for _, iface := range interfaces {
		addrs = append(addrs, iface.HardwareAddr)
	}

	return addrs, nil
}

func (s Service) GetNames() ([]string, error) {
	interfaces, err := s.WiFi.Interfaces()
	if err != nil {
		return nil, fmt.Errorf("failed to get interfaces: %w", err)
	}

	names := make([]string, 0, len(interfaces))
	for _, iface := range interfaces {
		names = append(names, iface.Name)
	}

	return names, nil
}
