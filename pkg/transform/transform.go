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

func MessagesFromProto(m *pb.Messages) model.Messages {
	messagesToAdd := make([]model.Message, 0, len(m.Messages))

	for _, msg := range m.Messages {
		messagesToAdd = append(messagesToAdd, MessageFromProto(msg))
	}

	return model.Messages{
		Labels:   m.Labels,
		Language: convertToLanguageTag(m.Language),
		Messages: messagesToAdd,
	}
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

func convertToLanguageTag(text string) language.Tag {
	tag, err := language.Parse(text)
	if err != nil {
		fmt.Println(fmt.Errorf("error while parsing: %w", err))
	}

	return tag
}
