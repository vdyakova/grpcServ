package app

import (
	"context"
	"google.golang.org/grpc"
	"grpcserv/internal/api"
	"grpcserv/internal/config"
	"grpcserv/internal/repository"
	serviceRepository "grpcserv/internal/repository/file_service"
	"grpcserv/internal/service"
	fileService "grpcserv/internal/service/file_service"
	"log"
	"sync"
)

type serviceProvider struct {
	mu                sync.Mutex
	grpcConfig        config.GRPCConfig
	serviceRepository repository.FileServiceRepository
	fileService       service.FileService
	activeConnections int32
	serviceImpl       *api.Implementation
	uploadLimiter     chan struct{}
	listFileLimiter   chan struct{}
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{
		uploadLimiter:   make(chan struct{}, 10),
		listFileLimiter: make(chan struct{}, 100),
	}
}
func (s *serviceProvider) connectionInterceptor(server interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {

	var limiter chan struct{}

	switch info.FullMethod {
	case "/fileService.UploadFile", "/fileService.DownloadFile":
		limiter = s.uploadLimiter
	case "/fileService.ListFiles":
		limiter = s.listFileLimiter
	default:
		limiter = nil
	}

	if limiter != nil {
		limiter <- struct{}{}
		log.Printf("Limiting concurrent requests for %s: %d active requests", info.FullMethod, len(limiter))
	}

	err := handler(server, stream)

	if limiter != nil {
		<-limiter
	}

	return err
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}
func (s *serviceProvider) ServiceRepository(ctx context.Context) repository.FileServiceRepository {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.serviceRepository == nil {

		s.serviceRepository = serviceRepository.NewFileServiceRepository("/upload")

	}

	return s.serviceRepository
}

func (s *serviceProvider) FileService(ctx context.Context) service.FileService {
	if s.fileService == nil {
		s.fileService = fileService.NewService(
			s.ServiceRepository(ctx),
		)
	}

	return s.fileService
}

func (s *serviceProvider) ServiceImpl(ctx context.Context) *api.Implementation {
	if s.serviceImpl == nil {
		fs := s.FileService(ctx)
		if fs == nil {
			log.Println("Error: FileService() returned nil")
			return nil
		}

		s.serviceImpl = api.NewImplementation(fs)
		if s.serviceImpl == nil {
			log.Println("Error: NewImplementation() returned nil")
		}
	}

	return s.serviceImpl
}
