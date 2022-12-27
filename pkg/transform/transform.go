package transform

import (
	"github.com/expect-digital/translate/pkg/model"
	pb "github.com/expect-digital/translate/pkg/server/translate/v1"
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
		Language: m.Language,
		Messages: messagesToAdd,
		Labels:   m.Labels,
	}
}

func MessagesToProtobuf(m model.Messages) *pb.Messages {
	messagesToAdd := make([]*pb.Message, 0, len(m.Messages))

	for _, msg := range m.Messages {
		messagesToAdd = append(messagesToAdd, MessageToProtobuf(msg))
	}

	return &pb.Messages{
		Language: m.Language,
		Messages: messagesToAdd,
		Labels:   m.Labels,
	}
}

func MessageToProtobuf(m model.Message) *pb.Message {
	return &pb.Message{
		Id:      m.ID,
		Message: m.Message,
		Fuzzy:   m.Fuzzy,
	}
}
