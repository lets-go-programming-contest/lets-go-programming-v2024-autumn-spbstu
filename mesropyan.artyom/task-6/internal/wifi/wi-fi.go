package wifi

import (
	"fmt"
	"net"

	"github.com/mdlayher/wifi"
)

//go:generate mockery --all --testonly --quiet --outpkg wifi_test --output .

type WiFi interface {
	Interfaces() ([]*wifi.Interface, error)
}

type Service struct {
	WiFi WiFi
}

func New(wifi WiFi) Service {
	return Service{WiFi: wifi}
}

func (service Service) GetAddresses() ([]net.HardwareAddr, error) {
	interfaces, err := service.WiFi.Interfaces()
	if err != nil {
		return nil, fmt.Errorf("error occurred: %w", err)
	}

	addrs := []net.HardwareAddr{}

	for _, iface := range interfaces {
		addrs = append(addrs, iface.HardwareAddr)
	}

	return addrs, nil
}

func (service Service) GetNames() ([]string, error) {
	interfaces, err := service.WiFi.Interfaces()
	if err != nil {
		return nil, fmt.Errorf("error occurred: %w", err)
	}

	nameList := []string{}

	for _, iface := range interfaces {
		nameList = append(nameList, iface.Name)
	}

	return nameList, nil
}
