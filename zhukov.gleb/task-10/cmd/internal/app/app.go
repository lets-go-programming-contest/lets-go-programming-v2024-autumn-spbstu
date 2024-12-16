package app

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"task-10/cmd/internal/config"
	gen "task-10/gen/proto/contact/v1"
	"task-10/internal/contact"
	"task-10/internal/db"
	"task-10/internal/grpc/server"
	"task-10/internal/http/handler"
	"task-10/internal/http/middleware"

	"github.com/bufbuild/protovalidate-go"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
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
	defer pgSQL.DB.Close()

	contactRepo := contact.NewContactRepo(&pgSQL)
	grpcServ, listener, err := a.setupGRPCServer(contactRepo)
	if err != nil {
		return err
	}

	go func() {
		if err = grpcServ.Serve(listener); err != nil {
			panic("gRPC server failed: " + err.Error())
		}
	}()

	httpHandler := a.setupHTTPHandlers(contactRepo)
	httpAddr := fmt.Sprintf(":%s", a.Cfg.RESTPort)
	httpServer := &http.Server{
		Addr:    httpAddr,
		Handler: httpHandler,
	}

	return httpServer.ListenAndServe()
}

func (a *App) setupHTTPHandlers(contactRepo *contact.ContactRepository) http.Handler {
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

	return h
}

func (a *App) setupGRPCServer(contactRepo *contact.ContactRepository) (*grpc.Server, net.Listener, error) {
	grpcServer := grpc.NewServer()
	validator, err := protovalidate.New()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to init validator %v", err)
	}

	grpcAddr := fmt.Sprintf(":%s", a.Cfg.GRPCPort)

	contactServer := server.NewContactServer(contactRepo, validator)
	gen.RegisterContactServiceServer(grpcServer, contactServer)

	listener, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to listen on port %s: %v", a.Cfg.GRPCPort, err)
	}

	return grpcServer, listener, nil
}
