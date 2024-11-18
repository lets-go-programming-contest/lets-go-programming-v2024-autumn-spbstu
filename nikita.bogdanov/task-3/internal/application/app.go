package application

import (
	"github.com/solomonalfred/task-3/internal/data_reader"
	"github.com/solomonalfred/task-3/internal/data_writer"
)

func App() {
	data := data_reader.GetConfig()
	valute_curse := data_reader.GetValuteCurse(data)
	data_writer.GetJSONReport(data, valute_curse)
}
