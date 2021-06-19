package main

import (
	"github.com/gofiber/fiber/v2"
)

type ToDo struct {
	ID   string `json:"id"`
	Task string `json:"task"`
}

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
