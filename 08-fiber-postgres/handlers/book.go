package handlers

import (
	"fiber-postgres/database"
	"fiber-postgres/models"

	"github.com/gofiber/fiber/v2"
)

func GetBooks(c *fiber.Ctx) error {
	var books []models.Book
	database.DB.Find(&books)
	return c.JSON(books)
}

func GetBook(c *fiber.Ctx) error {
	id := c.Params("id")
	var book models.Book
	result := database.DB.First(&book, id)
	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Book not found",
		})
	}
	return c.JSON(book)
}

func CreateBook(c *fiber.Ctx) error {
	var book models.Book
	if err := c.BodyParser(&book); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Cannot parse JSON",
		})
	}
	database.DB.Create(&book)
	return c.Status(201).JSON(book)
}

func UpdateBook(c *fiber.Ctx) error {
	id := c.Params("id")
	var book models.Book
	result := database.DB.First(&book, id)
	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Book not found",
		})
	}
	if err := c.BodyParser(&book); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Cannot parse JSON",
		})
	}
	database.DB.Save(&book)
	return c.JSON(book)
}

func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	var book models.Book
	result := database.DB.First(&book, id)
	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Book not found",
		})
	}
	database.DB.Delete(&book)
	return c.SendStatus(204)
}
