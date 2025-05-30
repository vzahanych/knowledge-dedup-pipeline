package storage

import (
	"encoding/json"
	"github.com/dgraph-io/badger/v4"
)

type ObjectMeta struct {
	ID         string
	Description string
	Price      string
	MinHashSig []uint64
	SimHashSig uint64
	FeatureVec []float64
}

type Store struct {
	DB *badger.DB
}

func NewStore(path string) (*Store, error) {
	db, err := badger.Open(badger.DefaultOptions(path))
	if err != nil {
		return nil, err
	}
	return &Store{DB: db}, nil
}

func (s *Store) PutMeta(meta ObjectMeta) error {
	return s.DB.Update(func(txn *badger.Txn) error {
		data, _ := json.Marshal(meta)
		return txn.Set([]byte("meta_"+meta.ID), data)
	})
}
