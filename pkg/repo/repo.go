package repo

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/dgraph-io/badger/v3"
	"github.com/expect-digital/translate/pkg/model"
	"golang.org/x/exp/slices"
)

var ErrNotFound = errors.New("not found in Database")

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
	key := []byte(m.TranslationID)

	messagesList, err := r.LoadMessages(m.TranslationID)
	if err != nil && !errors.Is(err, ErrNotFound) {
		return fmt.Errorf("loading messages: %w", err)
	}

	i := slices.IndexFunc(messagesList, func(v model.Messages) bool {
		return v.Language.String() == m.Language.String()
	})
	if i >= 0 {
		messagesList[i] = m
	} else {
		messagesList = append(messagesList, m)
	}

	messagesJson, err := json.Marshal(messagesList)
	if err != nil {
		return fmt.Errorf("marshaling model.Messages to JSON: %w", err)
	}

	err = r.db.Update(func(txn *badger.Txn) error {
		return txn.Set(key, messagesJson) //nolint:wrapcheck
	})
	if err != nil {
		return fmt.Errorf("inserting data into DB: %w", err)
	}

	return nil
}

func (r *Repo) LoadMessages(translationID string) ([]model.Messages, error) {
	var messagesList []model.Messages

	key := []byte(translationID)

	err := r.db.View(func(txn *badger.Txn) error {
		v, err := txn.Get(key)
		if err != nil {
			return err //nolint:wrapcheck
		}

		return v.Value(func(val []byte) error { //nolint:wrapcheck
			return json.Unmarshal(val, &messagesList) //nolint:wrapcheck
		})
	})

	switch {
	default:
		return messagesList, nil
	case errors.Is(err, badger.ErrKeyNotFound):
		return messagesList, ErrNotFound
	case err != nil:
		return nil, fmt.Errorf("getting '%s' from DB: %w", translationID, err)
	}
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
		return nil, fmt.Errorf("badgerDB read :%w", err)
	}

	return messages, nil
}
