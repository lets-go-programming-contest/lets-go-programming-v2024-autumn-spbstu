package input

import (
	"fmt"
	"os"
)

func InputInt(x *int) {
	_, err := fmt.Scan(x)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func InputSlice(slice []int) {
	var temp int
	for i := 0; i < len(slice); i++ {
		InputInt(&temp)
		slice[i] = temp
	}
}
