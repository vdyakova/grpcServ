package api

import (
	"grpcserv/internal/service"
	desc "grpcserv/pkg/file_service"
)

type Implementation struct {
	desc.UnimplementedFileServServer
	fileServ service.FileService
}

func NewImplementation(fileServ service.FileService) *Implementation {
	return &Implementation{
		fileServ: fileServ,
	}
}
