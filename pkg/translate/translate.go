package translate

import (
	"bytes"
	"errors"
	"fmt"
	"io"

	pb "github.com/expect-digital/translate/pkg/server/translate/v1"
)

type TranslateServiceServer struct {
	pb.UnimplementedTranslateServiceServer
}

func New() *TranslateServiceServer {
	return new(TranslateServiceServer)
}

func (t *TranslateServiceServer) UploadTranslationFile(
	stream pb.TranslateService_UploadTranslationFileServer,
) error {
	var fileData bytes.Buffer

	for {
		req, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			return fmt.Errorf("stream chunk loop: %w", err)
		}

		chunk := req.GetData()
		_, err = fileData.Write(chunk)

		if err != nil {
			return fmt.Errorf("write to bytes buffer: %w", err)
		}
	}

	// 1. Check the label and transform accordingly...
	// 2. Save to DB/FS...

	if err := stream.SendAndClose(&pb.UploadTranslationFileResponse{}); err != nil {
		return fmt.Errorf("send and close: %w", err)
	}

	return nil
}
