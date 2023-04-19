package agent

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rezaAmiri123/microservice/pkg/di"
	"github.com/rezaAmiri123/microservice/pkg/web"
	"github.com/rezaAmiri123/microservice/stores/internal/constants"
	"github.com/rezaAmiri123/microservice/stores/internal/ports/rest"
	"github.com/rezaAmiri123/microservice/stores/storespb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (a *Agent) setupHttpServer() error {
	mux := chi.NewMux()
	a.setupSwagger(mux)
	if err := a.setupGrpcEndpoint(mux); err != nil {
		return err
	}
	webServer := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", a.Config.HttpServerAddr, a.Config.HttpServerPort),
		Handler: mux,
	}
	a.container.AddSingleton(constants.HttpServerKey, func(c di.Container) (any, error) {
		return webServer, nil
	})
	//a.httpServer = webServer
	go func() {
		err := webServer.ListenAndServe()
		if err != nil {
			_ = a.Shutdown()
		}
	}()
	return nil

}
func (a *Agent) setupSwagger(mux *chi.Mux) {
	mux.Use(middleware.Heartbeat("/liveness"))
	mux.Method("GET", "/metrics", promhttp.Handler())
	mux.Mount("/", http.FileServer(http.FS(web.WebUI)))
	const specRoot = "/stores-spec/"
	mux.Mount(specRoot, http.StripPrefix(specRoot, http.FileServer(http.FS(rest.SwaggerUI))))
}
func (a *Agent) setupGrpcEndpoint(mux *chi.Mux) error {
	grpcAddress := fmt.Sprintf("%s:%d", a.Config.GRPCServerAddr, a.Config.GRPCServerPort)
	gateway := runtime.NewServeMux()
	err := storespb.RegisterStoresServiceHandlerFromEndpoint(context.Background(), gateway, grpcAddress, []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	})
	if err != nil {
		return err
	}

	// mount the GRPC gateway
	mux.Mount("/v1", gateway)
	return nil
}
