package api

import (
	"context"
	"fmt"
	"grpcserv/internal/model"
	desc "grpcserv/pkg/file_service"
	"io"
)

func (i *Implementation) UploadFile(stream desc.FileServ_UploadFileServer) error {
	var fileName string
	ctx := context.Background()
	for {
		chunk, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("failed to receive chunk: %w", err)
		}

		if fileName == "" {
			fileName = chunk.Filename
		}
		_, err = i.fileServ.UploadFile(ctx, &model.FileChunk{
			Name:    chunk.GetFilename(),
			Content: chunk.GetContent(),
		})
		if err != nil {
			return fmt.Errorf("failed to save chunk: %w", err)
		}
	}
	response := &desc.FileUploadResponse{
		Message: fmt.Sprintf("File %s uploaded successfully", fileName),
	}

	if err := stream.SendAndClose(response); err != nil {
		return fmt.Errorf("failed to send response: %w", err)
	}
	return nil

}
