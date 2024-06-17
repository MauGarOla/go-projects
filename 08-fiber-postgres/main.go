package main

import (
	"fiber-postgres/database"
	"fiber-postgres/handlers"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Conectar a la base de datos
	database.Connect()

	// Definir rutas
	app.Get("/books", handlers.GetBooks)
	app.Get("/books/:id", handlers.GetBook)
	app.Post("/books", handlers.CreateBook)
	app.Put("/books/:id", handlers.UpdateBook)
	app.Delete("/books/:id", handlers.DeleteBook)

	// Iniciar el servidor
	log.Fatal(app.Listen(":3000"))
}
