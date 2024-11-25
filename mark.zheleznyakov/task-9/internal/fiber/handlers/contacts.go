package fiberhandlers

import (
	"github.com/gofiber/fiber/v2"
	"strconv"

	"github.com/mrqiz/task-9/internal/database"
	"github.com/mrqiz/task-9/internal/models"
)

func GetContacts(c *fiber.Ctx) error {
	var contacts []models.Contact
	database.DB.Find(&contacts)
	return c.JSON(contacts)
}

func GetContact(c *fiber.Ctx) error {
	var contact models.Contact
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("NaN")
	}
	err = database.DB.First(&contact, id).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Contact is not found")
	}
	return c.JSON(contact)
}
