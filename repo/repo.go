package repo

import (
	"encoding/json"
	"fmt"

	"github.com/dgraph-io/badger/v3"
	"github.com/expect-digital/translate/pkg/model"
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
		return nil, fmt.Errorf("connect to DB: %w", err)
	}

	return db, nil
}

func (repo *Repo) SaveMessages(id string, m model.Messages) error {
	messagesJson, err := encodeMessages(m)
	if err != nil {
		return err
	}

	err = repo.db.Update(func(txn *badger.Txn) error {
		err = txn.Set([]byte(id), messagesJson)
		if err != nil {
			err = fmt.Errorf("setting key/value pairs %w", err)
		}
		return err
	})
	if err != nil {
		return fmt.Errorf("creating read/write transaction: %w", err)
	}

	return nil
}

func encodeMessages(m model.Messages) ([]byte, error) {
	data, err := json.Marshal(m)
	if err != nil {
		err = fmt.Errorf("model.Messages marshaliing to JSON %w", err)
		return nil, err
	}

	return data, err
}
