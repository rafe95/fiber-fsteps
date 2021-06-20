package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {

	migrateDb()

	populateDb()

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/static", "./static")

	app.Get("/", Home)

	app.Get("/todo/:id", GetById)

	app.Get("/todos", GetAll)

	app.Post("/todo", Save)

	app.Put("todo/:id", Update)

	app.Delete("/todo/:id", DeleteById)

	app.Listen(":8000")

}
