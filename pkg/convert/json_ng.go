package convert

import (
	"encoding/json"
	"fmt"

	"github.com/expect-digital/translate/pkg/model"
	"golang.org/x/text/language"
)

type ngJson struct {
	Locale       language.Tag      `json:"locale"`
	Translations map[string]string `json:"translations"`
}

func FromNgJson(b []byte) (model.Messages, error) {
	var angularJson ngJson

	err := json.Unmarshal(b, &angularJson)
	if err != nil {
		return model.Messages{}, fmt.Errorf("decode NG JSON to go messages: %w", err)
	}

	msg := model.Messages{
		Language: angularJson.Locale,
		Messages: make([]model.Message, len(angularJson.Translations)),
	}

	for key, value := range angularJson.Translations {
		msg.Messages[0] = model.Message{
			ID:      key,
			Message: value,
		}
	}

	return msg, nil
}

func ToNgJson(m model.Messages) ([]byte, error) {
	angularJson := ngJson{
		Locale:       m.Language,
		Translations: make(map[string]string, len(m.Messages)),
	}

	for _, message := range m.Messages {
		angularJson.Translations[message.ID] = message.Message
	}

	result, err := json.Marshal(angularJson)
	if err != nil {
		return nil, fmt.Errorf("encode go NgJson to []byte: %w", err)
	}

	return result, nil
}
