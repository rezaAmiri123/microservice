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

// MessageServiceClient is the client API for MessageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessageServiceClient interface {
	CreateEmail(ctx context.Context, in *CreateEmailRequest, opts ...grpc.CallOption) (*CreateEmailResponse, error)
	GetEmailByUUID(ctx context.Context, in *GetEmailByUUIDRequest, opts ...grpc.CallOption) (*GetEmailByUUIDResponse, error)
	GetEmails(ctx context.Context, in *GetEmailsRequest, opts ...grpc.CallOption) (*GetEmailsResponse, error)
}

type messageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMessageServiceClient(cc grpc.ClientConnInterface) MessageServiceClient {
	return &messageServiceClient{cc}
}

func (c *messageServiceClient) CreateEmail(ctx context.Context, in *CreateEmailRequest, opts ...grpc.CallOption) (*CreateEmailResponse, error) {
	out := new(CreateEmailResponse)
	err := c.cc.Invoke(ctx, "/message.MessageService/CreateEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) GetEmailByUUID(ctx context.Context, in *GetEmailByUUIDRequest, opts ...grpc.CallOption) (*GetEmailByUUIDResponse, error) {
	out := new(GetEmailByUUIDResponse)
	err := c.cc.Invoke(ctx, "/message.MessageService/GetEmailByUUID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) GetEmails(ctx context.Context, in *GetEmailsRequest, opts ...grpc.CallOption) (*GetEmailsResponse, error) {
	out := new(GetEmailsResponse)
	err := c.cc.Invoke(ctx, "/message.MessageService/GetEmails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessageServiceServer is the server API for MessageService service.
// All implementations must embed UnimplementedMessageServiceServer
// for forward compatibility
type MessageServiceServer interface {
	CreateEmail(context.Context, *CreateEmailRequest) (*CreateEmailResponse, error)
	GetEmailByUUID(context.Context, *GetEmailByUUIDRequest) (*GetEmailByUUIDResponse, error)
	GetEmails(context.Context, *GetEmailsRequest) (*GetEmailsResponse, error)
	mustEmbedUnimplementedMessageServiceServer()
}

// UnimplementedMessageServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMessageServiceServer struct {
}

func (UnimplementedMessageServiceServer) CreateEmail(context.Context, *CreateEmailRequest) (*CreateEmailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEmail not implemented")
}
func (UnimplementedMessageServiceServer) GetEmailByUUID(context.Context, *GetEmailByUUIDRequest) (*GetEmailByUUIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEmailByUUID not implemented")
}
func (UnimplementedMessageServiceServer) GetEmails(context.Context, *GetEmailsRequest) (*GetEmailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEmails not implemented")
}
func (UnimplementedMessageServiceServer) mustEmbedUnimplementedMessageServiceServer() {}

// UnsafeMessageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessageServiceServer will
// result in compilation errors.
type UnsafeMessageServiceServer interface {
	mustEmbedUnimplementedMessageServiceServer()
}

func RegisterMessageServiceServer(s grpc.ServiceRegistrar, srv MessageServiceServer) {
	s.RegisterService(&MessageService_ServiceDesc, srv)
}

func _MessageService_CreateEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).CreateEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.MessageService/CreateEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).CreateEmail(ctx, req.(*CreateEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_GetEmailByUUID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEmailByUUIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).GetEmailByUUID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.MessageService/GetEmailByUUID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).GetEmailByUUID(ctx, req.(*GetEmailByUUIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_GetEmails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEmailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).GetEmails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.MessageService/GetEmails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).GetEmails(ctx, req.(*GetEmailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MessageService_ServiceDesc is the grpc.ServiceDesc for MessageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MessageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "message.MessageService",
	HandlerType: (*MessageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEmail",
			Handler:    _MessageService_CreateEmail_Handler,
		},
		{
			MethodName: "GetEmailByUUID",
			Handler:    _MessageService_GetEmailByUUID_Handler,
		},
		{
			MethodName: "GetEmails",
			Handler:    _MessageService_GetEmails_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message.proto",
}
