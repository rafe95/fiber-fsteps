package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {

	migrateDb()

	populateDb()

	app := fiber.New()

	app.Get("/todo/:id", GetById)

	app.Get("/todos", GetAll)

	app.Post("/todo", Save)

	app.Put("todo/:id", Update)

	app.Delete("/todo/:id", DeleteById)

	app.Listen(":8000")

}
