package main

import (
	"bufio"
	"os"

	"github.com/KRYST4L614/task-2-1/internal/io"
	"github.com/KRYST4L614/task-2-1/internal/temperature"
)

func main() {
	temperConstraints := temperature.TemperatureConstraints{
		Max: 30,
		Min: 15,
	}

	reader := bufio.NewReader(os.Stdin)

	n, err := io.ReadInt(reader)
	if err != nil {
		panic(err.Error())
	}

	for range n {
		err := temperature.TemperatureRequestHandler(temperConstraints, reader, bufio.NewWriter(os.Stdout))
		if err != nil {
			panic(err.Error())
		}
	}
}
