package database

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

type Todo struct {
	ID    int
	Value string
}

func InitDB(dbPath string) *bolt.DB {
	// Open BoltDB data file in DB_PATH.
	// todo.db created if doesnt exist
	db, err := bolt.Open(dbPath, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create db bucket "Todos" if doesnt exist
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("Todos"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})

	return db
}

func CreateTodoEntry(db *bolt.DB, entry *Todo) error {
	return db.Update(func(tx *bolt.Tx) error {
		// Retrieve the Todos bucket.
		// Created when the DB is first opened if not already existing
		b := tx.Bucket([]byte("Todos"))

		// Generate ID for the todo.
		// This returns an error only if the Tx is closed or not writeable.
		id, _ := b.NextSequence()
		entry.ID = int(id)

		// Marshal user data into bytes.
		buf, err := json.Marshal(entry)
		if err != nil {
			return err
		}

		// Persist bytes to "todos" bucket.
		return b.Put(itob(entry.ID), buf)
	})
}

// itob returns an 8-byte big endian representation of v.
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func ViewTodoList(db *bolt.DB) {

	db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte("Todos"))

		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("key=%s, value=%s\n", k, v)
			// fmt.Printf("%s. %s\n", k, v) final format
		}

		return nil
	})
}
