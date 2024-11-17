package wifi_test

import (
	"bytes"
	"errors"
	"net"
	"testing"

	"github.com/mdlayher/wifi"
	myWifi "github.com/mrqiz/task-6/internal/wifi"
)

type MockWiFi struct {
	interfaces []*wifi.Interface
	err        error
}

func (m *MockWiFi) Interfaces() ([]*wifi.Interface, error) {
	return m.interfaces, m.err
}

func TestGetAddresses(t *testing.T) {
	mockWiFi := &MockWiFi{
		interfaces: []*wifi.Interface{
			{Name: "wlan0", HardwareAddr: net.HardwareAddr{0x00, 0x11, 0x22, 0x33, 0x44, 0x55}},
			{Name: "wlan1", HardwareAddr: net.HardwareAddr{0x66, 0x77, 0x88, 0x99, 0xAA, 0xBB}},
		},
	}

	service := myWifi.New(mockWiFi)

	addresses, err := service.GetAddresses()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	expected := []net.HardwareAddr{
		{0x00, 0x11, 0x22, 0x33, 0x44, 0x55},
		{0x66, 0x77, 0x88, 0x99, 0xAA, 0xBB},
	}

	if len(addresses) != len(expected) {
		t.Fatalf("expected %d addresses, got %d", len(expected), len(addresses))
	}

	for i, addr := range addresses {
		if !bytes.Equal(addr, expected[i]) {
			t.Errorf("expected address %v, got %v", expected[i], addr)
		}
	}
}

func TestGetNames(t *testing.T) {
	mockWiFi := &MockWiFi{
		interfaces: []*wifi.Interface{
			{Name: "wlan0", HardwareAddr: net.HardwareAddr{0x00, 0x11, 0x22, 0x33, 0x44, 0x55}},
			{Name: "wlan1", HardwareAddr: net.HardwareAddr{0x66, 0x77, 0x88, 0x99, 0xAA, 0xBB}},
		},
	}

	service := myWifi.New(mockWiFi)

	names, err := service.GetNames()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	expected := []string{"wlan0", "wlan1"}

	if len(names) != len(expected) {
		t.Fatalf("expected %d names, got %d", len(expected), len(names))
	}

	for i, name := range names {
		if name != expected[i] {
			t.Errorf("expected name %s, got %s", expected[i], name)
		}
	}
}

var errSomeError = errors.New("some error occurred")

func TestGetAddressesError(t *testing.T) {
	mockWiFi := &MockWiFi{
		err: errSomeError,
	}

	service := myWifi.New(mockWiFi)

	addresses, err := service.GetAddresses()
	if err == nil {
		t.Fatal("expected an error, got none")
	}

	if addresses != nil {
		t.Fatalf("expected nil addresses, got %v", addresses)
	}
}

func TestGetNamesError(t *testing.T) {
	mockWiFi := &MockWiFi{
		err: errSomeError,
	}

	service := myWifi.New(mockWiFi)

	names, err := service.GetNames()
	if err == nil {
		t.Fatal("expected an error, got none")
	}

	if names != nil {
		t.Fatalf("expected nil names, got %v", names)
	}
}

