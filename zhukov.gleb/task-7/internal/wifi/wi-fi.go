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
		return nil, fmt.Errorf("GetAddresses interfaces err: %v", err)
	}

	addrs := make([]net.HardwareAddr, 0)
	for _, iface := range interfaces {
		addrs = append(addrs, iface.HardwareAddr)
	}

	return addrs, nil
}

func (service Service) GetNames() ([]string, error) {
	interfaces, err := service.WiFi.Interfaces()
	if err != nil {
		return nil, fmt.Errorf("GetNames interfaces err: %v", err)
	}

	name_list := make([]string, 0)
	for _, iface := range interfaces {
		name_list = append(name_list, iface.Name)
	}

	return name_list, nil
}
