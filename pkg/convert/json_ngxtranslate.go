package convert

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/expect-digital/translate/pkg/model"
)

var ErrWrongKey = errors.New("Wrong key type")

func fromNgxTranslate(b []byte) (messages model.Messages, err error) {
	var dst map[string]interface{}

	if err = json.Unmarshal(b, &dst); err != nil {
		return messages, fmt.Errorf("unmarshal ngx translate: %w", err)
	}

	traverseMap := func(key string, value any) (err error) {
		return err
	}

	traverseMap = func(key string, v any) (err error) {
		switch value := v.(type) {
		default:
			return ErrWrongKey
		case string:
			messages.Messages = append(messages.Messages, model.Message{ID: key, Message: value})
		case map[string]interface{}:
			for key2, val2 := range value {
				var newKey string
				if key != "" {
					newKey += key + "."
				}

				newKey += key2
				if err = traverseMap(newKey, val2); err != nil {
					return err
				}
			}
		}

		return err
	}

	if err = traverseMap("", dst); err != nil {
		return messages, fmt.Errorf("traverse ngx: %w", err)
	}

	return messages, nil
}

func toNgxTranslate(messages model.Messages) (b []byte, err error) {
	// We Omit Fuzzy key as it is not in specifications
	type tmpMessage struct {
		ID, Message string
	}

	tmpMessages := make([]tmpMessage, 0, len(messages.Messages))

	for _, msg := range messages.Messages {
		tmpMessages = append(tmpMessages, tmpMessage{ID: msg.ID, Message: msg.Message})
	}

	b, err = json.Marshal(tmpMessages)
	if err != nil {
		return nil, fmt.Errorf("marshal messages to ngx translate : %w", err)
	}

	return b, nil
}
