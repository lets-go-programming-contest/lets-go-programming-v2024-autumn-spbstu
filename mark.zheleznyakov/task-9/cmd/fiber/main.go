package main

import (
	"github.com/gofiber/fiber/v2"
	"log"

	"github.com/mrqiz/task-9/internal/database"
	"github.com/mrqiz/task-9/internal/fiber/routes"
	"github.com/mrqiz/task-9/internal/models"
)

func init() {
	if err := database.Connect(); err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}

	if err := database.DB.AutoMigrate(&models.Contact{}); err != nil {
		log.Fatalf("migration failed: %v", err)
	}
}

func main() {
	app := fiber.New()
	contacts := app.Group("/contacts")
	fiberroutes.ContactsRouter(contacts)
	log.Fatal(app.Listen(":3000"))

	if db, err := database.DB.DB(); err == nil {
		defer db.Close()
	} else {
		log.Printf("failed to close connection: %v", err)
	}
}
