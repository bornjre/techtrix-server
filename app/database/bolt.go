package database

import (
	"encoding/json"
	"errors"

	"github.com/bornjre/techtrix-server/app/config"

	bolt "go.etcd.io/bbolt"
)

var (
	ErrorNoBucket         = errors.New("Bucket not found of that name")
	ErrorValueNotFound    = errors.New("Cannot value with that key")
	ErrorNotExistToupdate = errors.New("key doesnot exist to update the value")
	ErrorAlreadyExist     = errors.New("key already exist to update the value")
)

type Boltdb struct {
	conn *bolt.DB
}

var DB = &Boltdb{}

func Close() {
	DB.Close()
}

func init() {
	firstrun := config.IsFirstRun()

	err := DB.Open("block.db")
	if err != nil {
		panic(err)
	}
	if firstrun {
		runMigration()
	}
}

func runMigration() {

}

func (b *Boltdb) Open(path string) error {

	db, err := bolt.Open(path, 0666, nil)
	if err != nil {
		return err
	}
	b.conn = db

	return nil
}

func (b *Boltdb) Create(key, value, bucket []byte) error {
	print("KEY::")
	print(key)

	err := b.conn.Update(func(tx *bolt.Tx) error {
		bk, err := tx.CreateBucketIfNotExists(bucket)
		if err != nil {
			return err
		}

		olddata := bk.Get(key)
		if olddata != nil {
			return ErrorAlreadyExist
		}

		if err := bk.Put(key, value); err != nil {
			return err
		}
		return nil
	})
	return err
}

func (b *Boltdb) Read(key, bucket []byte) ([]byte, error) {

	var value []byte

	err := b.conn.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)
		if b == nil {
			return ErrorNoBucket
		}

		value = b.Get(key)
		if value == nil {
			return ErrorValueNotFound
		}
		return nil
	})
	return value, err
}

func (b *Boltdb) ReadAll(bucket []byte) ([][]byte, error) {

	var values [][]byte

	err := b.conn.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)
		if b == nil {
			return ErrorNoBucket
		}

		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {

			values = append(values, v)
		}

		return nil
	})
	return values, err
}

func (b *Boltdb) Update(key, value, bucket []byte) error {

	err := b.conn.Update(func(tx *bolt.Tx) error {
		bk := tx.Bucket(bucket)

		if bk == nil {
			return ErrorNoBucket
		}

		olddata := bk.Get(key)
		if olddata == nil {
			return ErrorNotExistToupdate
		}

		if err := bk.Put(key, value); err != nil {
			return err
		}
		return nil
	})
	return err

}

func (b *Boltdb) Delete(key, bucket []byte) error {
	err := b.conn.Update(func(tx *bolt.Tx) error {
		bk := tx.Bucket(bucket)
		if bk == nil {
			return ErrorNoBucket
		}
		if err := bk.Delete(key); err != nil {
			return err
		}
		return nil
	})
	return err

}

func (b *Boltdb) Close() {
	b.conn.Close()
}

// Encode is unified way to serilize data, this way it might be easier to change encoding(json to gob/msgpack or protobuf?)
// without having to huntdown and change all serilization code
func Encode(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// Decode opposite of encode
func Decode(raw []byte, out interface{}) error {
	return json.Unmarshal(raw, out)
}
