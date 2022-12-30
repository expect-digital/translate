package convert

import (
	"encoding/json"
	"fmt"

	"github.com/expect-digital/translate/pkg/model"
	"golang.org/x/text/language"
)

type NgJson struct {
	Locale       language.Tag      `json:"locale"`
	Translations map[string]string `json:"translations"`
}

func FromNgJson(b []byte) (model.Messages, error) {
	var ngJson NgJson

	err := json.Unmarshal(b, &ngJson)
	if err != nil {
		return model.Messages{}, fmt.Errorf("decode NG JSON to go messages: %w", err)
	}

	msg := model.Messages{
		Language: ngJson.Locale,
		Messages: []model.Message{},
	}

	for key, value := range ngJson.Translations {
		msg.Messages = append(msg.Messages, model.Message{
			ID:      key,
			Message: value,
		})
	}

	return msg, nil
}

func ToNgJson(m model.Messages) ([]byte, error) {
	ngJson := NgJson{
		Locale:       m.Language,
		Translations: make(map[string]string),
	}

	for _, message := range m.Messages {
		ngJson.Translations[message.ID] = message.Message
	}

	result, err := json.Marshal(ngJson)
	if err != nil {
		return nil, fmt.Errorf("encode go NgJson to []byte: %w", err)
	}

	return result, nil
}
