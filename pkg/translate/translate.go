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

const chunkSize = 1024 * 3

func (t *TranslateServiceServer) DownloadTranslationFile(
	req *pb.DownloadTranslationFileRequest,
	stream pb.TranslateService_DownloadTranslationFileServer,
) error {
	chunk := &pb.DownloadTranslationFileResponse{Data: make([]byte, chunkSize)}

	// 1. find file from DB/FS
	// 2. transform to the label's structure

	// Placeholder for real data
	translation := []byte("ewoiYSI6ImIiLAoiYS5jIiA6ICJkIiwKInZsYWRpc2xhdnMiOiJwZXJrYW5rcyIKfQ==")

	reader := bytes.NewReader(translation)

	for {
		n, err := reader.Read(chunk.Data)

		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			return fmt.Errorf("stream chunk loop: %w", err)
		}

		chunk.Data = translation[:n]
		if err := stream.Send(chunk); err != nil {
			return fmt.Errorf("send response: %w", err)
		}
	}

	return nil
}
