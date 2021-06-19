package main

import "github.com/hashicorp/go-memdb"

var schema *memdb.DBSchema
var db *memdb.MemDB
var txn *memdb.Txn

type ToDo struct {
	ID   string `json:"id"`
	Task string `json:"task"`
}

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

	// Create a new data base
	db, _ = memdb.NewMemDB(schema)
	defer txn.Abort()
}

func populateDb() {

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

}
