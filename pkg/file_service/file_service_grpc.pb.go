// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.0
// source: file_service.proto

package file_service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	FileServ_UploadFile_FullMethodName   = "/file_service.FileServ/UploadFile"
	FileServ_ListFile_FullMethodName     = "/file_service.FileServ/ListFile"
	FileServ_DownloadFile_FullMethodName = "/file_service.FileServ/DownloadFile"
)

// FileServClient is the client API for FileServ service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FileServClient interface {
	UploadFile(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[FileChunk, FileUploadResponse], error)
	ListFile(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*FileListResponse, error)
	DownloadFile(ctx context.Context, in *FileRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[FileChunk], error)
}

type fileServClient struct {
	cc grpc.ClientConnInterface
}

func NewFileServClient(cc grpc.ClientConnInterface) FileServClient {
	return &fileServClient{cc}
}

func (c *fileServClient) UploadFile(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[FileChunk, FileUploadResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &FileServ_ServiceDesc.Streams[0], FileServ_UploadFile_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[FileChunk, FileUploadResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type FileServ_UploadFileClient = grpc.ClientStreamingClient[FileChunk, FileUploadResponse]

func (c *fileServClient) ListFile(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*FileListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FileListResponse)
	err := c.cc.Invoke(ctx, FileServ_ListFile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServClient) DownloadFile(ctx context.Context, in *FileRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[FileChunk], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &FileServ_ServiceDesc.Streams[1], FileServ_DownloadFile_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[FileRequest, FileChunk]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type FileServ_DownloadFileClient = grpc.ServerStreamingClient[FileChunk]

// FileServServer is the server API for FileServ service.
// All implementations must embed UnimplementedFileServServer
// for forward compatibility.
type FileServServer interface {
	UploadFile(grpc.ClientStreamingServer[FileChunk, FileUploadResponse]) error
	ListFile(context.Context, *emptypb.Empty) (*FileListResponse, error)
	DownloadFile(*FileRequest, grpc.ServerStreamingServer[FileChunk]) error
	mustEmbedUnimplementedFileServServer()
}

// UnimplementedFileServServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedFileServServer struct{}

func (UnimplementedFileServServer) UploadFile(grpc.ClientStreamingServer[FileChunk, FileUploadResponse]) error {
	return status.Errorf(codes.Unimplemented, "method UploadFile not implemented")
}
func (UnimplementedFileServServer) ListFile(context.Context, *emptypb.Empty) (*FileListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFile not implemented")
}
func (UnimplementedFileServServer) DownloadFile(*FileRequest, grpc.ServerStreamingServer[FileChunk]) error {
	return status.Errorf(codes.Unimplemented, "method DownloadFile not implemented")
}
func (UnimplementedFileServServer) mustEmbedUnimplementedFileServServer() {}
func (UnimplementedFileServServer) testEmbeddedByValue()                  {}

// UnsafeFileServServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FileServServer will
// result in compilation errors.
type UnsafeFileServServer interface {
	mustEmbedUnimplementedFileServServer()
}

func RegisterFileServServer(s grpc.ServiceRegistrar, srv FileServServer) {
	// If the following call pancis, it indicates UnimplementedFileServServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&FileServ_ServiceDesc, srv)
}

func _FileServ_UploadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FileServServer).UploadFile(&grpc.GenericServerStream[FileChunk, FileUploadResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type FileServ_UploadFileServer = grpc.ClientStreamingServer[FileChunk, FileUploadResponse]

func _FileServ_ListFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServServer).ListFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FileServ_ListFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServServer).ListFile(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileServ_DownloadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FileRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FileServServer).DownloadFile(m, &grpc.GenericServerStream[FileRequest, FileChunk]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type FileServ_DownloadFileServer = grpc.ServerStreamingServer[FileChunk]

// FileServ_ServiceDesc is the grpc.ServiceDesc for FileServ service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FileServ_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "file_service.FileServ",
	HandlerType: (*FileServServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListFile",
			Handler:    _FileServ_ListFile_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadFile",
			Handler:       _FileServ_UploadFile_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "DownloadFile",
			Handler:       _FileServ_DownloadFile_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "file_service.proto",
}
