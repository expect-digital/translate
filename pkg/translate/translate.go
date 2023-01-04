package translate

import (
	"context"
	"fmt"

	"github.com/expect-digital/translate/pkg/convert"
	"github.com/expect-digital/translate/pkg/model"
	pb "github.com/expect-digital/translate/pkg/server/translate/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TranslateServiceServer struct {
	pb.UnimplementedTranslateServiceServer
}

func New() *TranslateServiceServer {
	return new(TranslateServiceServer)
}

func (t *TranslateServiceServer) UploadTranslationFile(
	ctx context.Context,
	req *pb.UploadTranslationFileRequest,
) (*pb.UploadTranslationFileResponse, error) {
	var (
		messages model.Messages
		err      error
	)

	switch req.GetType() {
	default:
		return nil, status.Errorf(codes.Internal, "")
	case pb.Type_TYPE_UNSPECIFIED:
		return nil, status.Errorf(codes.InvalidArgument, "Type is missing")
	case pb.Type_NG_LOCALISE:
		messages, err = convert.FromNgJson([]byte(req.Data))
	case pb.Type_NGX_TRANSLATE:
		messages, err = convert.FromNgxTranslate([]byte(req.Data))
	case pb.Type_GO:
		messages, err = convert.FromGo([]byte(req.Data))
	}

	// todo Proper error handling with status codes
	if err != nil {
		return nil, fmt.Errorf("convert to messages: %w", err)
	}

	// Save to DB/FS...
	_ = messages

	return &pb.UploadTranslationFileResponse{}, nil
}
