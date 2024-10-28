package write

import (
	"fmt"
	"os"

	"github.com/Madyarov-Gleb/task-3/internal/read"
)

func WriteResult(data []byte, config read.Config) {
	file, err := os.OpenFile(config.Output, os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("the output file could not be found, the file will be created")
		file, err = os.Create("output.json")
		if err != nil {
			panic("failed to create a file")
		}
	}
	_, err = file.Write(data)
	if err != nil {
		panic("failed to write to a file")
	}
	defer file.Close()
}
