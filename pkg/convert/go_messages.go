package convert

import (
	"encoding/json"
	"fmt"

	"github.com/expect-digital/translate/pkg/model"
	"golang.org/x/text/message/pipeline"
)

func FromGoMessages(m model.Messages) ([]byte, error) {
	pipelineMsg := messagesToPipeline(m)

	msg, err := json.Marshal(pipelineMsg)
	if err != nil {
		return nil, fmt.Errorf("failed pipeline.Messages marshaling: %w", err)
	}

	return msg, nil
}

func ToGoMessages(b []byte) (model.Messages, error) {
	var pipelineMsg pipeline.Messages

	err := json.Unmarshal(b, &pipelineMsg)
	if err != nil {
		return model.Messages{}, fmt.Errorf("failed pipeline.Messages unmarshaling: %w", err)
	}

	msg := messagesFromPipeline(pipelineMsg)

	return msg, nil
}

func messagesToPipeline(m model.Messages) pipeline.Messages {
	pipelineMsg := pipeline.Messages{
		Language: m.Language,
		Messages: make([]pipeline.Message, 0, len(m.Messages)),
	}

	for _, value := range m.Messages {
		pipelineMsg.Messages = append(pipelineMsg.Messages, pipeline.Message{
			ID:      pipeline.IDList{value.ID},
			Fuzzy:   value.Fuzzy,
			Message: pipeline.Text{Msg: value.Message},
		})
	}

	return pipelineMsg
}

func messagesFromPipeline(m pipeline.Messages) model.Messages {
	msg := model.Messages{
		Language: m.Language,
		Messages: make([]model.Message, 0, len(m.Messages)),
	}

	for i, value := range m.Messages {
		msg.Messages = append(msg.Messages, model.Message{
			ID:      value.ID[i],
			Fuzzy:   value.Fuzzy,
			Message: value.Message.Msg,
		})
	}

	return msg
}
