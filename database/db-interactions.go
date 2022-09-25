package database

import (
	"fmt"
	"log"
	"strconv"

	"github.com/boltdb/bolt"
)

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

func CreateTodoEntry(dbPath string, entry string) error {

	// access db
	db, err := bolt.Open(dbPath, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// generate ID (key) and add entry (value) to db
	return db.Update(func(tx *bolt.Tx) error {
		// Retrieve the Todos bucket.
		// Created when the DB is first opened if not already existing
		b := tx.Bucket([]byte("Todos"))

		// Generate ID for the todo.
		// This returns an error only if the Tx is closed or not writeable.
		id, _ := b.NextSequence()

		// Persist bytes to "todos" bucket.
		return b.Put([]byte(strconv.Itoa(int(id))), []byte(entry))
	})
}

func ViewTodoList(dbPath string) {

	db, err := bolt.Open(dbPath, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte("Todos"))

		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("%s. %s\n", k, v)
		}

		return nil
	})
}

func MarkTodoDone(dbPath string, key string) error {

	db, err := bolt.Open(dbPath, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	return db.Update(func(tx *bolt.Tx) error {
		// Retrieve the Todos bucket.
		// Created when the DB is first opened if not already existing
		b := tx.Bucket([]byte("Todos"))

		// Delete entry from "Todos" bucket with provided key.
		return b.Delete([]byte(key))
	})
}

func ReorderKeys(dbPath string) error {

	db, err := bolt.Open(dbPath, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	entriesMap := make(map[string]string)

	db.View(func(tx *bolt.Tx) error {
		// Retrieve the Todos bucket.
		// Created when the DB is first opened if not already existing
		b := tx.Bucket([]byte("Todos"))

		b.ForEach(func(k, v []byte) error {

			entriesMap[string(k)] = string(v)

			return nil
		})
		return nil
	})

	db.Update(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte("Todos"))

		for k, _ := range entriesMap {
			// Delete entry from "Todos" bucket with provided key.
			b.Delete([]byte(k))
		}
		return nil
	})

	db.Update(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte("Todos"))
		b.SetSequence(0)

		for _, v := range entriesMap {
			// Generate ID for the todo.
			// This returns an error only if the Tx is closed or not writeable.
			id, _ := b.NextSequence()

			// Persist bytes to "todos" bucket.
			b.Put([]byte(strconv.Itoa(int(id))), []byte(v))
		}
		return nil

	})

	return nil
}
