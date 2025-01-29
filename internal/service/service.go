package service

import (
	"context"
	"grpcserv/internal/model"
)

type FileService interface {
	UploadFile(ctx context.Context, fileCh *model.FileChunk) (*model.Message, error)
	ListFile(ctx context.Context) ([]*model.File, error)
	DownloadFile(ctx context.Context, message string) ([]*model.FileChunk, error)
}
