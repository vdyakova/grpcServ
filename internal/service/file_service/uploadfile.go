package file_service

import (
	"context"
	"fmt"
	"grpcserv/internal/model"
)

func (s *serv) UploadFile(ctx context.Context, fileCh *model.FileChunk) (*model.Message, error) {
	message, err := s.fileRepository.UploadFile(ctx, fileCh)
	if err != nil {

		fmt.Println("Error service")
	}
	return message, nil
}
