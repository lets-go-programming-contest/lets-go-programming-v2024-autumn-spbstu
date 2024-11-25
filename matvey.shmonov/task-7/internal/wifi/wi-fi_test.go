package wifi_test

import (
	"errors"
	"net"
	"testing"

	myWifi "github.com/Koshsky/task-6/internal/wifi"
	"github.com/mdlayher/wifi"
	"github.com/stretchr/testify/require"
)

type rowTestSysInfo struct {
	addrs       []string
	name_list   []string
	errExpected error
}

var testTable = []rowTestSysInfo{
	{
		addrs:       []string{"01:23:45:67:89:ab", "cd:ef:01:23:45:67"},
		name_list:   []string{"Alice", "Bob"},
		errExpected: nil,
	},
	{
		addrs:       []string{"01:23:45:67:89:ab", "cd:ef:01:23:45:67"},
		name_list:   nil,
		errExpected: errors.New("ExpectedError"),
	},
	{
		addrs:       nil,
		name_list:   []string{"Alice", "Bob"},
		errExpected: errors.New("ExpectedError"),
	},
	{
		addrs:       nil,
		name_list:   nil,
		errExpected: errors.New("ExpectedError"),
	},
}

func TestGetAddresses(t *testing.T) {
	t.Parallel()
	mockWiFi := NewWiFi(t)
	wifiService := myWifi.New(mockWiFi)

	for i, row := range testTable {
		mockWiFi.On("Interfaces").Unset()
		mockWiFi.On("Interfaces").Return(mockInterfaces(row), row.errExpected)
		actualAddrs, err := wifiService.GetAddresses()

		if row.errExpected != nil {
			require.ErrorIs(t, err, row.errExpected, "row: %d, expected error: %w, actual error: %w", i, row.errExpected, err)
			continue
		}

		require.NoError(t, err, "row: %d, error must be nil", i)
		require.Equal(t, parseMACs(row.addrs), actualAddrs)
	}
}

func TestGetNames(t *testing.T) {
	t.Parallel()
	mockWiFi := NewWiFi(t)
	wifiService := myWifi.New(mockWiFi)

	for i, row := range testTable {
		mockWiFi.On("Interfaces").Unset()
		mockWiFi.On("Interfaces").Return(mockInterfaces(row), row.errExpected)
		actualNames, err := wifiService.GetNames()

		if row.errExpected != nil {
			require.ErrorIs(t, err, row.errExpected, "row: %d, expected error: %w, actual error: %w", i, row.errExpected, err)
			continue
		}

		require.NoError(t, err, "row: %d, error must be nil", i)
		require.Equal(t, row.name_list, actualNames)
	}
}

func parseMACs(addrs []string) []net.HardwareAddr {
	var hardwareAddrs []net.HardwareAddr
	for _, addr := range addrs {
		hwAddr, _ := net.ParseMAC(addr)
		hardwareAddrs = append(hardwareAddrs, hwAddr)
	}
	return hardwareAddrs
}

func mockInterfaces(row rowTestSysInfo) []*wifi.Interface {
	if len(row.addrs) != len(row.name_list) {
		return nil
	}

	mockInterfaces := make([]*wifi.Interface, len(row.addrs))
	for i, addr := range row.addrs {
		hwAddr, _ := net.ParseMAC(addr)
		mockInterfaces[i] = &wifi.Interface{
			Name:         row.name_list[i],
			HardwareAddr: hwAddr,
		}
	}
	return mockInterfaces
}
