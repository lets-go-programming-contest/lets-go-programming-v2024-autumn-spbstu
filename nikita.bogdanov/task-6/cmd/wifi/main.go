package main

import (
	"fmt"

	"github.com/mdlayher/wifi"

	myWifi "github.com/solomonalfred/task-6/internal/wifi"
)

func main() {
	wifiClient, _ := wifi.New()

	wifiService := myWifi.New(wifiClient)
	addrs, _ := wifiService.GetAddresses()

	for _, addr := range addrs {
		fmt.Println(addr)
	}
}
