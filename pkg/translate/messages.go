package translate

import (
	"context"

	pb "github.com/expect-digital/translate/pkg/server/translate/v1"
	"github.com/expect-digital/translate/pkg/transform"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

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
