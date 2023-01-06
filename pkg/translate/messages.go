package translate

import (
	"context"

	"github.com/expect-digital/translate/pkg/model"
	pb "github.com/expect-digital/translate/pkg/server/translate/v1"
	"github.com/expect-digital/translate/pkg/transform"
)

func (t *TranslateServiceServer) ListMessages(
	ctx context.Context,
	req *pb.ListMessagesRequest,
) (*pb.ListMessagesResponse, error) {
	// Get all messages from DB
	allMessages := model.Messages{}

	resp := pb.ListMessagesResponse{
		Messages: transform.MessagesToProtobuf(allMessages),
	}

	return &resp, nil
}
