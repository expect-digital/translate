package translate

import (
	"context"

	"github.com/expect-digital/translate/pkg/convert"
	"github.com/expect-digital/translate/pkg/model"
	pb "github.com/expect-digital/translate/pkg/server/translate/v1"
	"golang.org/x/text/language"
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
	var (
		reqTranslationID = req.GetTranslationId()
		reqType          = req.GetType()
		reqLanguage      = req.GetLanguage()
	)

	if len(reqTranslationID) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "missing translation_id")
	}

	if len(reqLanguage) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "language is missing")
	}

	language, err := language.Parse(reqLanguage)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "parse language: %s", err)
	}

	var to convert.To

	switch reqType {
	case pb.Type_TYPE_UNSPECIFIED:
		return nil, status.Errorf(codes.InvalidArgument, "type is missing")
	case pb.Type_NG_LOCALISE:
		to = convert.ToNgJson
	case pb.Type_NGX_TRANSLATE:
		to = convert.ToNgxTranslate
	case pb.Type_GO:
		to = convert.ToGo
	}
	// find file from DB/FS with language
	_ = language
	messages := model.Messages{}

	data, err := to(messages)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "serialize data: %s", err)
	}

	return &pb.DownloadTranslationFileResponse{Data: data}, nil
}
