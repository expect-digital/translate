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

func (t *TranslateServiceServer) DownloadTranslationFile(
	ctx context.Context,
	req *pb.DownloadTranslationFileRequest,
) (*pb.DownloadTranslationFileResponse, error) {
	if len(req.GetTranslationId()) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Missing param translation_id")
	}

	var (
		data []byte
		err  error
	)

	// find file from DB/FS
	messages := model.Messages{}

	switch req.GetType() {
	default:
		return nil, status.Errorf(codes.Internal, "")
	case pb.Type_TYPE_UNSPECIFIED:
		return nil, status.Errorf(codes.InvalidArgument, "Type is missing")
	case pb.Type_NG_LOCALISE:
		data, err = convert.ToNgJson(messages)
	case pb.Type_NGX_TRANSLATE:
		data, err = convert.ToNgxTranslate(messages)
	case pb.Type_GO:
		data, err = convert.FromGoMessages(messages)
	}

	// todo Proper error handling with status codes
	if err != nil {
		return nil, fmt.Errorf("convert to bytes: %w", err)
	}

	return &pb.DownloadTranslationFileResponse{Data: data}, nil
}
