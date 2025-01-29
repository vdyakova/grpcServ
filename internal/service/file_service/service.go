package file_service

import (
	"grpcserv/internal/repository"
	"grpcserv/internal/service"
)

type serv struct {
	fileRepository repository.FileServiceRepository
}

func NewService(fileRepository repository.FileServiceRepository) service.FileService {
	return &serv{
		fileRepository: fileRepository,
	}
}
