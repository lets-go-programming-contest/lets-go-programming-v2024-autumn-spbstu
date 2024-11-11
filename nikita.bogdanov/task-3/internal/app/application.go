package app

import (
	"github.com/solomonalfred/task-3/internal/src"
)

func App() {
	data := src.ParseConfig()
	valute_curse := src.GetData(data)
	src.GetJson(data, valute_curse)
}
