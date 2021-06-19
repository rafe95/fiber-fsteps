package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hashicorp/go-uuid"
)

func GetAll(c *fiber.Ctx) error {
	// List all
	var all []ToDo
	todos, err := txn.Get("toDo", "id")
	if err != nil {
		panic(err)
	}

	for obj := todos.Next(); obj != nil; obj = todos.Next() {
		p := obj.(*ToDo)
		all = append(all, *p)
	}

	return c.JSON(all)
}

func GetById(c *fiber.Ctx) error {
	// Lookup by id
	id := c.Params("id")
	raw, err := txn.First("toDo", "id", id)
	if err != nil {
		panic(err)
	}
	return c.JSON(raw.(*ToDo))
}

func Update(c *fiber.Ctx) error {

	todo := new(ToDo)
	err := c.BodyParser(todo)
	if err != nil {
		return err
	}

	txn = db.Txn(true)
	txn.Insert("toDo", todo)

	txn.Commit()
	txn = db.Txn(false)

	return c.JSON(todo)
}

func Save(c *fiber.Ctx) error {

	txn = db.Txn(true)
	todo := new(ToDo)
	err := c.BodyParser(todo)
	if err != nil {
		return err
	}
	todo.ID, _ = uuid.GenerateUUID()
	txn.Insert("toDo", todo)

	txn.Commit()
	txn = db.Txn(false)

	return c.Status(201).JSON(todo)
}

func DeleteById(c *fiber.Ctx) error {
	id := c.Params("id")
	txn = db.Txn(true)

	raw, err := txn.First("toDo", "id", id)
	if err != nil {
		panic(err)
	}
	txn.Delete("toDo", raw.(*ToDo))

	txn.Commit()
	txn = db.Txn(false)

	return c.Status(204).JSON("")
}
