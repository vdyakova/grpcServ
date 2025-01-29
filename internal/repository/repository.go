package repository

import (
	"context"
	"grpcserv/internal/model"
)

type FileServiceRepository interface {
	UploadFile(ctx context.Context, file *model.FileChunk) (*model.Message, error)
	ListFile(ctx context.Context) ([]*model.File, error)
	DownloadFile(ctx context.Context, fileName string) ([]*model.FileChunk, error)
}
