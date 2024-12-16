package main

import (
	"fmt"

	"github.com/mdlayher/wifi"

	myWifi "github.com/EmptyInsid/task-5/internal/wifi"
)

func main() {
	wifiClient, err := wifi.New()
	if err != nil {
		fmt.Printf("Error while create wifiClient: %s\n", err.Error())

		return
	}

	wifiService := myWifi.New(wifiClient)

	addrs, err := wifiService.GetAddresses()
	if err != nil {
		fmt.Printf("Error getting address: %s\n", err.Error())

		return
	}

	for _, addr := range addrs {
		fmt.Println(addr)
	}
}
