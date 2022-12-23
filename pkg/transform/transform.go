package transform

import (
	pb "github.com/expect-digital/translate/api/translate/translate/v1/gen"
	"github.com/expect-digital/translate/pkg/model"
)

func MessageFromProto(m *pb.Message) model.Message {
	return model.Message{
		ID:      m.Id,
		Message: m.Message,
		Fuzzy:   m.Fuzzy,
	}
}

func MessagesFromProto(m *pb.Messages) model.Messages {
	var messagesToAdd []model.Message

	for i := range m.Messages {
		messagesToAdd = append(messagesToAdd, MessageFromProto(m.Messages[i]))
	}

	return model.Messages{
		Language: m.Language,
		Messages: messagesToAdd,
	}
}

func MessagesToProtobuf(m model.Messages) *pb.Messages {
	var messagesToAdd []*pb.Message

	for i := range m.Messages {
		messagesToAdd = append(messagesToAdd, MessageToProtobuf(m.Messages[i]))
	}

	return &pb.Messages{
		Language: m.Language,
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
