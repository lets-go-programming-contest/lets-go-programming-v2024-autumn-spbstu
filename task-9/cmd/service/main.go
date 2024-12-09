package main

import (
	config "contactManager/internal/config"
	"contactManager/internal/dbase"
	"contactManager/internal/manager"

	// "fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	// fmt.Println("hello world!")

	config, err := config.ReadConfig()
	if err != nil {
		panic(err)
	}

	db, err := dbase.ConnectToDB(config)
	if err != nil {
		panic(err)
	}

	manager.InitDataBase(db)
	router := mux.NewRouter()
	manager.CreateRoutes(router)
	port := config.Server.Port

	headers := handlers.AllowedHeaders([]string{"Content-Type"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"http://localhost" + strconv.Itoa(port)})
	loggedRouter := handlers.CORS(headers, methods, origins)(router)

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), loggedRouter))

	err = dbase.CloseDB(db)
	if err != nil {
		panic(err)
	}
}
