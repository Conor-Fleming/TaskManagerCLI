package db

import (
	"fmt"
	"log"
	"time"

	"github.com/boltdb/bolt"
)

func Init() {
	//connecting to DB, using timeout option to prevent hanging
	db, err := bolt.Open("tasks.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("TODOs"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return b.Put([]byte("task-1"), []byte("Testing DB"))
	})

}
