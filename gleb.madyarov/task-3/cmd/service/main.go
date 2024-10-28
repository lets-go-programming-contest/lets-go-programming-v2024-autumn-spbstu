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
	var config read.Config = read.ReadConfig()
	xmltostruct.XMLtoStruct(&data, config)
	sort.Sorting(data.ValCurs)
	datajson := structtojson.StructtoJSON(&data)
	write.WriteResult(datajson, config)
}
