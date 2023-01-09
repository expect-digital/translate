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

func (r *Repo) SaveMessages(id string, m model.Messages) error {
	messagesJson, err := json.Marshal(m)
	if err != nil {
		return fmt.Errorf("marshaling model.Messages to JSON: %w", err)
	}

	err = r.db.Update(func(txn *badger.Txn) error {
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

func (r *Repo) ListMessages() ([]model.Messages, error) {
	var messages []model.Messages

	err := r.db.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()

		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			var message model.Messages

			err := item.Value(func(val []byte) error {
				return json.Unmarshal(val, &message) //nolint:wrapcheck
			})
			if err != nil {
				return fmt.Errorf("unmarshal value to messages: %w", err)
			}

			messages = append(messages, message)
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf(":%w", err)
	}

	return messages, nil
}
