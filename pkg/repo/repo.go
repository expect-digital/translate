package repo

import (
	"encoding/json"
	"fmt"

	"github.com/dgraph-io/badger/v3"
	"github.com/expect-digital/translate/pkg/model"
	"golang.org/x/exp/slices"
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

func (r *Repo) SaveMessages(m model.Messages) error {
	var messages []model.Messages

	key := []byte(m.TranslationID)

	err := r.db.View(func(txn *badger.Txn) error {
		v, err := txn.Get(key)
		if err != nil {
			return err
		}

		return v.Value(func(val []byte) error {
			return json.Unmarshal(val, &messages)
		})
	})
	if err != nil {
		return err
	}

	i := slices.IndexFunc(messages, func(v model.Messages) bool {
		return v.Language == m.Language
	})
	if i >= 0 {
		messages[i] = m
	} else {
		messages = append(messages, m)
	}

	messagesJson, err := json.Marshal(messages)
	if err != nil {
		return fmt.Errorf("marshaling model.Messages to JSON: %w", err)
	}

	err = r.db.Update(func(txn *badger.Txn) error {
		return txn.Set(key, messagesJson) //nolint:wrapcheck
	})
	if err != nil {
		return fmt.Errorf("creating read/write transaction: %w", err)
	}

	return nil
}

func (r *Repo) LoadMessages(id string) ([]model.Messages, error) {
	var msg []model.Messages

	err := r.db.View(func(txn *badger.Txn) error {
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
		return nil, fmt.Errorf("creating read transaction: %w", err)
	}

	return msg, nil
}
