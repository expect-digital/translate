package transform

import (
	"fmt"

	"github.com/expect-digital/translate/pkg/model"
	pb "github.com/expect-digital/translate/pkg/server/translate/v1"
	"golang.org/x/text/language"
)

func MessageFromProto(m *pb.Message) model.Message {
	return model.Message{
		ID:      m.Id,
		Message: m.Message,
		Fuzzy:   m.Fuzzy,
	}
}

func MessagesFromProto(m *pb.Messages) (model.Messages, error) {
	messagesToAdd := make([]model.Message, 0, len(m.Messages))

	for _, msg := range m.Messages {
		messagesToAdd = append(messagesToAdd, MessageFromProto(msg))
	}

	tag, err := language.Parse(m.Language)
	if err != nil {
		return model.Messages{}, fmt.Errorf("language code (BCP 47) is invalid: %w", err)
	}

	messageModel := model.Messages{
		Labels:   m.Labels,
		Language: tag,
		Messages: messagesToAdd,
	}

	return messageModel, nil
}

func MessagesToProtobuf(m model.Messages) *pb.Messages {
	messagesToAdd := make([]*pb.Message, 0, len(m.Messages))

	for _, msg := range m.Messages {
		messagesToAdd = append(messagesToAdd, MessageToProtobuf(msg))
	}

	return &pb.Messages{
		Labels:   m.Labels,
		Language: m.Language.String(),
		Messages: messagesToAdd,
	}
}

func MessageToProtobuf(m model.Message) *pb.Message {
	return &pb.Message{
		Id:      m.ID,
		Message: m.Message,
		Fuzzy:   m.Fuzzy,
	}
}
