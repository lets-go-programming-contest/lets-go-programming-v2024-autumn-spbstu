package main

import (
	"fmt"
	"net/http"

	handler "github.com/artem6554/task-9/internal/handler/http"
	"github.com/artem6554/task-9/internal/service"
	"github.com/gorilla/mux"
)

func main() {
	var service service.Service
	r := mux.NewRouter()
	handler := handler.NewHandler(service, r)
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}
	fmt.Println(server.ListenAndServe())
}
