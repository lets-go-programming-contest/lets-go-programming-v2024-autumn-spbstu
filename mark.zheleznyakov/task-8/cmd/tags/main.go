package main

import (
  "fmt"
  "os"
  "strconv"

	"github.com/mrqiz/task-8/internal/feats"
	"github.com/mrqiz/task-8/internal/output"
)

func main() {
  output.CliHeader()

  if (len(os.Args) < 4) {
    output.CliTooFewArgs()
    os.Exit(1)
  }

  feat := feats.FindByName(os.Args[1])
  if feat == nil {
    output.CliOpNotFound()
    os.Exit(1)
  }

  args := make([]float64, len(os.Args[2:]))

  for i, str := range os.Args[2:] {
		if value, err := strconv.ParseFloat(str, 64); err == nil {
			args[i] = value
		} else {
			fmt.Printf("the '%s' arg is not... you know... convertible 🚗\n", str)
		}
	}

  fmt.Println(feat.Exec(args...))
}
