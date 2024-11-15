package wifi_test

import (
	"fmt"
	"net"
	"testing"

	myWiFi "github.com/katagiriwhy/task-7/internal/wifi"
	"github.com/mdlayher/wifi"
	"github.com/stretchr/testify/require"
)

type testWifi struct {
	addrs       []string
	errExpected error
}

var testTable = []testWifi{
	{
		addrs:       nil, //[]string{"1111:2222:3333:4444", "8800.0000.8800:9999"}
		errExpected: nil,
	},
	{
		addrs:       nil,
		errExpected: nil, //errors.New("no addresses provided")
	},
}

func mockIfaces(addrs []string) []*wifi.Interface {
	var ifaces []*wifi.Interface
	for i, addr := range addrs {
		hardwareAddr, err := net.ParseMAC(addr)
		if err != nil {
			continue
		}
		intrface := &wifi.Interface{
			Index:        i,
			Name:         fmt.Sprintf("eth%d", i),
			HardwareAddr: hardwareAddr,
			PHY:          i,
			Device:       i,
			Type:         wifi.InterfaceTypeAPVLAN,
			Frequency:    0,
		}
		ifaces = append(ifaces, intrface)
	}
	return ifaces
}

func parseMACs(str []string) []net.HardwareAddr {
	var addrs []net.HardwareAddr
	for _, address := range str {
		tmp, err := net.ParseMAC(address)
		if err != nil {
			return nil
		}
		addrs = append(addrs, tmp)
	}
	return addrs
}

func TestGetAddresses(t *testing.T) {
	t.Parallel()

	mockWifi := NewWiFi(t)
	wifiService := myWiFi.Service{WiFi: mockWifi}
	for i, row := range testTable {
		mockWifi.On("Interfaces").Unset()
		mockWifi.On("Interfaces").Return(mockIfaces(row.addrs), row.errExpected)
		ourAddrs, err := wifiService.GetAddresses()
		if row.errExpected == nil {
			require.ErrorIs(t, err, row.errExpected, "row %d, expected error: %w, actual error: %w", i, row.errExpected, err)
			continue
		}
		require.NoError(t, err, "row %d, error must be nil", i)
		require.Equal(t, parseMACs(row.addrs), ourAddrs, "row: %d, expected addrs: %s, actual addrs: %s", i, parseMACs(row.addrs), ourAddrs)
	}
}

func TestGetNames(t *testing.T) {
	t.Parallel()

	mockWifi := NewWiFi(t)
	wifiService := myWiFi.Service{WiFi: mockWifi}
	for i, row := range testTable {
		mockWifi.On("Interfaces").Unset()
		mockWifi.On("Interfaces").Return(mockIfaces(row.addrs), row.errExpected)
		ourNames, err := wifiService.GetNames()
		if row.errExpected == nil {
			require.ErrorIs(t, err, row.errExpected, "row %d, expected error: %w, actual error: %w", i, row.errExpected, err)
			continue
		}
		require.NoError(t, err, "row %d, error must be nil", i)
		require.Equal(t, []string{"eth0", "eth1"}, ourNames, "row: %d, expected names: %s, actual names: %s", i, row.addrs, ourNames)
	}
}
