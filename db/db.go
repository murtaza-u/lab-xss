package db

import (
	"bytes"

	"github.com/boltdb/bolt"
)

type DB struct {
	Conn   *bolt.DB
	bucket string
}

func Init(bucket, file string) (*DB, error) {
	db, err := bolt.Open(file, 0600, nil)
	if err != nil {
		return nil, err
	}

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucket))
		return err
	})

	if err != nil {
		return nil, err
	}

	return &DB{Conn: db, bucket: bucket}, nil
}

func (db *DB) Get(key []byte) []byte {
	var value []byte

	db.Conn.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(db.bucket))
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			if bytes.Compare(key, k) == 0 {
				value = v
				break
			}
		}

		return nil
	})

	return value
}

func (db *DB) Put(key, value []byte) error {
	err := db.Conn.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(db.bucket))
		return b.Put(key, value)
	})

	return err
}

func (db *DB) Exists(key []byte) bool {
	var exists bool

	db.Conn.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(db.bucket))
		c := b.Cursor()

		for k, _ := c.First(); k != nil; k, _ = c.Next() {
			if bytes.Compare(key, k) == 0 {
				exists = true
				break
			}
		}

		return nil
	})

	return exists
}
