package wifi_test

import (
	"errors"
	"fmt"
	"net"
	"testing"

	myWiFi "github.com/katagiriwhy/task-6/internal/wifi"
	"github.com/mdlayher/wifi"
	"github.com/stretchr/testify/require"
)

type MockWifi struct{}

func (mockWifi *MockWifi) Interfaces() ([]*wifi.Interface, error) {
	return nil, nil
}

func TestNew(t *testing.T) {
	t.Parallel()

	mockWifi := &MockWifi{}
	wiFiService := myWiFi.New(mockWifi)

	require.NotNil(t, wiFiService)
	require.Equal(t, mockWifi, wiFiService.WiFi, "Wifi service should be equal with")
}

type testWifi struct {
	addrs       []string
	errExpected error
}

var testTable = []testWifi{
	{
		addrs: []string{"1111:2222:3333:4444", "8800:0000:8800:9999"},
	},
	{
		errExpected: errors.New("no addresses provided"),
	},
}

func mockIfaces(addrs []string) []*wifi.Interface {
	var ifaces []*wifi.Interface
	for i, addr := range addrs {
		hardwareAddr := parseMAC(addr)
		if hardwareAddr == nil {
			continue
		}
		intrface := &wifi.Interface{
			Index:        i + 1,
			Name:         fmt.Sprintf("eth%d", i+1),
			HardwareAddr: hardwareAddr,
			PHY:          1,
			Device:       1,
			Type:         wifi.InterfaceTypeAPVLAN,
			Frequency:    0,
		}
		ifaces = append(ifaces, intrface)
	}
	return ifaces
}

func TestGetAddresses(t *testing.T) {
	t.Parallel()

	mockWifi := NewWiFi(t)
	wifiService := myWiFi.Service{WiFi: mockWifi}
	for i, row := range testTable {
		mockWifi.On("Interfaces").Unset()
		mockWifi.On("Interfaces").Return(mockIfaces(row.addrs), row.errExpected)
		ourAddrs, err := wifiService.GetAddresses()
		if row.errExpected != nil {
			require.ErrorIs(t, err, row.errExpected, "row %d, expected error: %w, actual error: %w", i, row.errExpected, err)
			continue
		}
		require.NoError(t, err, "row %d, error must be nil", i)
		require.Equal(t, parseMACs(row.addrs), parseMACs(row.addrs), "row: %d, expected addrs: %s, actual addrs: %s", i, parseMACs(row.addrs), ourAddrs)
	}
}

func TestGetNames(t *testing.T) {
	t.Parallel()

	mockWifi := NewWiFi(t)
	wifiService := myWiFi.Service{WiFi: mockWifi}
	for i, row := range testTable {
		mockWifi.On("Interfaces").Unset()
		mockWifi.On("Interfaces").Return(mockIfaces(row.addrs), row.errExpected)
		_, err := wifiService.GetNames()
		if row.errExpected != nil {
			require.ErrorIs(t, err, row.errExpected, "row %d, expected error: %w, actual error: %w", i, row.errExpected, err)
			continue
		}
		require.NoError(t, err, "row %d, error must be nil", i)
	}
}

func parseMACs(macStr []string) []net.HardwareAddr {
	var addrs []net.HardwareAddr
	for _, addr := range macStr {
		addrs = append(addrs, parseMAC(addr))
	}
	return addrs
}

func parseMAC(macStr string) net.HardwareAddr {
	hwAddr, err := net.ParseMAC(macStr)
	if err != nil {
		return nil
	}
	return hwAddr
}

// go tool cover -html=prof
