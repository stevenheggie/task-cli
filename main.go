/*
Copyright © 2022 Steven Heggie github.com/stevenheggie
*/
package main

import (
	"github.com/stevenheggie/task-cli/cmd"
	"github.com/stevenheggie/task-cli/db"
)

const DB_PATH string = "./todo.db"

func main() {

	// Initialise BoltDB
	db.InitDB(DB_PATH)

	// Execute cli cmd
	cmd.Execute()
}
