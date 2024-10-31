package main

import (
	xmltostruct "github.com/Madyarov-Gleb/task-3/internal/XMLtoStruct"
	"github.com/Madyarov-Gleb/task-3/internal/read"
	"github.com/Madyarov-Gleb/task-3/internal/sort"
	structtojson "github.com/Madyarov-Gleb/task-3/internal/structtoJSON"
	"github.com/Madyarov-Gleb/task-3/internal/write"
)

func main() {
	data := xmltostruct.ValCurs{}
	var config read.Config
	config, err := read.ReadConfig()
	if err != nil {
		panic(err)
	}
	data, err = xmltostruct.XMLtoStruct(&data, config)
	if err != nil {
		panic(err)
	}
	sort.Sorting(data.ValCurs)
	datajson, err := structtojson.StructtoJSON(&data)
	if err != nil {
		panic(err)
	}
	err = write.WriteResult(datajson, config)
	if err != nil {
		panic(err)
	}
}
