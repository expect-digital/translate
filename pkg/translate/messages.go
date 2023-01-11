package translate

import (
	"context"
	"errors"

	"github.com/expect-digital/translate/pkg/model"
	"github.com/expect-digital/translate/pkg/repo"
	pb "github.com/expect-digital/translate/pkg/server/translate/v1"
	"github.com/expect-digital/translate/pkg/transform"
	"golang.org/x/exp/slices"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (t *TranslateServiceServer) UpdateMessages(
	ctx context.Context,
	req *pb.UpdateMessagesRequest,
) (*pb.UpdateMessagesResponse, error) {
	var (
		reqTranslationID = req.GetTranslationId()
		reqMessages      = req.GetMessages()
	)

	if len(reqTranslationID) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "missing translation_id")
	}

	reqMessages.TranslationId = reqTranslationID

	messages, err := transform.MessagesFromProto(reqMessages)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "parsing messages: %s", err)
	}

	messagesList, err := t.repo.LoadMessages(reqTranslationID)

	switch {
	case errors.Is(err, repo.ErrNotFound):
		return nil, status.Errorf(codes.NotFound, "no translations with id '%s'", reqTranslationID)
	case err != nil:
		return nil, status.Errorf(codes.Internal, "loading messages: %s", err)
	}

	if cnt := slices.ContainsFunc(messagesList, func(m model.Messages) bool {
		return m.Language.String() == messages.Language.String()
	}); !cnt {
		return nil, status.Errorf(codes.NotFound, "no such language exists in Database '%s'", messages.Language.String())
	}

	err = t.repo.SaveMessages(messages)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "saving messages: %s", err)
	}

	return &pb.UpdateMessagesResponse{}, nil
}

func (t *TranslateServiceServer) ListMessages(
	ctx context.Context,
	req *pb.ListMessagesRequest,
) (*pb.ListMessagesResponse, error) {
	allMessages, err := t.repo.ListMessages()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "listing all messages: %s", err)
	}

	allMessagesPb := make([]*pb.Messages, 0, len(allMessages))

	for _, msg := range allMessages {
		allMessagesPb = append(allMessagesPb, transform.MessagesToProtobuf(msg))
	}

	return &pb.ListMessagesResponse{Messages: allMessagesPb}, nil
}
