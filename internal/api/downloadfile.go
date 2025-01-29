package api

import (
	"context"
	desc "grpcserv/pkg/file_service"
)

func (i *Implementation) DownloadFiles(ctx context.Context, request desc.FileRequest) (stream desc.FileServ_DownloadFileServer) {
	files, err := i.fileServ.DownloadFile(ctx, request.Filename)
	if err != nil {
		return nil

	}

	for _, fileChunk := range files {
		err := stream.Send(&desc.FileChunk{
			Filename: fileChunk.Name,
			Content:  fileChunk.Content,
		})
		if err != nil {
			return nil
		}
	}

	return nil
}
