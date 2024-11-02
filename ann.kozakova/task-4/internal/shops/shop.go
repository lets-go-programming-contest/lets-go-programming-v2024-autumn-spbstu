package shops

import (
	"strconv"
	"sync"
)

type Shop struct {
	All   int
	Items []string
	mutex sync.Mutex
}

func Add100Items(shop *Shop, product int) {
	for i := 0; i < 100; i++ {
		shop.Items = append(shop.Items, "Product"+strconv.Itoa(product+i))
		shop.All += 1
	}
}
