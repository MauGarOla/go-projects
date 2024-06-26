package handlers

import (
	"fiber-postgres/database"
	"fiber-postgres/models"

	"github.com/gofiber/fiber/v2"
)

func GetCategories(c *fiber.Ctx) error {
	var categories []models.Category
	database.DB.Find(&categories)
	return c.JSON(categories)
}

func GetCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	var category models.Category
	result := database.DB.First(&category, id)
	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Category not found",
		})
	}
	return c.JSON(category)
}

func CreateCategory(c *fiber.Ctx) error {
	var category models.Category
	if err := c.BodyParser(&category); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Cannot parse JSON",
		})
	}
	database.DB.Create(&category)
	return c.Status(201).JSON(category)
}

func UpdateCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	var category models.Category
	result := database.DB.First(&category, id)
	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Category not found",
		})
	}
	if err := c.BodyParser(&category); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Cannot parse JSON",
		})
	}
	database.DB.Save(&category)
	return c.JSON(category)
}

func DeleteCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	var category models.Category
	result := database.DB.First(&category, id)
	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Category not found",
		})
	}
	database.DB.Delete(&category)
	return c.SendStatus(204)
}
