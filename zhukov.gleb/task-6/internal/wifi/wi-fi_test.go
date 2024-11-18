package wifi_test

import (
	"errors"
	"net"
	"testing"

	"github.com/mdlayher/wifi"
	"github.com/stretchr/testify/require"

	myWifi "task-6/internal/wifi"
)

//go:generate mockery --all --testonly --quiet --outpkg wifi_test --output .
type rowTestSysInfo struct {
	addrs       []string
	names       []string
	errExpected error
}

func TestGetAddresses(t *testing.T) {
	t.Parallel()

	testTable := []rowTestSysInfo{
		{
			addrs: []string{"00:11:22:33:44:55", "aa:bb:cc:dd:ee:ff"},
		},
		{
			errExpected: errors.New("ExpectedError"),
		},
	}

	t.Run("whole GetAddresses test", func(t *testing.T) {
		mockWifi := NewWiFi(t)
		wifiService := myWifi.Service{WiFi: mockWifi}
		for i, row := range testTable {
			mockWifi.On("Interfaces").
				Unset()
			mockWifi.On("Interfaces").
				Return(mockIfaces(row.addrs), row.errExpected)

			actualAddrs, err := wifiService.GetAddresses()
			if row.errExpected != nil {
				require.Contains(t, err.Error(), row.errExpected.Error(), "row: %d, expected error: %v, actual error: %v",
					i,
					row.errExpected,
					err,
				)
				continue
			}

			require.NoError(t, err, "row: %d, error must be nil", i)
			require.Equal(t, parseMACs(row.addrs), actualAddrs,
				"row: %d, expected addrs: %s, actual addrs: %s", i,
				parseMACs(row.addrs), actualAddrs)
		}
	})
}

func TestGetNamess(t *testing.T) {
	t.Parallel()

	testTable := []rowTestSysInfo{
		{
			names: []string{"name"},
		},
		{
			addrs:       []string{"not_valid"},
			errExpected: errors.New("some err"),
		},
	}

	t.Run("whole GetAddresses test", func(t *testing.T) {
		mockWifi := NewWiFi(t)
		wifiService := myWifi.New(mockWifi)

		for i, row := range testTable {
			mockWifi.On("Interfaces").
				Unset()
			mockWifi.On("Interfaces").
				Return(mockIfaces(row.names), row.errExpected)

			actualNames, err := wifiService.GetNames()
			if row.errExpected != nil {
				require.Contains(t, err.Error(), row.errExpected.Error(), "row: %d, expected error: %v, actual error: %v",
					i,
					row.errExpected,
					err,
				)
				continue
			}

			require.NoError(t, err, "row: %d, error must be nil", i)
			require.Equal(t, row.names, actualNames, "row: %d, expected addrs: %s, actual addrs: %s",
				i,
				row.names,
				actualNames,
			)
		}
	})
}

func mockIfaces(data []string) []*wifi.Interface {
	var interfaces []*wifi.Interface

	for i, v := range data {
		hwAddr := parseMAC(v)

		iface := &wifi.Interface{
			Index:        i + 1,
			Name:         v,
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
