package main

import (
	"fmt"

	"github.com/mdlayher/wifi"

	myWifi "github.com/IDevFrye/task-7/internal/wifi"
)

func main() {
	wifiClient, err := wifi.New()
	if err != nil {
		fmt.Printf("Ошибка при создании wifiClient: %s\n", err.Error())

		return
	}

	wifiService := myWifi.New(wifiClient)

	addrs, err := wifiService.GetAddresses()
	if err != nil {
		fmt.Printf("Ошибка при получении адресов: %s\n", err.Error())

		return
	}

	for _, addr := range addrs {
		fmt.Println(addr)
	}
}
