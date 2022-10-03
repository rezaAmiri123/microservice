package grpc

import (
	"context"
	"crypto/tls"
	"time"

	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
)

const (
	backoffLinear  = 100 * time.Millisecond
	backoffRetries = 3
)

func NewGrpcClient(
	addr string,
	clientTLSConfig *tls.Config,
	logger logger.Logger,
) (grpc.ClientConnInterface, error) {
	im := NewInterceptorManager(logger)

	ctx := context.Background()

	retryOpts := []grpc_retry.CallOption{
		grpc_retry.WithBackoff(grpc_retry.BackoffLinear(backoffLinear)),
		grpc_retry.WithCodes(codes.NotFound, codes.Aborted),
		grpc_retry.WithMax(backoffRetries),
	}

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithChainUnaryInterceptor(
		grpc_retry.UnaryClientInterceptor(retryOpts...),
		im.ClientRequestLoggerInterceptor(),
	))
	if clientTLSConfig != nil {
		clientCreds := credentials.NewTLS(clientTLSConfig)
		opts = append(opts, grpc.WithTransportCredentials(clientCreds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}

	return grpc.DialContext(ctx, addr, opts...)
}
