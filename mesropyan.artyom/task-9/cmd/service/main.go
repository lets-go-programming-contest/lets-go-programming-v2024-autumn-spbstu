package main

import (
	"fmt"
	"net/http"

	"github.com/artem6554/task-9/internal/config"
	handler "github.com/artem6554/task-9/internal/handler/http"
	"github.com/artem6554/task-9/internal/service"
	"github.com/gorilla/mux"
)

func main() {
	conf := config.ReadServerConfig()
	var service service.Service
	r := mux.NewRouter()
	handler := handler.NewHandler(service, r)
	server := http.Server{
		Addr:    conf.Addr,
		Handler: handler,
	}
	fmt.Println(server.ListenAndServe())
}
