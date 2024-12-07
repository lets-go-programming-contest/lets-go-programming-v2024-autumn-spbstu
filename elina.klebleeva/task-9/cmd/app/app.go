package app

import (
	"net/http"

	"github.com/EmptyInsid/task-9/internal/config"
)

type MyApp struct {
	server *http.Server
}

func NewApp(cfg config.Config) *MyApp {
	//init db

	//init router

	//init handlers

	//init server

	return &MyApp{}
}

func (a *MyApp) Run() {

}

func (a *MyApp) Close() {

}
