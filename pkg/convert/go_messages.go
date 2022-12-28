package convert

import (
	"encoding/json"
	"fmt"

	"github.com/expect-digital/translate/pkg/model"
)

func FromGoMessages(m model.Messages) ([]byte, error) {
	msg, err := json.Marshal(m)
	if err != nil {
		return []byte{}, fmt.Errorf("error while marshaling: %w", err)
	}

	return msg, nil
}

func ToGoMessages(b []byte) (model.Messages, error) {
	var msg model.Messages

	err := json.Unmarshal(b, &msg)
	if err != nil {
		return model.Messages{}, fmt.Errorf("error while unmarshaling: %w", err)
	}

	return msg, nil
}
