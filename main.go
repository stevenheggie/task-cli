/*
Copyright © 2022 Steven Heggie github.com/stevenheggie
*/
package main

import (
	"fmt"

	"github.com/stevenheggie/task/cmd"
)

func main() {

	// Initialise BoltDB
	// database.InitDB("./todo.db") //TODO: use viper and set this in config file

	fmt.Println("blah")
	// dbPath := string.Join(viper.GetString("DB_DIR") + "/" + viper.Get("DB_NAME"))

	// fmt.Println(dbPath)
	// database.InitDB(dbPath)
	// Execute cli cmd
	cmd.Execute()
}
