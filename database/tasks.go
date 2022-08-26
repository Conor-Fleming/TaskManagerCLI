package database

import (
	"encoding/binary"
	"fmt"
	"time"

	"github.com/boltdb/bolt"
)

var taskBucket = []byte("tasks")

var db *bolt.DB

type Task struct {
	Key   int
	Value string
}

func Init(dbPath string) error {
	var err error
	db, err = bolt.Open("tasks.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	defer db.Close()

	fn := func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err
	}
	return db.Update(fn)
}

func CreateTask(task string) (int, error) {
	var id int
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		//Generating the key for the user
		key, _ := b.NextSequence()
		id = int(key)
		return b.Put(itob(id), []byte(task))
	})
	if err != nil {
		return -1, err
	}
	return id, nil
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}

func ViewList() ([]Task, error) {
	//var result []Task
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		if b == nil {
			return fmt.Errorf("Bucket %s not found", taskBucket)
		}

		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("key = %s, value = %s\n", k, v)
		}

		return nil
	})
}
