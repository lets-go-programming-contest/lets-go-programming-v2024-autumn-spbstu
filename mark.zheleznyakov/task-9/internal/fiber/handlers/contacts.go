package fiberhandlers

import (
	"github.com/go-playground/validator/v10"
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
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "id is nan",
		})

	}
	err = database.DB.First(&contact, id).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "not found",
		})
	}
	return c.JSON(contact)
}

func PostContacts(c *fiber.Ctx) error {
	var contact models.Contact

	if err := c.BodyParser(&contact); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "malformed body",
		})
	}

	validate := validator.New()
	if err := validate.Struct(contact); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "name and valid phone are required",
		})
	}

	if err := database.DB.Create(&contact).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to create contact",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(contact)
}

func PutContact(c *fiber.Ctx) error {
	id := c.Params("id")
	if _, err := strconv.Atoi(id); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "id is nan",
		})
	}

	var contact models.Contact

	if err := c.BodyParser(&contact); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "malformed body",
		})
	}

	validate := validator.New()
	if err := validate.Struct(contact); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "name and valid phone are required",
		})
	}

	var existingContact models.Contact
	if database.DB.First(&existingContact, id).Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "contact not found",
		})
	}

	existingContact.Name = contact.Name
	existingContact.Phone = contact.Phone

	if err := database.DB.Save(&existingContact).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to update contact",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(contact)
}

func DeleteContact(c *fiber.Ctx) error {
	id := c.Params("id")
	if _, err := strconv.Atoi(id); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "id is nan",
		})
	}

	var existingContact models.Contact
	if database.DB.First(&existingContact, id).Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "contact not found",
		})
	}

	if err := database.DB.Delete(&existingContact).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to delete contact",
		})
	}

  return c.Status(fiber.StatusNoContent).Send([]byte{})
}
