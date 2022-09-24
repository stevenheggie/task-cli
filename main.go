/*
Copyright © 2022 Steven Heggie github.com/stevenheggie
*/
package main

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
	"github.com/stevenheggie/task-cli/cmd"
)

func main() {

	// Open my.db (Bolt) data file in current dir.
	// todo.db created if doesnt exist
	db, err := bolt.Open("todo.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("List"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})

	// Execute cli cmd
	cmd.Execute()
}
