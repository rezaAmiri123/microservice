package agent

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rezaAmiri123/microservice/pkg/web"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rezaAmiri123/microservice/baskets/internal/constants"
	"github.com/rezaAmiri123/microservice/baskets/internal/ports/rest"
	"github.com/rezaAmiri123/microservice/pkg/di"
	"github.com/rezaAmiri123/microservice/pkg/logger"
)

func (a *Agent) setupHttpServer() error {
	mux := chi.NewMux()
	mux.Use(middleware.Heartbeat("/liveness"))
	mux.Method("GET", "/metrics", promhttp.Handler())
	mux.Mount("/", http.FileServer(http.FS(web.WebUI)))
	//a.setupSwagger(mux)
	rest.RegisterSwagger(mux)
	grpcAddress := fmt.Sprintf("%s:%d", a.Config.GRPCServerAddr, a.Config.GRPCServerPort)
	if err := rest.RegisterGateway(context.Background(), mux, grpcAddress); err != nil {
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
		log := a.container.Get(constants.LoggerKey).(logger.Logger)
		log.Infof("run http at %s", webServer.Addr)
		err := webServer.ListenAndServe()
		if err != nil {
			_ = a.Shutdown()
		}
	}()
	return nil

}

//func (a *Agent) setupSwagger(mux *chi.Mux) {
//	const specRoot = "/baskets-spec/"
//	mux.Mount(specRoot, http.StripPrefix(specRoot, http.FileServer(http.FS(rest.SwaggerUI))))
//}

//func (a *Agent) setupGrpcEndpoint(mux *chi.Mux) error {
//	grpcAddress := fmt.Sprintf("%s:%d", a.Config.GRPCServerAddr, a.Config.GRPCServerPort)
//	gateway := runtime.NewServeMux()
//	err := basketspb.RegisterBasketServiceHandlerFromEndpoint(context.Background(), gateway, grpcAddress, []grpc.DialOption{
//		grpc.WithTransportCredentials(insecure.NewCredentials()),
//	})
//	if err != nil {
//		return err
//	}
//
//	// mount the GRPC gateway
//	mux.Mount("/v1", gateway)
//	return nil
//}
