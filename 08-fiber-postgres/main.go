package main

import (
	"fiber-postgres/database"
	"fiber-postgres/handlers"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	database.Connect()

	app.Get("/books", handlers.GetBooks)
	app.Get("/books/:id", handlers.GetBook)
	app.Post("/books", handlers.CreateBook)
	app.Put("/books/:id", handlers.UpdateBook)
	app.Delete("/books/:id", handlers.DeleteBook)

	app.Get("/categories", handlers.GetCategories)
	app.Get("/categories/:id", handlers.GetCategory)
	app.Post("/categories", handlers.CreateCategory)
	app.Put("/categories/:id", handlers.UpdateCategory)
	app.Delete("/categories/:id", handlers.DeleteCategory)

	log.Fatal(app.Listen(":3000"))
}
