package convert

import (
	"encoding/json"
	"fmt"

	"github.com/expect-digital/translate/pkg/model"
)

func EncodeToBytes(m model.Messages) ([]byte, error) {
	message := model.Messages{
		Messages: m.Messages,
		Language: m.Language,
		Labels:   m.Labels,
	}

	msg, err := json.Marshal(message)
	if err != nil {
		return []byte{}, fmt.Errorf("error while marshaling: %w", err)
	}

	return msg, nil
}

func DecodeToMessages(b []byte) (model.Messages, error) {
	var msg model.Messages

	err := json.Unmarshal(b, &msg)
	if err != nil {
		return model.Messages{}, fmt.Errorf("error while unmarshaling: %w", err)
	}

	return msg, nil
}
