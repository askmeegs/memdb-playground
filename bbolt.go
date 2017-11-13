// Transaction-based, in-memory DB for Sigma
// Yass

package main

import (
	"encoding/json"
	"log"

	bolt "github.com/coreos/bbolt"
)

const (
	bn = "sigma"
)

// one DB -> one Sigma bucket
type MemDB struct {
	D *bolt.DB
}

type Vacation struct {
	Cost          float64
	Kind          string
	International bool
	PackingList   []string
}

func main() {
	// Create a Memdb
	mdb, err := New()
	if err != nil {
		log.Fatal(err)
	}

	// Put a generic struct in there
	v := Vacation{
		Cost:          314.15,
		Kind:          "backpacking",
		International: true,
		PackingList:   []string{"backpack", "boots", "compass", "map", "train pass"},
	}

	vBytes, err := json.Marshal(v)
	if err != nil {
		log.Fatal(err)
	}

	err = mdb.Upsert("EuropeAdventure", vBytes)
	if err != nil {
		log.Fatal(err)
	}

	//overwrite
	v.Cost = 1000.01
	vBytes, err = json.Marshal(v)
	if err != nil {
		log.Fatal(err)
	}
	err = mdb.Upsert("EuropeAdventure", vBytes)
	if err != nil {
		log.Fatal(err)
	}

	output, err := mdb.Get("EuropeAdventure")
	if err != nil {
		log.Fatal(err)
	}
	var gotV Vacation
	err = json.Unmarshal(output, &gotV)
	log.Printf("AFTER GET, SUCCESSFULLY GOT OUTPUT: \n %#v \n", gotV)

	// Delete then get
	err = mdb.Delete("EuropeAdventure")
	if err != nil {
		log.Fatal(err)
	}

	var postDel Vacation
	output, err = mdb.Get("EuropeAdventure")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(output, &postDel)
	log.Printf("AFTER DELETE, SUCCESSFULLY GOT OUTPUT: \n %#v \n", gotV)

}

func New() (*MemDB, error) {
	db, err := bolt.Open("sigma.mem.db", 0600, nil)
	if err != nil {
		return &MemDB{}, err
	}
	return &MemDB{
		D: db,
	}, nil
}

// -------- GENERIC OPERATIONS ------------------------------------------------

func (m *MemDB) Upsert(key string, valBytes []byte) error {
	keyBytes := []byte(key)
	// Start a write transaction.
	if err := m.D.Update(func(tx *bolt.Tx) error {
		// Create bucket around which to execute the txn
		b, err := tx.CreateBucketIfNotExists([]byte(bn))
		if err != nil {
			return err
		}

		// Set the value "bar" for the key "foo".
		if err := b.Put(keyBytes, valBytes); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (m *MemDB) Get(key string) ([]byte, error) {
	value := []byte{}
	if err := m.D.View(func(tx *bolt.Tx) error {
		value = tx.Bucket([]byte(bn)).Get([]byte(key))
		return nil
	}); err != nil {
		return value, err
	}
	return value, nil
}

// Maps ID to object
func (m *MemDB) Query() (map[string][]byte, error) {
	return map[string][]byte{}, nil
}

func (m *MemDB) Delete(key string) error {
	if err := m.D.Update(func(tx *bolt.Tx) error {
		return tx.Bucket([]byte(bn)).Delete([]byte(key))
	}); err != nil {
		return err
	}
	return nil
}
