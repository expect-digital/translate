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
	messagesJson, err := json.Marshal(m)
	if err != nil {
		return fmt.Errorf("marshaling model.Messages to JSON: %w", err)
	}

	err = repo.db.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(id), messagesJson) //nolint:wrapcheck
	})
	if err != nil {
		return fmt.Errorf("creating read/write transaction: %w", err)
	}

	return nil
}

func (repo *Repo) LoadMessages(id string) (model.Messages, error) {
	var msg model.Messages

	err := repo.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(id))
		if err != nil {
			return fmt.Errorf("getting key/value pair: %w", err)
		}

		err = item.Value(func(val []byte) error {
			err = json.Unmarshal(val, &msg)

			if err != nil {
				return fmt.Errorf("unmarshaling JSON to model.Messages: %w", err)
			}

			return nil
		})

		if err != nil {
			return fmt.Errorf("retrieving value from badger db %w", err)
		}

		return nil
	})
	if err != nil {
		return model.Messages{}, fmt.Errorf("creating read transaction: %w", err)
	}

	return msg, nil
}
