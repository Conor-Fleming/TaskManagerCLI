package database

import (
	"encoding/binary"
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

	fn := func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err
	}

	return db.Update(fn)
}

func DoneTask(id int) error {
	task := ViewTask(id)
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		return b.Put(itob(id), []byte(task+" - DONE"))
	})
	if err != nil {
		return err
	}
	return nil
}

func RemoveTask(id int) (int, error) {
	var key int
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		key = id
		return b.Delete(itob(id))
	})
	if err != nil {
		return -1, err
	}
	return key, nil
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
	var result []Task
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			task := Task{
				Key:   btoi(k),
				Value: string(v),
			}
			result = append(result, task)
		}
		return nil
	})
	return result, nil
}

func ViewTask(id int) string {
	var task string
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		v := b.Get(itob(id))
		task = string(v)
		return nil
	})
	return task
}
