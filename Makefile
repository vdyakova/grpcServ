include .env
generate-note-api:

	protoc --proto_path=api \
		   --go_out=pkg/file_service --go_opt=paths=source_relative \
 		  --go-grpc_out=pkg/file_service --go-grpc_opt=paths=source_relative \
 		  api/file_service.proto