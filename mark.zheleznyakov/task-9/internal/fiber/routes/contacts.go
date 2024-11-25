package fiberroutes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/mrqiz/task-9/internal/fiber/handlers"
)

func ContactsRouter(app fiber.Router) {
	app.Get("/", fiberhandlers.GetContacts)
	app.Get("/:id", fiberhandlers.GetContact)
	app.Post("/", fiberhandlers.PostContacts)
	app.Put("/:id", fiberhandlers.PutContact)
  app.Delete("/:id", fiberhandlers.DeleteContact)
}
