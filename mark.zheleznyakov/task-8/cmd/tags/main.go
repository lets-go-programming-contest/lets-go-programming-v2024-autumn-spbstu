package main

import (
  "fmt"
  "os"

	"github.com/mrqiz/task-8/internal/feats"
	"github.com/mrqiz/task-8/internal/output"
)

func main() {
  output.CliHeader()

  if (len(os.Args) < 4) {
    output.CliTooFewArgs()
    os.Exit(1)
  }

  featName := os.Args[1]
  if feats.FindByName(featName) == nil {
    output.CliOpNotFound()
    os.Exit(1)
  }

  fmt.Println(os.Args[2:])
}
