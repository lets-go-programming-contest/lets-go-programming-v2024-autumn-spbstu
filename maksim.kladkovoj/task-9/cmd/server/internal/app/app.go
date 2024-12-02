package app

import (
	"fmt"
	"net/http"

	"github.com/Mmmakskl/task-9/cmd/server/internal/config"
	database "github.com/Mmmakskl/task-9/internal/database/file"
	handler "github.com/Mmmakskl/task-9/internal/handler/http"
	service "github.com/Mmmakskl/task-9/internal/service/file"
	"github.com/gorilla/mux"
)

type App struct {
	server *http.Server
}

func NewApp(cfg *config.Config) (*App, error) {
	readDB, err := database.NewDatabaseReader(database.DBstruct(cfg.Database))
	if err != nil {
		return nil, fmt.Errorf("error while create reader database: %w", err)
	}

	writeDB, err := database.NewDatabaseWriter(database.DBstruct(cfg.Database))
	if err != nil {
		return nil, fmt.Errorf("error while create writer database: %w", err)
	}

	contactService := service.NewService(readDB, writeDB)

	r := mux.NewRouter()

	contactHandler := handler.NewHandler(&contactService, r)

	server := &http.Server{
		Addr:    cfg.Server.Host,
		Handler: contactHandler,
	}

	return &App{
		server: server,
	}, nil
}

func (a *App) Run() error {
	defer a.server.Close()
	return a.server.ListenAndServe()
}
