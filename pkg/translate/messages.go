package translate

import (
	"context"

	pb "github.com/expect-digital/translate/pkg/server/translate/v1"
	"github.com/expect-digital/translate/pkg/transform"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (t *TranslateServiceServer) UpdateMessages(
	ctx context.Context,
	req *pb.UpdateMessagesRequest,
) (*pb.UpdateMessagesResponse, error) {
	var (
		reqMessagesID = req.GetMessagesId()
		reqMessages   = req.GetMessages()
	)

	if len(reqMessagesID) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "missing translation_id")
	}

	messages, err := transform.MessagesFromProto(reqMessages)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "parse messages: %s", err)
	}

	// Load from DB
	// Update with new values
	// Store updated

	_ = messages

	return &pb.UpdateMessagesResponse{}, nil
}

func (t *TranslateServiceServer) ListMessages(
	ctx context.Context,
	req *pb.ListMessagesRequest,
) (*pb.ListMessagesResponse, error) {
	allMessages, err := t.repo.ListMessages()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "list all messages: %s", err)
	}

	allMessagesPb := make([]*pb.Messages, 0, len(allMessages))

	for _, msg := range allMessages {
		allMessagesPb = append(allMessagesPb, transform.MessagesToProtobuf(msg))
	}

	return &pb.ListMessagesResponse{Messages: allMessagesPb}, nil
}
