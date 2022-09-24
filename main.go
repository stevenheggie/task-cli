/*
Copyright © 2022 Steven Heggie github.com/stevenheggie
*/
package main

import (
	"github.com/boltdb/bolt"
	"github.com/stevenheggie/task-cli/cmd"
	"github.com/stevenheggie/task-cli/database"
)

var DB *bolt.DB

const DB_PATH string = "./todo.db"

func main() {

	// Initialise BoltDB
	DB = database.InitDB(DB_PATH)

	// Execute cli cmd
	cmd.Execute()
}

func GetDB() *bolt.DB {
	return DB
}
