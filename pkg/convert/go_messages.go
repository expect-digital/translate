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
		return nil, fmt.Errorf("error while marshaling: %w", err)
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

func messagesToPipeline(m model.Messages) pipeline.Messages {
	pipelineMsg := pipeline.Messages{
		Language: m.Language,
		Messages: make([]pipeline.Message, 0, len(m.Messages)),
	}

	for _, value := range m.Messages {
		pipelineMsg.Messages = append(pipelineMsg.Messages, pipeline.Message{
			ID:      append(pipeline.IDList{}, value.ID),
			Fuzzy:   value.Fuzzy,
			Message: pipeline.Text{Msg: value.Message},
		})
	}

	return pipelineMsg
}
