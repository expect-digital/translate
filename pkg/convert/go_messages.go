package convert

import (
	"encoding/json"
	"fmt"

	"github.com/expect-digital/translate/pkg/model"
	"golang.org/x/text/message/pipeline"
)

func FromGoMessages(m model.Messages) ([]byte, error) {
	pipelineMsg := MessagesToPipeline(m)

	msg, err := json.Marshal(pipelineMsg)
	if err != nil {
		return []byte{}, fmt.Errorf("error while marshaling: %w", err)
	}

	return msg, nil
}

func ToGoMessages(b []byte) (pipeline.Messages, error) {
	var msg pipeline.Messages

	err := json.Unmarshal(b, &msg)
	if err != nil {
		return pipeline.Messages{}, fmt.Errorf("error while unmarshaling: %w", err)
	}

	return msg, nil
}

func MessagesToPipeline(m model.Messages) pipeline.Messages {
	pipelineMsg := pipeline.Messages{}

	pipelineMsg.Language = m.Language

	for key, value := range m.Labels {
		pipelineMsg.Macros[key] = pipeline.Text{Msg: value}
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
