package app

import (
	"fmt"
	"net/http"
	"strings"

	"task-9/cmd/internal/config"
	"task-9/internal/contact"
	"task-9/internal/db"
	"task-9/internal/handler"

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
		a.Cfg.UDBName,
		a.Cfg.UDBPass,
		a.Cfg.PgSQLHost,
		a.Cfg.DBPgSQLName,
		a.Cfg.PortPgSQL,
	)
	if err != nil {
		return fmt.Errorf("error with db.NewPgSQLController: %w", err)
	}

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

	addr := strings.Builder{}
	addr.WriteString(":")
	addr.WriteString(a.Cfg.Host)

	return http.ListenAndServe(addr.String(), r)
}
