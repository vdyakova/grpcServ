package file_service

import (
	"context"
	"grpcserv/internal/model"
)

func (s *serv) DownloadFile(ctx context.Context, message string) ([]*model.FileChunk, error) {

	files, err := s.fileRepository.DownloadFile(ctx, message)

	if err != nil {
		return nil, err
	}
	var messages []*model.FileChunk
	for _, file := range files {
		messages = append(messages, &model.FileChunk{
			Name:    file.Name,
			Content: file.Content,
		})
	}

	return messages, nil
}
