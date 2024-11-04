package controller

import (
	"fmt"

	"task-3/internal/fileWorkers"
)

type Controller struct {
	filesWorker fileWorkers.FilesWorker
}

func NewController(w fileWorkers.FilesWorker) *Controller {
	return &Controller{
		filesWorker: w,
	}
}

func (c *Controller) Run() error {
	inOutFilePaths, err := c.filesWorker.CfgParse()
	if err != nil {
		return fmt.Errorf("Error in parsing cfg file: %v", err)
	}

	data, err := c.filesWorker.InputFileParse(inOutFilePaths.InFile)
	if err != nil {
		return fmt.Errorf("Error in parsing input file: %v", err)
	}

	err = c.filesWorker.OutputFileParse(inOutFilePaths.OutFile, data)
	if err != nil {
		return fmt.Errorf("Error in output data in file: %v", err)
	}

	return nil
}
