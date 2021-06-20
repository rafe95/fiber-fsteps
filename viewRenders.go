package main

import "github.com/gofiber/fiber/v2"

func Home(c *fiber.Ctx) error {
	// Render index template
	return c.Render("index", fiber.Map{
		"Title":       "Todos",
		"CreateLabel": "Create a new ToDo",
	})
}
