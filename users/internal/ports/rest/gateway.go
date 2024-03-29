package rest

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rezaAmiri123/microservice/users/userspb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func RegisterGateway(ctx context.Context, mux *chi.Mux, grpcAddr string) error {
	const apiRoot = "/api/users"

	gateway := runtime.NewServeMux()
	err := userspb.RegisterUserServiceHandlerFromEndpoint(ctx, gateway, grpcAddr, []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	})
	if err != nil {
		return err
	}

	// mount the GRPC gateway
	mux.Mount(apiRoot, gateway)

	return nil
}
