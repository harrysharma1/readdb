package database

import (
	"encoding/binary"
	"time"

	bolt "go.etcd.io/bbolt"
)

func itob(v uint) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func nextId(b *bolt.Bucket) (uint, error) {
	id, err := b.NextSequence()
	return uint(id), err
}

func InitDB() (*bolt.DB, error) {
	db, err := bolt.Open("./test.db", 0600, &bolt.Options{Timeout: 3 * time.Second})
	if err != nil {
		return nil, err
	}
	err = db.Update(func(tx *bolt.Tx) error {
		buckets := []string{
			"authors",
			"books",
			"publishers",
			"lists"}
		for _, bucket := range buckets {
			_, err := tx.CreateBucketIfNotExists([]byte(bucket))
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
