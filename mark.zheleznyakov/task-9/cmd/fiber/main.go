package main

import (
	"github.com/gofiber/fiber/v2"
	"log"

	"github.com/mrqiz/task-9/internal/database"
	"github.com/mrqiz/task-9/internal/fiber/routes"
	"github.com/mrqiz/task-9/internal/models"
)

func init() {
	database.Connect()
	database.DB.AutoMigrate(&models.Contact{})
}

func main() {
	app := fiber.New()
	contacts := app.Group("/contacts")
	fiberroutes.ContactsRouter(contacts)
	log.Fatal(app.Listen(":3000"))

	db, _ := database.DB.DB()
	defer db.Close()
}
