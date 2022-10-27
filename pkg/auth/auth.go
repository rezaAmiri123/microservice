package auth

import (
	"context"

	"github.com/rezaAmiri123/microservice/pkg/token"
)

const (
	AuthorizationHeaderKey  = "authorization"
	AuthorizationTypeBearer = "bearer"
	AuthorizationPayloadKey = "authorization_payload"
)

func PayloadFromCtx(ctx context.Context) *token.Payload {
	payload, ok := ctx.Value(AuthorizationPayloadKey).(*token.Payload)
	if ok {
		return payload
	}

	return &token.Payload{}
}
