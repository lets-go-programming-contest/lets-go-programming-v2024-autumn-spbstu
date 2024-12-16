package app

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/EmptyInsid/task-9/internal/config"
	"github.com/EmptyInsid/task-9/internal/database"
	handlers "github.com/EmptyInsid/task-9/internal/handlers/http"
	service "github.com/EmptyInsid/task-9/internal/service/database"
)

type MyApp struct {
	server *http.Server
	db     *database.Database
}

func NewApp(logger *slog.Logger, cfg config.Config) (*MyApp, error) {
	db, err := database.NewDB(&cfg.DBCfg)
	if err != nil {
		logger.Error("connect with db", "error", err)

		return nil, err
	}

	logger.Info("connect with db")

	dbService := service.NewDBService(db, logger)
	logger.Info("create new db service")

	router := mux.NewRouter()
	router.Use(handlers.LoggingMiddleware(logger))
	dbHandler := handlers.NewHandler(&dbService, router)

	logger.Info("create router")

	server := &http.Server{
		Addr:    cfg.ServerCfg.Port,
		Handler: dbHandler,
	}

	return &MyApp{
		server: server,
		db:     db,
	}, nil
}

func (a *MyApp) Run() error {
	log.Println("Run app")

	defer a.server.Close()

	defer a.db.Close()

	return fmt.Errorf("run app err: %w", a.server.ListenAndServe())
}
