package wifi_test

import (
	"errors"
	"net"
	"testing"

	myWiFi "github.com/Madyarov-Gleb/task-8/point_3"
	"github.com/mdlayher/wifi"
	"github.com/stretchr/testify/require"
)

//go:generate mockery --all --testonly --quiet --outpkg wifi_test --output .

type rowTestSysInfo struct {
	addrs       []string
	name_list   []string
	errExpected error
}

var testTable = []rowTestSysInfo{
	{
		addrs:       []string{"00:11:22:33:44:55", "aa:bb:cc:dd:ee:ff"},
		name_list:   []string{"Ivan", "Gena"},
		errExpected: nil,
	},
	{
		addrs:       []string{"00:11:22:33:44:55", "aa:bb:cc:dd:ee:ff"},
		name_list:   nil,
		errExpected: errors.New("ExpectedError"),
	},
	{
		addrs:       nil,
		name_list:   []string{"Ivan", "Gena"},
		errExpected: errors.New("ExpectedError"),
	},
	{
		addrs:       nil,
		name_list:   nil,
		errExpected: errors.New("ExpectedError"),
	},
}

func TestGetNames(t *testing.T) {
	t.Parallel()

	mockWifi := NewWiFi(t)
	wifiService := myWiFi.Service{WiFi: mockWifi}

	for i, row := range testTable {
		mockWifi.On("Interfaces").Unset()
		mockWifi.On("Interfaces").Return(mockIfaces(row.name_list), row.errExpected)
		actualNames, err := wifiService.GetNames()
		if row.errExpected != nil {
			require.ErrorIs(t, err, row.errExpected, "row: %d, expected error: %w, actual error: %w", i, row.errExpected, err)
			continue
		}

		require.NoError(t, err, "row: %d, error must be nil", i)
		require.Equal(t, row.name_list, actualNames, "row: %d, expected names: %s, actual names: %s", i, row.name_list, actualNames)
	}
}

func TestGetAddresses(t *testing.T) {
	t.Parallel()

	mockWifi := NewWiFi(t)
	wifiService := myWiFi.Service{WiFi: mockWifi}

	for i, row := range testTable {
		mockWifi.On("Interfaces").Unset()
		mockWifi.On("Interfaces").Return(mockIfaces(row.addrs), row.errExpected)
		actualAddrs, err := wifiService.GetAddresses()
		if row.errExpected != nil {
			require.ErrorIs(t, err, row.errExpected, "row: %d, expected error: %w, actual error: %w", i, row.errExpected, err)
			continue
		}

		require.NoError(t, err, "row: %d, error must be nil", i)
		require.Equal(t, parseMACs(row.addrs), actualAddrs, "row: %d, expected addrs: %s, actual addrs: %s", i, parseMACs(row.addrs), actualAddrs)
	}
}

func mockIfaces(namesOrAddrs []string) []*wifi.Interface {
	var interfaces []*wifi.Interface

	for i, val := range namesOrAddrs {
		hwAddr := parseMAC(val)
		if hwAddr == nil && val != "" {
			hwAddr = net.HardwareAddr{}
		}

		iface := &wifi.Interface{
			Index:        i + 1,
			Name:         val,
			HardwareAddr: hwAddr,
			PHY:          1,
			Device:       1,
			Type:         wifi.InterfaceTypeAPVLAN,
			Frequency:    0,
		}

		interfaces = append(interfaces, iface)
	}
	return interfaces
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
