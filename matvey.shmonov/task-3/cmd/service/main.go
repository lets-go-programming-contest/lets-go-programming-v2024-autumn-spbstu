package main

import (
	"log"

	"github.com/Koshsky/task-3/internal/configutil"
	"github.com/Koshsky/task-3/internal/currencies"
)

func main() {
	cfg, err := configutil.ReadConfig()
	if err != nil {
		log.Panicf("err: %v", err)
	}

	crcs := currencies.Currencies{}
	err = currencies.UnmarshalXML(&crcs, cfg.InFilename)
	if err != nil {
		log.Panicf("err: %v", err)
	}
	currencies.Sort(crcs.List)
	err = currencies.MarshalJSON(&crcs, cfg.OutFilename)
	if err != nil {
		log.Panicf("err: %v", err)
	}
}
