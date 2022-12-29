package convert

import (
	"encoding/json"
	"fmt"

	"github.com/expect-digital/translate/pkg/model"
	"golang.org/x/text/language"
)

func FromNgJson(b []byte) (model.Messages, error) {
	incoming := make(map[string]interface{})

	err := json.Unmarshal(b, &incoming)
	if err != nil {
		return model.Messages{}, fmt.Errorf("decode NG JSON to go map: %w", err)
	}

	tag, err := language.Parse(fmt.Sprint(incoming["locale"]))
	if err != nil {
		return model.Messages{}, fmt.Errorf("parse NG JSON key: locale : %w", err)
	}

	msg := model.Messages{
		Language: tag,
		Messages: []model.Message{},
	}

	f := func(in interface{}) {
		switch v := in.(type) {
		case map[string]interface{}:
			for key, value := range v {
				msg.Messages = append(msg.Messages, model.Message{ID: key, Message: fmt.Sprint(value)})
			}
		default:
			return
		}
	}

	f(incoming["translations"])

	return msg, nil
}

func ToNgJson(m model.Messages) ([]byte, error) {
	messages := make(map[string]interface{})
	result := make(map[string]interface{})

	for _, message := range m.Messages {
		messages[message.ID] = message.Message
	}

	result["locale"] = m.Language
	result["translations"] = messages

	msg, err := json.Marshal(result)
	if err != nil {
		return nil, fmt.Errorf("encode model.Messages to NG JSON: %w", err)
	}

	return msg, nil
}
