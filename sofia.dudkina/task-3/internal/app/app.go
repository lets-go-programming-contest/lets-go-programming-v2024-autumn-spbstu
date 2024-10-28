package app

import (
	"sort"

	"github.com/sssidkn/task-3/config"
	"github.com/sssidkn/task-3/internal/parser"
	"github.com/sssidkn/task-3/internal/writer"
)

func Run(config *config.Config) error {
	curs, err := parser.ParseFile(config.InputFile)
	if err != nil {
		return err
	}
	sort.Sort(curs)
	err = writer.WriteFile(config.OutputFile, curs)
	if err != nil {
		return err
	}
	return nil
}
