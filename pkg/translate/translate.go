package translate

import (
	"context"

	"github.com/expect-digital/translate/pkg/convert"
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

func (t *TranslateServiceServer) UploadTranslationFile(
	ctx context.Context,
	req *pb.UploadTranslationFileRequest,
) (*pb.UploadTranslationFileResponse, error) {
	var (
		reqTranslationID = req.GetTranslationId()
		reqData          = req.GetData()
		reqLabels        = req.GetLabels()
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

	var from convert.From

	switch reqType {
	case pb.Type_TYPE_UNSPECIFIED:
		return nil, status.Errorf(codes.InvalidArgument, "type is missing")
	case pb.Type_NG_LOCALISE:
		from = convert.FromNgJson
	case pb.Type_NGX_TRANSLATE:
		from = convert.FromNgxTranslate
	case pb.Type_GO:
		from = convert.FromGo
	}

	messages, err := from([]byte(reqData))
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "parse data: %s", err)
	}

	// Save to DB/FS...
	messages.Language = language
	messages.Labels = reqLabels
	_ = messages

	return &pb.UploadTranslationFileResponse{}, nil
}
