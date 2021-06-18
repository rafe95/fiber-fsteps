package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hashicorp/go-memdb"
)

type ToDo struct {
	ID   string `json:"id"`
	Task string `json:"task"`
}

var schema *memdb.DBSchema
var db *memdb.MemDB
var txn *memdb.Txn

func migrateDb() {
	// Create the DB schema
	schema = &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"toDo": {
				Name: "toDo",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
				},
			},
		},
	}
}

func populateDb() {
	// Create a new data base
	db, _ = memdb.NewMemDB(schema)

	// Create a write transaction
	txn = db.Txn(true)

	// Insert some toDos
	people := []*ToDo{
		{"h1h2", "Bake bread"},
		{"h5h6", "Walk the dog"},
		{"r432", "Complete a form"},
	}
	for _, p := range people {
		if err := txn.Insert("toDo", p); err != nil {
			panic(err)
		}
	}

	// Commit the transaction
	txn.Commit()

	// Create read-only transaction
	txn = db.Txn(false)
	defer txn.Abort()
}

func main() {

	migrateDb()

	populateDb()

	app := fiber.New()

	app.Get("/todo/:id", func(c *fiber.Ctx) error {
		// Lookup by id
		id := c.Params("id")
		raw, err := txn.First("toDo", "id", id)
		if err != nil {
			panic(err)
		}
		return c.JSON(raw.(*ToDo))
	})

	app.Get("/todos", func(c *fiber.Ctx) error {
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
	})

	app.Post("/todo", func(c *fiber.Ctx) error {

		txn = db.Txn(true)
		todo := new(ToDo)
		err := c.BodyParser(todo)
		if err != nil {
			return err
		}
		txn.Insert("toDo", todo)

		txn.Commit()
		txn = db.Txn(false)

		return c.JSON(todo)
	})

	app.Listen(":8000")

}
