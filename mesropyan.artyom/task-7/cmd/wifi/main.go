package main

import (
	"fmt"

	myWifi "github.com/artem6554/task-7/internal/wifi"
	"github.com/mdlayher/wifi"
)

func main() {
	wifiClient, err := wifi.New()
	if err != nil {
		fmt.Printf("Ошибка при создании wifiClient: %s\n", err.Error())

		return
	}

	wifiService := myWifi.New(wifiClient)

	addrs, _ := wifiService.GetAddresses()

	for _, addr := range addrs {
		fmt.Println(addr)
	}
}
