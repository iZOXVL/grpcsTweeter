// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.27.0
// source: analysis.proto

package analysis

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	AnalysisService_ChatCompletions_FullMethodName = "/AnalysisService/chatCompletions"
)

// AnalysisServiceClient is the client API for AnalysisService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AnalysisServiceClient interface {
	ChatCompletions(ctx context.Context, in *AnalyzeRequest, opts ...grpc.CallOption) (*AnalyzeResponse, error)
}

type analysisServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAnalysisServiceClient(cc grpc.ClientConnInterface) AnalysisServiceClient {
	return &analysisServiceClient{cc}
}

func (c *analysisServiceClient) ChatCompletions(ctx context.Context, in *AnalyzeRequest, opts ...grpc.CallOption) (*AnalyzeResponse, error) {
	out := new(AnalyzeResponse)
	err := c.cc.Invoke(ctx, AnalysisService_ChatCompletions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AnalysisServiceServer is the server API for AnalysisService service.
// All implementations must embed UnimplementedAnalysisServiceServer
// for forward compatibility
type AnalysisServiceServer interface {
	ChatCompletions(context.Context, *AnalyzeRequest) (*AnalyzeResponse, error)
	mustEmbedUnimplementedAnalysisServiceServer()
}

// UnimplementedAnalysisServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAnalysisServiceServer struct {
}

func (UnimplementedAnalysisServiceServer) ChatCompletions(context.Context, *AnalyzeRequest) (*AnalyzeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChatCompletions not implemented")
}
func (UnimplementedAnalysisServiceServer) mustEmbedUnimplementedAnalysisServiceServer() {}

// UnsafeAnalysisServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AnalysisServiceServer will
// result in compilation errors.
type UnsafeAnalysisServiceServer interface {
	mustEmbedUnimplementedAnalysisServiceServer()
}

func RegisterAnalysisServiceServer(s grpc.ServiceRegistrar, srv AnalysisServiceServer) {
	s.RegisterService(&AnalysisService_ServiceDesc, srv)
}

func _AnalysisService_ChatCompletions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AnalyzeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalysisServiceServer).ChatCompletions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AnalysisService_ChatCompletions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalysisServiceServer).ChatCompletions(ctx, req.(*AnalyzeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AnalysisService_ServiceDesc is the grpc.ServiceDesc for AnalysisService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AnalysisService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "AnalysisService",
	HandlerType: (*AnalysisServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "chatCompletions",
			Handler:    _AnalysisService_ChatCompletions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "analysis.proto",
}
