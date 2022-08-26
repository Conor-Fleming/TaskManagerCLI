package database

import (
	"encoding/binary"
	"encoding/json"
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

func CreateTask(task *Task) (int, error) {
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		//Generating the key for the user
		key, _ := b.NextSequence()
		task.Key = int(key)

		buf, err := json.Marshal(task)
		if err != nil {
			return err
		}

		return b.Put(itob(task.Key), buf)
	})
	if err != nil {
		return 0, err
	}
	return 0, nil
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

/*
func ViewList() {
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
	})
}*/
