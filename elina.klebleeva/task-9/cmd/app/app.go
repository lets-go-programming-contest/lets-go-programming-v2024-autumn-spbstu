package app

import (
	"log"
	"log/slog"
	"net/http"

	"github.com/EmptyInsid/task-9/internal/config"
	"github.com/EmptyInsid/task-9/internal/database"
	handlers "github.com/EmptyInsid/task-9/internal/handlers/http"
	service "github.com/EmptyInsid/task-9/internal/service/database"
	"github.com/gorilla/mux"
)

type MyApp struct {
	server *http.Server
	db     *database.Database
}

func NewApp(logger *slog.Logger, cfg config.Config) (*MyApp, error) {
	//init db
	db, err := database.NewDb(&cfg.DBCfg)
	if err != nil {
		logger.Error("connect with db", "error", err)
		return nil, err
	}
	logger.Info("connect with db")

	dbService := service.NewDbService(db, logger)
	logger.Info("create new db service")

	//init router
	router := mux.NewRouter()
	router.Use(handlers.LoggingMiddleware(logger))
	dbHandler := handlers.NewHandler(&dbService, router)

	logger.Info("create router")

	//init server
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
	defer a.db.Close()
	defer a.server.Close()
	return a.server.ListenAndServe()
}
