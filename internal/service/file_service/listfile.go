package file_service

import (
	"context"
	"fmt"
	"grpcserv/internal/model"
)

func (s *serv) ListFile(ctx context.Context) ([]*model.File, error) {

	files, err := s.fileRepository.ListFile(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list files: %w", err)
	}

	var messages []*model.File
	for _, file := range files {
		messages = append(messages, &model.File{
			Name:      file.Name,
			CreatedAt: file.CreatedAt,
			UpdatedAt: file.UpdatedAt,
		})
	}

	return messages, nil
}
