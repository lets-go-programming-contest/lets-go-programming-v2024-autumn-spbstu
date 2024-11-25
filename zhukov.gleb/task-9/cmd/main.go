package main

import (
	"net/http"
	"os"
	"strconv"
	"strings"
	"task-9/internal/contact"
	"task-9/internal/handler"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

const (
	envPath = ".env"
)

type Config struct {
	db db.DBcfg
}

// я ж запушу энв чтоб вам проще было запускать да xDDD
func loadConfig() (Config, error) {
	err := godotenv.Load(envPath)
	if err != nil {
		panic("error loading .env file: " + err.Error())
	}

	jwtSecret := os.Getenv("SESSION_TOKEN")

	dbConfigs := db.DBcfg{
		UPass:         os.Getenv("USER_PASS"),
		DBName:        os.Getenv("DB_NAME"),
		UName:         os.Getenv("USER_NAME"),
		MySQLHost:     os.Getenv("MYSQL_HOST"),
		MongoHost:     os.Getenv("MONGO_HOST"),
		PgSQLHost:     os.Getenv("PGSQL_HOST"),
		CollMongoName: os.Getenv("COLL_MONGO_NAME"),
	}

	dbConfigs.PortMySQL, err = strconv.Atoi(os.Getenv("PORT_MYSQL"))
	if err != nil {
		return Config{}, err
	}

	dbConfigs.PortPgSQL, err = strconv.Atoi(os.Getenv("PORT_PGSQL"))
	if err != nil {
		return Config{}, err
	}
	dbConfigs.PortMongo, err = strconv.Atoi(os.Getenv("PORT_MONGO"))
	if err != nil {
		return Config{}, err
	}

	return Config{
		db: dbConfigs,
	}, nil
}

func main() {
	config, err := loadConfig()
	if err != nil {
		panic("error reading data .env file: " + err.Error())
	}

	pgSQL, err := user.NewPgSQLController(
		config.db.UName,
		config.db.UPass,
		config.db.PgSQLHost,
		config.db.DBName,
		config.db.PortPgSQL,
	)
	if err != nil {
		panic("error with pgSQL DB: " + err.Error())
	}

	contactRepo := contact.NewContactRepo(pgSQL)
	contactHandler := &handler.ContactHandler{
		ContactRepo: contactRepo,
	}

	r := mux.NewRouter()
	r.HandleFunc("/contacts", contactHandler.Register).Methods("GET")
	r.HandleFunc("/contacts/{id}", contactHandler.Login).Methods("GET")

	staticHandler := http.StripPrefix(
		"/static",
		http.FileServer(http.Dir("./static")),
	)
	r.PathPrefix("/static/").Handler(staticHandler)

	r.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/api/") || strings.HasPrefix(r.URL.Path, "/static/") {
			http.NotFound(w, r)
			return
		}

		err = tmpl.ExecuteTemplate(w, "index.html", nil)
		if err != nil {
			logger.Errorw("Template execution error", "error", err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	})

	hndlr := middleware.Auth(logger, sm, r)
	hndlr = middleware.AccessLog(logger, hndlr)
	hndlr = middleware.PanicHandler(logger, hndlr)

	addr := ":8080"
	logger.Infow("starting server",
		"type", "START",
		"addr", addr,
	)

	err = http.ListenAndServe(addr, hndlr)
	if err != nil {
		panic("ListenAndServe fail" + err.Error())
	}
}
