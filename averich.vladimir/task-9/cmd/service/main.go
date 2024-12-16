package main

import (
	"log"
	"net/http"
	"strconv"
	"task-9/internal/config"
	"task-9/internal/db"
	"task-9/internal/handlers"

	handlersGorilla "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	cfg, err := config.ReadConfig()
	if err != nil {
		panic(err)
	}

	database, err := db.InitDB(cfg)
	if err != nil {
		panic(err)
	}

	router := mux.NewRouter()
	handlers.CreateRoutes(router)
	port := cfg.Server.Port

	headers := handlersGorilla.AllowedHeaders([]string{"Content-Type"})
	methods := handlersGorilla.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlersGorilla.AllowedOrigins([]string{"http://localhost" + strconv.Itoa(port)})
	loggedRouter := handlersGorilla.CORS(headers, methods, origins)(router)

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), loggedRouter))

	err = db.CloseDB(database)
	if err != nil {
		panic(err)
	}
}

/*
-- Создание базы данных
CREATE DATABASE Contacts;

-- Подключение к базе данных
\c Contacts;

-- Создание таблицы contacts
CREATE TABLE contacts (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    phone VARCHAR(20) NOT NULL
);

-- Добавление данных
INSERT INTO contacts (name, phone) VALUES
('Alice', '+1234567890'),
('Bob', '+0987654321'),
('Charlie', '+1122334455');

-- Просмотр данных
SELECT * FROM contacts;
*/
