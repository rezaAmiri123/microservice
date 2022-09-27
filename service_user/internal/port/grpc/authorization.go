package grpc

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/rezaAmiri123/microservice/pkg/token"
	"google.golang.org/grpc/metadata"
)

const (
	authorizationHeader = "authorization"
	authorizationBearer = "bearer"
)

var (
	ErrMissMetadata            = errors.New("missing metadata")
	ErrMissAuthHeader          = errors.New("missing authorization header")
	ErrInvalidAuthHeaderFormat = errors.New("invalid authorization header format")
)

func (s *UserGRPCServer) AuthorizeUser(ctx context.Context) (*token.Payload, error) {
	// span, ctx := opentracing.StartSpanFromContext(ctx, "UserGRPCServer.authorizeUser")
	// defer span.Finish()

	// s.cfg.Metric.LoginRequests.Inc()

	// violations := validateLoginRequest(req)
	// if violations != nil {
	// 	return nil, pkgGrpc.InvalidArgumentError(violations)
	// }

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, ErrMissMetadata
	}

	valuse := md.Get(authorizationHeader)
	if len(valuse) == 0 {
		return nil, ErrMissAuthHeader
	}

	authHeader := valuse[0]
	fields := strings.Fields(authHeader)
	if len(fields) < 2 {
		return nil, ErrInvalidAuthHeaderFormat
	}

	authType := strings.ToLower(fields[0])
	if authType != authorizationBearer {
		return nil, fmt.Errorf("unsupported authorization type: %s", authType)
	}

	accessToken := fields[1]
	payload, err := s.cfg.Maker.VerifyToken(accessToken)
	if err != nil {
		return nil, fmt.Errorf("invalid access token: %s", err)
	}
	return payload, nil
}
