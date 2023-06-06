package rest

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rezaAmiri123/microservice/ordering/orderingpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func RegisterGateway(ctx context.Context, mux *chi.Mux, grpcAddr string) error {
	gateway := runtime.NewServeMux()
	err := orderingpb.RegisterOrderingServiceHandlerFromEndpoint(ctx, gateway, grpcAddr, []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	})
	if err != nil {
		return err
	}

	// mount the GRPC gateway
	mux.Mount("/v1", gateway)

	return nil
}
