package file_service

import (
	"context"
	"errors"
	"fmt"
	"grpcserv/internal/model"
	"grpcserv/internal/repository"
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"
)

type repo struct {
	uploadDir string
	mu        sync.RWMutex
}

func NewFileServiceRepository(uploadDir string) repository.FileServiceRepository {
	return &repo{
		uploadDir: uploadDir,
	}
}
func (r *repo) UploadFile(ctx context.Context, file *model.FileChunk) (*model.Message, error) {
	filePath := filepath.Join(r.uploadDir, file.Name)

	r.mu.Lock()
	defer r.mu.Unlock()
	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("upload canceled: %w", ctx.Err())
	default:
	}
	f, err := os.Create(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to create file: %w", err)
	}
	defer f.Close()
	_, err = f.Write(file.Content)
	if err != nil {
		return nil, fmt.Errorf("failed to write file content: %w", err)
	}

	return &model.Message{Message: fmt.Sprintf("File %s uploaded successfully", file.Name)}, nil
}
func (r *repo) ListFile(ctx context.Context) ([]*model.File, error) {
	var files []*model.File
	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("upload canceled: %w", ctx.Err())
	default:
	}
	err := filepath.Walk(r.uploadDir, func(path string, info os.FileInfo, err error) error {
		select {
		case <-ctx.Done():
			return fmt.Errorf("listfile canceled: %w", ctx.Err())
		default:
		}
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		files = append(files, &model.File{
			Name:      info.Name(),
			CreatedAt: info.ModTime(),
			UpdatedAt: info.ModTime(),
		})

		return nil
	})

	if err != nil {
		log.Printf("Error with skanning dir: %v", err)
		return []*model.File{}, fmt.Errorf("Sorry")
	}

	if len(files) == 0 {
		return []*model.File{}, fmt.Errorf("Sorry")
	}

	return files, nil
}

func (r *repo) DownloadFile(ctx context.Context, fileName string) ([]*model.FileChunk, error) {
	filePath := filepath.Join(r.uploadDir, fileName)
	r.mu.Lock()
	defer r.mu.Unlock()
	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("upload canceled: %w", ctx.Err())
	default:
	}
	f, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer f.Close()

	const chunkSize = 1024
	var fileChunks []*model.FileChunk
	buffer := make([]byte, chunkSize)
	for {
		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("download canceled: %w", ctx.Err())
		default:
		}
		n, err := f.Read(buffer)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return nil, fmt.Errorf("failed to read file: %w", err)
		}
		if n == 0 {
			break
		}

		fileChunk := &model.FileChunk{
			Name:    fileName,
			Content: buffer[:n],
		}
		fileChunks = append(fileChunks, fileChunk)
	}

	return fileChunks, nil
}
