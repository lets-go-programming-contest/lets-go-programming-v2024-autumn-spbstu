package wifi_test

import (
	"errors"
	"fmt"
	"net"
	"testing"

	"github.com/mdlayher/wifi"
	"github.com/stretchr/testify/require"

	myWifi "github.com/Mmmakskl/task-6/internal/wifi"
)

type testTableWiFi struct {
	addrs         []string
	names         []string
	expectedError error
}

var testTable = []testTableWiFi{
	{
		addrs:         []string{"00:11:22:33:44:55", "aa:bb:cc:dd:ee:ff"},
		names:         []string{"wlan0", "wlan1"},
		expectedError: nil,
	},
	{
		addrs:         []string{"00:11:22:33:44:55", "aa:bb:cc:dd:ee:ff"},
		names:         []string{},
		expectedError: errors.New("ExpectedError"),
	},
	{
		addrs:         []string{},
		names:         []string{"wlan0", "wlan1"},
		expectedError: errors.New("ExpectedError"),
	},
	{
		addrs:         []string{"00:11:22:33:44:55", "aa:bb:cc:dd:ee:ff"},
		names:         nil,
		expectedError: errors.New("ExpectedError"),
	},
	{
		addrs:         nil,
		names:         []string{"wlan0", "wlan1"},
		expectedError: errors.New("ExpectedError"),
	},
	{
		addrs:         []string{},
		names:         []string{},
		expectedError: errors.New("ExpectedError"),
	},
	{
		addrs:         nil,
		names:         nil,
		expectedError: errors.New("ExpectedError"),
	},
}

func TestGetAddresses(t *testing.T) {
	t.Parallel()

	mockWiFi := NewWiFi(t)
	wifiService := myWifi.Service{WiFi: mockWiFi}

	for i, row := range testTable {
		mockWiFi.On("Interfaces").Unset()
		mockWiFi.On("Interfaces").Return(mockIfaces(row.addrs), row.expectedError)
		actualAddrs, err := wifiService.GetAddresses()

		if row.expectedError != nil {
			require.ErrorIs(t, err, row.expectedError, "row: %d, expected error: %w, actual error: %w", i, row.expectedError, err)
			continue
		}

		require.NoError(t, err, "row: %d, error must be nil", i)
		require.Equal(t, parseMACs(row.addrs), actualAddrs, "row %d, expected addrs: %s, actual addrs: %s", i, parseMACs(row.addrs), actualAddrs)
	}
}

func TestGetNames(t *testing.T) {
	t.Parallel()

	mockWiFi := NewWiFi(t)
	wifiService := myWifi.Service{WiFi: mockWiFi}

	for i, row := range testTable {
		mockWiFi.On("Interfaces").Unset()
		mockWiFi.On("Interfaces").Return(mockIfaces(row.addrs), row.expectedError)
		actualNames, err := wifiService.GetNames()

		if row.expectedError != nil {
			require.ErrorIs(t, err, row.expectedError, "row: %d, expected error: %w, actual error: %w", i, row.expectedError, err)
			continue
		}

		require.NoError(t, err, "row: %d, error must be nil", i)
		require.Equal(t, row.names, actualNames, "row %d, expected names: %s, actual names: %s", i, row.names, actualNames)
	}
}

func TestNew(t *testing.T) {
	t.Parallel()

	mockWiFi := NewWiFi(t)

	service := myWifi.New(mockWiFi)

	require.NotNil(t, service)
	require.Equal(t, mockWiFi, service.WiFi, "Service should have the provided WIFI")
}

func mockIfaces(addrs []string) []*wifi.Interface {

	var wifiIfaces []*wifi.Interface

	for i, addr := range addrs {

		hwAddr := parseMAC(addr)
		if hwAddr == nil {
			continue
		}

		iface := wifi.Interface{
			Index:        i + 1,
			Name:         fmt.Sprintf("wlan%d", i),
			HardwareAddr: hwAddr,
			PHY:          1,
			Device:       1,
			Type:         wifi.InterfaceTypeAP,
			Frequency:    2412,
		}

		wifiIfaces = append(wifiIfaces, &iface)
	}

	return wifiIfaces
}

func parseMACs(addrs []string) []net.HardwareAddr {
	var hwAddrs []net.HardwareAddr

	for _, addr := range addrs {
		hwAddrs = append(hwAddrs, parseMAC(addr))
	}

	return hwAddrs
}

func parseMAC(addrs string) net.HardwareAddr {
	hwAddrs, err := net.ParseMAC(addrs)
	if err != nil {
		return nil
	}

	return hwAddrs
}
