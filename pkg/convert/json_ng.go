package convert

import (
	"encoding/json"
	"fmt"

	"github.com/expect-digital/translate/pkg/model"
)

func FromNgJson(b []byte) (model.Messages, error) {
	var msg model.Messages

	err := json.Unmarshal(b, &msg)
	if err != nil {
		return model.Messages{}, fmt.Errorf("decode model.Messages to NG JSON: %w", err)
	}

	return msg, nil
}

func ToNgJson(m model.Messages) ([]byte, error) {
	msg, err := json.Marshal(m)
	if err != nil {
		return nil, fmt.Errorf("encode model.Messages to NG JSON: %w", err)
	}

	return msg, nil
}
