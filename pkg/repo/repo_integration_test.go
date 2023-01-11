package repo

import (
	"log"
	"reflect"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/expect-digital/translate/pkg/model"
	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/slices"
	"golang.org/x/text/language"
)

var repo *Repo

func init() {
	db, err := Connect()
	if err != nil {
		log.Fatal(err)
	}

	repo = NewRepo(db)
}

func Test_SaveMessages(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		messages model.Messages
	}{
		{
			name: "Save messages 1",
			messages: model.Messages{
				TranslationID: gofakeit.UUID(),
				Labels:        map[string]string{},
				Language:      language.English,
				Messages: []model.Message{
					{
						ID:      gofakeit.Word(),
						Message: gofakeit.Word(),
					},
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := repo.SaveMessages(tt.messages)
			assert.NoError(t, err)

			messageList, err := repo.LoadMessages(tt.messages.TranslationID)
			assert.NoError(t, err)

			contains := slices.ContainsFunc(messageList, func(m model.Messages) bool {
				return reflect.DeepEqual(tt.messages, m)
			})
			assert.True(t, contains)

			log.Println(messageList)
		})
	}
}

func Test_SaveMessagesWithSameID(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		messages model.Messages
	}{
		{
			name: "Same id 1",
			messages: model.Messages{
				TranslationID: "asd",
				Labels:        map[string]string{},
				Language:      language.English,
				Messages: []model.Message{
					{
						ID:      gofakeit.Word(),
						Message: gofakeit.Word(),
					},
				},
			},
		},
		{
			name: "Same id 2",
			messages: model.Messages{
				TranslationID: "asd",
				Labels:        map[string]string{},
				Language:      language.Georgian,
				Messages: []model.Message{
					{
						ID:      gofakeit.Word(),
						Message: gofakeit.Word(),
					},
				},
			},
		},
		{
			name: "Replace english",
			messages: model.Messages{
				TranslationID: "asd",
				Labels:        map[string]string{},
				Language:      language.Georgian,
				Messages: []model.Message{
					{
						ID:      gofakeit.Word(),
						Message: gofakeit.Word(),
					},
				},
			},
		},
	}
	for _, tt := range tests { //nolint:paralleltest
		t.Run(tt.name, func(t *testing.T) {
			err := repo.SaveMessages(tt.messages)
			assert.NoError(t, err)

			messageList, err := repo.LoadMessages(tt.messages.TranslationID)
			assert.NoError(t, err)

			contains := slices.ContainsFunc(messageList, func(m model.Messages) bool {
				return reflect.DeepEqual(tt.messages, m)
			})
			assert.True(t, contains)

			log.Println(messageList)
		})
	}
}
