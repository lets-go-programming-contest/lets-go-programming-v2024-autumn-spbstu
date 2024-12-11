package app

import (
	"log"
	"net/http"

	"github.com/EmptyInsid/task-9/internal/config"
	"github.com/EmptyInsid/task-9/internal/database"
	handlers "github.com/EmptyInsid/task-9/internal/handlers/http"
	service "github.com/EmptyInsid/task-9/internal/service/database"
	"github.com/gorilla/mux"
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
	defer db.Close()
	log.Println("Succsess connect with db")

	dbService := service.NewDbService(db)

	//init router
	router := mux.NewRouter()
	dbHandler := handlers.NewHandler(&dbService, router)

	//init server
	server := &http.Server{
		Addr:    cfg.ServerCfg.Port,
		Handler: dbHandler,
	}

	return &MyApp{server: server}, nil
}

func (a *MyApp) Run() error {
	defer a.server.Close()
	return a.server.ListenAndServe()
}
