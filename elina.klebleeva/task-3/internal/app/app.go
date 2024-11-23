package app

import (
	"fmt"
	"sort"

	"github.com/EmptyInsid/task-3/configs/configData"
	"github.com/EmptyInsid/task-3/internal/errorUtils"
	"github.com/EmptyInsid/task-3/internal/parseXml"
	"github.com/EmptyInsid/task-3/internal/writeJson"
)

func Run(config *configData.Config) error {

	if config == nil {
		return errorUtils.ErrorWithLocation(fmt.Errorf("configData is nil"))
	}

	xmlData, err := parseXml.ProcessXml(config.InputFile)
	if err != nil {
		return errorUtils.ErrorWithLocation(err)
	}

	sort.Sort(xmlData)

	err = writeJson.ProcessJson(xmlData, config.OutputFile)
	if err != nil {
		return errorUtils.ErrorWithLocation(err)
	}

	return nil
}
