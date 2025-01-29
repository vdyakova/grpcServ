package api

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	desc "grpcserv/pkg/file_service"
)

func (i *Implementation) ListFile(ctx context.Context, a *emptypb.Empty) (response *desc.FileListResponse, err error) {
	files, err := i.fileServ.ListFile(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list files: %w", err)
	}

	var fileMetadataList []*desc.FileMetadata
	for _, file := range files {
		fileMetadataList = append(fileMetadataList, &desc.FileMetadata{
			Name:      file.Name,
			CreatedAt: timestamppb.New(file.CreatedAt),
			UpdatedAt: timestamppb.New(file.UpdatedAt),
		})
	}

	response = &desc.FileListResponse{
		Files: fileMetadataList,
	}
	return response, nil
}
