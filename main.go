/*
Copyright © 2022 Steven Heggie github.com/stevenheggie
*/
package main

import (
	"github.com/stevenheggie/task-cli/cmd"
	"github.com/stevenheggie/task-cli/database"
)

func main() {

	// Initialise BoltDB
	database.InitDB("./todo.db") //TODO: use viper and set this in config file

	// Execute cli cmd
	cmd.Execute()
}
