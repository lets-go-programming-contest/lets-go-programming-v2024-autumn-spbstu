package main

import (
	"bufio"
	"fmt"
	"os"

	temp "github.com/Koshsky/task-2-1/internal/temperature"
)

func main() {
	b := bufio.NewReader(os.Stdin)

	N := temp.ReadInt()
	for range N {
		minT := temp.MinTemperature
		maxT := temp.MaxTemperature
		K := temp.ReadInt()
		for range K {
			prompt := temp.ReadPrompt(b)
			if temp.ProcessInput(prompt, &minT, &maxT) {
				fmt.Println(minT)
			} else {
				fmt.Println(-1)
				break
			}
		}
	}
}
