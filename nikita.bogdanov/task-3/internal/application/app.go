package application

import (
	"github.com/solomonalfred/task-3/internal/data_reader"
	"github.com/solomonalfred/task-3/internal/data_writer"
)

const (
	DataReaderError  = "Can`t read config file"
	ValuteCurseError = "Can`t get valute curse"
	OutputError      = "Smth wents wrong"
)

func App() {
	data, err := data_reader.GetConfig()
	if err != nil {
		panic(DataReaderError)
	}
	valute_curse, err := data_reader.GetValuteCurse(*data)
	if err != nil {
		panic(ValuteCurseError)
	}
	err = data_writer.GetJSONReport(*data, valute_curse)
	if err != nil {
		panic(OutputError)
	}
}
