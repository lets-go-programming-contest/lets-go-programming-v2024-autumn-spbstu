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
	db     *database.Database
}

func NewApp(cfg config.Config) (*MyApp, error) {
	//init db
	db, err := database.NewDb(&cfg.DBCfg)
	if err != nil {
		log.Printf("Error while connect with database :: %v", err)
		return nil, err
	}
	log.Println("Succsess connect with db")

	dbService := service.NewDbService(db)
	log.Println("Succsess create db service")

	//init router
	router := mux.NewRouter()
	dbHandler := handlers.NewHandler(&dbService, router)
	log.Println("Succsess create db service")

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
