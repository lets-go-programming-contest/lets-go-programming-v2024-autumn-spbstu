package wifi_test

import (
	"errors"
	"fmt"
	"net"
	"testing"

	"github.com/mdlayher/wifi"
	"github.com/stretchr/testify/require"

	myWifi "erdem.istaev/task-6/internal/wifi"
)

//go:generate mockery --all --testonly --quiet --outpkg wifi_test --output .

type rowTestSysInfo struct {
	addrs       []string
	names       []string
	errExpected error
}

var testTable = []rowTestSysInfo{
	{
		addrs:       []string{"00:11:22:33:44:55", "aa:bb:cc:dd:ee:ff"},
		names:       []string{"eth0", "eth1"},
		errExpected: nil,
	},
	{
		addrs:       []string{"12:34:56:78:90:AB"},
		names:       []string{"eth0"},
		errExpected: nil,
	},
	{
		addrs:       []string{},
		names:       []string{},
		errExpected: errors.New("ExpectedError"),
	},
	{
		addrs:       []string{"invalid_mac"},
		names:       []string{"eth3"},
		errExpected: errors.New("ExpectedError"),
	},
	{
		addrs:       nil,
		names:       []string{"eth0"},
		errExpected: errors.New("ExpectedError"),
	},
	{
		addrs:       []string{"aa:bb:cc:dd:ee:ff"},
		names:       nil,
		errExpected: errors.New("ExpectedError"),
	},
}

func TestGetAddresses(t *testing.T) {
	for i, row := range testTable {
		t.Run("", func(t *testing.T) {
			t.Parallel()

			mockWifi := NewWiFi(t)
			wifiService := myWifi.Service{WiFi: mockWifi}

			mockWifi.On("Interfaces").Unset()
			mockWifi.On("Interfaces").Return(mockIfaces(row.addrs), row.errExpected)

			actualAddrs, err := wifiService.GetAddresses()
			if row.errExpected != nil {
				require.ErrorIs(t, err, row.errExpected,
					"row: %d, expected error: %w, actual error: %w",
					i, row.errExpected, err)
				return
			}

			require.NoError(t, err, "row: %d, error must be nil", i)
			require.Equal(t, parseMACs(row.addrs), actualAddrs,
				"row: %d, expected addrs: %s, actual addrs: %s",
				i, parseMACs(row.addrs), actualAddrs)
		})
	}
}

func mockIfaces(addrs []string) []*wifi.Interface {
	var interfaces []*wifi.Interface
	for i, addrStr := range addrs {
		hwAddr := parseMAC(addrStr)
		if hwAddr == nil {
			continue
		}

		iface := &wifi.Interface{
			Index:        i + 1,
			Name:         fmt.Sprintf("eth%d", i),
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

func TestGetNames(t *testing.T) {
	for i, row := range testTable {
		t.Run(fmt.Sprintf("TestCase%d", i), func(t *testing.T) {
			t.Parallel()

			mockWifi := NewWiFi(t)
			wifiService := myWifi.Service{WiFi: mockWifi}

			mockWifi.On("Interfaces").Unset()
			mockWifi.On("Interfaces").Return(mockIfaces(row.addrs), row.errExpected)

			actualNames, err := wifiService.GetNames()
			if row.errExpected != nil {
				require.ErrorIs(t, err, row.errExpected,
					"row: %d, expected error: %w, actual error: %w",
					i, row.errExpected, err)
				return
			}

			require.NoError(t, err, "row: %d, error must be nil", i)
			require.Equal(t, row.names, actualNames,
				"row: %d, expected names: %s, actual names: %s",
				i, row.names, actualNames)
		})
	}
}
