package app

import (
	"fmt"

	"github.com/KRYST4L614/task-9/internal/db"
	"github.com/KRYST4L614/task-9/internal/handler"
	"github.com/KRYST4L614/task-9/internal/repository"
	"github.com/KRYST4L614/task-9/internal/route"
	"github.com/KRYST4L614/task-9/internal/server"
	"github.com/KRYST4L614/task-9/internal/service"
	"github.com/gorilla/mux"
)

type App struct {
	server *server.Server
	config *Config
}

func NewApp(config *Config) (*App, error) {
	dbPool, err := db.NewDB(config.DbConfig, 3)
	if err != nil {
		return nil, fmt.Errorf("failed to connect db with error: %w", err)
	}

	router := mux.NewRouter()
	repository := repository.NewRepository(dbPool)
	service := service.NewContactService(repository)
	handler := handler.NewContactHandler(service)
	route.RegisterHandlers(router, *handler)

	server := server.NewServer(&config.ServerConfig, router)

	return &App{
		server: server,
		config: config,
	}, nil
}

func (app *App) Start() error {
	app.server.Start()
	return nil
}

func (app *App) Stop() error {
	var errComposition error
	if err := app.server.Stop(); err != nil {
		errComposition = fmt.Errorf("%w. %w", err, errComposition)
	}
	return errComposition
}
