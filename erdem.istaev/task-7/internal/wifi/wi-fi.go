package wifi

import (
	"fmt"
	"net"

	"github.com/mdlayher/wifi"
)

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
		return nil, fmt.Errorf("get addresses: %w", err)
	}

	addrs := make([]net.HardwareAddr, len(interfaces))
	for i, iface := range interfaces {
		addrs[i] = iface.HardwareAddr
	}

	return addrs, nil
}

func (service Service) GetNames() ([]string, error) {
	interfaces, err := service.WiFi.Interfaces()
	if err != nil {
		return nil, fmt.Errorf("get names: %w", err)
	}

	nameList := make([]string, len(interfaces))
	for i, iface := range interfaces {
		nameList[i] = iface.Name
	}

	return nameList, nil
}
