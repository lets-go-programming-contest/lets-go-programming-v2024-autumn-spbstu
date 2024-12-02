package application

import (
	"fmt"

	"github.com/solomonalfred/task-3/internal/data_reader"
	"github.com/solomonalfred/task-3/internal/data_writer"
)

const (
	DataReaderError  = "Can`t read config file %e"
	ValuteCurseError = "Can`t get valute curse %e"
	OutputError      = "Smth wents wrong %e"
)

func App() error {
	data, err := data_reader.GetConfig()
	if err != nil {
		return fmt.Errorf(DataReaderError, err)
	}
	valute_curse, err := data_reader.GetValuteCurse(*data)
	if err != nil {
		return fmt.Errorf(ValuteCurseError, err)
	}
	err = data_writer.GetJSONReport(*data, valute_curse)
	if err != nil {
		return fmt.Errorf(OutputError, err)
	}
	return nil
}
