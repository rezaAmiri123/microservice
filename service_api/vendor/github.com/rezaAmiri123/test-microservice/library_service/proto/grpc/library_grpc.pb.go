// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package grpc

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

// ArticleServiceClient is the client API for ArticleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ArticleServiceClient interface {
	GetArticleBySlug(ctx context.Context, in *GetArticleBySlugRequest, opts ...grpc.CallOption) (*GetArticleBySlugResponse, error)
	GetArticles(ctx context.Context, in *GetArticlesRequest, opts ...grpc.CallOption) (*GetArticlesResponse, error)
}

type articleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewArticleServiceClient(cc grpc.ClientConnInterface) ArticleServiceClient {
	return &articleServiceClient{cc}
}

func (c *articleServiceClient) GetArticleBySlug(ctx context.Context, in *GetArticleBySlugRequest, opts ...grpc.CallOption) (*GetArticleBySlugResponse, error) {
	out := new(GetArticleBySlugResponse)
	err := c.cc.Invoke(ctx, "/library.ArticleService/GetArticleBySlug", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleServiceClient) GetArticles(ctx context.Context, in *GetArticlesRequest, opts ...grpc.CallOption) (*GetArticlesResponse, error) {
	out := new(GetArticlesResponse)
	err := c.cc.Invoke(ctx, "/library.ArticleService/GetArticles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArticleServiceServer is the server API for ArticleService service.
// All implementations must embed UnimplementedArticleServiceServer
// for forward compatibility
type ArticleServiceServer interface {
	GetArticleBySlug(context.Context, *GetArticleBySlugRequest) (*GetArticleBySlugResponse, error)
	GetArticles(context.Context, *GetArticlesRequest) (*GetArticlesResponse, error)
	mustEmbedUnimplementedArticleServiceServer()
}

// UnimplementedArticleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedArticleServiceServer struct {
}

func (UnimplementedArticleServiceServer) GetArticleBySlug(context.Context, *GetArticleBySlugRequest) (*GetArticleBySlugResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticleBySlug not implemented")
}
func (UnimplementedArticleServiceServer) GetArticles(context.Context, *GetArticlesRequest) (*GetArticlesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticles not implemented")
}
func (UnimplementedArticleServiceServer) mustEmbedUnimplementedArticleServiceServer() {}

// UnsafeArticleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ArticleServiceServer will
// result in compilation errors.
type UnsafeArticleServiceServer interface {
	mustEmbedUnimplementedArticleServiceServer()
}

func RegisterArticleServiceServer(s grpc.ServiceRegistrar, srv ArticleServiceServer) {
	s.RegisterService(&ArticleService_ServiceDesc, srv)
}

func _ArticleService_GetArticleBySlug_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArticleBySlugRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).GetArticleBySlug(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/library.ArticleService/GetArticleBySlug",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).GetArticleBySlug(ctx, req.(*GetArticleBySlugRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleService_GetArticles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArticlesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).GetArticles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/library.ArticleService/GetArticles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).GetArticles(ctx, req.(*GetArticlesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ArticleService_ServiceDesc is the grpc.ServiceDesc for ArticleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ArticleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "library.ArticleService",
	HandlerType: (*ArticleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetArticleBySlug",
			Handler:    _ArticleService_GetArticleBySlug_Handler,
		},
		{
			MethodName: "GetArticles",
			Handler:    _ArticleService_GetArticles_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "library.proto",
}