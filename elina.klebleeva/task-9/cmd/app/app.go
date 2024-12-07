package app

import (
	"log"
	"net/http"

	"github.com/EmptyInsid/task-9/internal/config"
	"github.com/EmptyInsid/task-9/internal/database"
)

type MyApp struct {
	server *http.Server
}

func NewApp(cfg config.Config) (*MyApp, error) {
	//init db
	db, err := database.NewDb(&cfg.DBCfg)
	if err != nil {
		log.Printf("Error while connect with database: %v", err)
		return nil, err
	}
	log.Println("Succsess connect with db")

	//init router

	//init handlers

	//init server

	return &MyApp{}, nil
}

func (a *MyApp) Run() {

}

func (a *MyApp) Close() {

}
