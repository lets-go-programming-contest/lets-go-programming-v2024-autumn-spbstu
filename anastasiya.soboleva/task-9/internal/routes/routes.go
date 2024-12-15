package routes

import (
	"database/sql"
	"github.com/nayzzerr/task-9/internal/handlers"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nayzzerr/task-9/internal/repository"
	"github.com/nayzzerr/task-9/internal/services"
)

func SetupRoutes(db *sql.DB) *mux.Router {
	contactRepo := repository.NewContactRepository(db)
	contactService := services.NewContactService(contactRepo)
	contactHandler := handlers.NewContactHandler(contactService)

	router := mux.NewRouter()
	router.HandleFunc("/contacts", contactHandler.GetAll).Methods(http.MethodGet)
	router.HandleFunc("/contacts/{id}", contactHandler.GetByID).Methods(http.MethodGet)
	router.HandleFunc("/contacts", contactHandler.Create).Methods(http.MethodPost)
	router.HandleFunc("/contacts/{id}", contactHandler.Update).Methods(http.MethodPut)
	router.HandleFunc("/contacts/{id}", contactHandler.Delete).Methods(http.MethodDelete)

	return router
}
