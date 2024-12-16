package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hahapathetic/task-9/internal/config"
	"github.com/hahapathetic/task-9/internal/handler"
	"github.com/hahapathetic/task-9/internal/service"
)

func main() {
	conf := config.ReadServerConfig()

	svc := service.Service{}

	r := mux.NewRouter()

	handler.NewHandler(svc, r)

	server := &http.Server{
		Addr:    conf.Addr,
		Handler: r,
	}

	fmt.Printf("Server started at %s\n", conf.Addr)
	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("Server failed: %v\n", err)
	}
}
