package repo

import (
	"fmt"
	"github.com/dgraph-io/badger/v3"
	"github.com/expect-digital/translate/pkg/model"
	"log"
)

type Repo struct {
	db *badger.DB
}

func NewRepo(db *badger.DB) *Repo {
	return &Repo{
		db: db,
	}
}

func Connect() (*badger.DB, error) {
	db, err := badger.Open(badger.DefaultOptions("").WithInMemory(true))
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}

func (repo *Repo) SaveMessages(m model.Messages) error {
	for _, value := range m.Messages {
		err := repo.db.Update(func(txn *badger.Txn) error {
			err := txn.Set([]byte(value.ID), []byte(value.Message))
			return fmt.Errorf("setting key/value pair: %w", err)
		})
		if err != nil {
			return fmt.Errorf("creating read/write transaction: %w", err)
		}
	}
	return nil
}
