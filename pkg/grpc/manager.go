package grpc

import (
	"context"
	"time"

	"github.com/rezaAmiri123/microservice/pkg/auth"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/rezaAmiri123/microservice/pkg/token"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type InterceptorManager interface {
	Logger(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error)
	ClientRequestLoggerInterceptor() func(
		ctx context.Context,
		method string,
		req interface{},
		reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error
	Payload(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error)
	ClientRequestPayloadInterceptor() func(
		ctx context.Context,
		method string,
		req interface{},
		reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error
}

// InterceptorManager struct
type interceptorManager struct {
	logger logger.Logger
}

// NewInterceptorManager InterceptorManager constructor
func NewInterceptorManager(logger logger.Logger) *interceptorManager {
	return &interceptorManager{logger: logger}
}

// Logger Interceptor
func (im *interceptorManager) Logger(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	start := time.Now()
	md, _ := metadata.FromIncomingContext(ctx)
	reply, err := handler(ctx, req)
	im.logger.GrpcMiddlewareAccessLogger(info.FullMethod, time.Since(start), md, err)
	return reply, err
}

// ClientRequestLoggerInterceptor gRPC client interceptor
func (im *interceptorManager) ClientRequestLoggerInterceptor() func(
	ctx context.Context,
	method string,
	req interface{},
	reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	return func(
		ctx context.Context,
		method string,
		req interface{},
		reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		start := time.Now()
		err := invoker(ctx, method, req, reply, cc, opts...)
		md, _ := metadata.FromIncomingContext(ctx)
		im.logger.GrpcClientInterceptorLogger(method, req, reply, time.Since(start), md, err)
		return err
	}
}

// Payload Interceptor
func (im *interceptorManager) Payload(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	md, _ := metadata.FromIncomingContext(ctx)
	payloadString := md.Get(auth.AuthorizationPayloadKey)[0]
	payload, _ := token.UnarshalPayload(payloadString)
	ctx = context.WithValue(ctx, auth.AuthorizationPayloadKey, &payload)
	reply, err := handler(ctx, req)
	return reply, err
}

// ClientRequestPayloadInterceptor gRPC client interceptor
func (im *interceptorManager) ClientRequestPayloadInterceptor() func(
	ctx context.Context,
	method string,
	req interface{},
	reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	return func(
		ctx context.Context,
		method string,
		req interface{},
		reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		payload := auth.PayloadFromCtx(ctx)
		ctx = metadata.AppendToOutgoingContext(ctx, auth.AuthorizationPayloadKey, payload.Marshal())
		err := invoker(ctx, method, req, reply, cc, opts...)
		return err
	}
}
