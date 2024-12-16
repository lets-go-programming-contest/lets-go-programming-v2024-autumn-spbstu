package app

import (
	"log"
	"net/http"

	"erdem.istaev/task-9/cmd/server/internal/config"
	handler "erdem.istaev/task-9/internal/handler/http"
	"erdem.istaev/task-9/internal/repository/database"
	"erdem.istaev/task-9/internal/service"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

type app struct {
	httpServer *http.Server
}

func NewApp(cfg config.Config) (*app, *sqlx.DB) {
	log.Printf("create new database config")
	db, err := database.NewDBConfig(cfg.DB)
	if err != nil {
		panic(err)
	}

	log.Printf("create new database repository")
	dbRepository := database.NewRepository(db)

	log.Printf("create new service")
	newService := service.NewService(dbRepository)

	log.Printf("create new http server")
	r := mux.NewRouter()

	r.Use(handler.LogMiddleware)
	newHandler := handler.NewHandler(newService, r)
	server := &http.Server{
		Handler: newHandler,
		Addr:    ":" + cfg.ServerPort,
	}

	return &app{
		httpServer: server,
	}, db
}

func (s *app) Run() {
	log.Printf("start http server")

	if err := s.httpServer.ListenAndServe(); err != nil {
		panic(err)
	}
}

func (s *app) Close() error {
	return s.httpServer.Close()
}
