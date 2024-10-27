package app

import (
	"fmt"
	"github.com/sssidkn/task-3/config"
	"github.com/sssidkn/task-3/internal/parser"
	"sort"
)

func Run(config *config.Config) {
	curs := parser.ParseFile(config.InputFile)
	sort.Sort(curs)

	for i := range curs.Valute {
		fmt.Println(curs.Valute[i])
	}
}
