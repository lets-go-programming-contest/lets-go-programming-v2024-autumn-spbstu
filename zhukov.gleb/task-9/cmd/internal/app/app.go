package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"task-9/cmd/internal/config"
	"task-9/internal/contact"
	"task-9/internal/db"
	"task-9/internal/http/handler"
	"task-9/internal/http/middleware"

	"github.com/gorilla/mux"
)

type App struct {
	Cfg config.AppCfg
}

func (a *App) Run() error {
	var err error
	a.Cfg, err = config.LoadConfig()
	if err != nil {
		return fmt.Errorf("error with config.LoadConfig(): %w", err)
	}

	pgSQL, err := db.NewPgSQLController(
		a.Cfg.DBCfg,
	)
	if err != nil {
		return fmt.Errorf("error with db.NewPgSQLController: %w", err)
	}
	//TODO defer close
	defer pgSQL.DB.Close()

	contactRepo := contact.NewContactRepo(&pgSQL)
	contactHandler := &handler.ContactHandler{
		ContactRepo: contactRepo,
	}

	r := mux.NewRouter()
	r.HandleFunc("/contacts", contactHandler.GetAll).Methods("GET")
	r.HandleFunc("/contacts/{id}", contactHandler.GetByID).Methods("GET")
	r.HandleFunc("/contacts", contactHandler.AddContact).Methods("POST")
	r.HandleFunc("/contacts/{id}", contactHandler.UpdateContact).Methods("PUT")
	r.HandleFunc("/contacts/{id}", contactHandler.DeleteContact).Methods("DELETE")

	logger := log.New(os.Stdout, "LOG:", log.LUTC)
	h := middleware.LoggingMiddleware(logger, r)
	h = middleware.PanicHandler(logger, h)

	addr := strings.Builder{}
	addr.WriteString(":")
	addr.WriteString(a.Cfg.Port)
	//TODO full address - зачем
	return http.ListenAndServe(addr.String(), h)
}
