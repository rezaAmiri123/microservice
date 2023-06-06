package agent

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rezaAmiri123/microservice/ordering/internal/constants"
	"github.com/rezaAmiri123/microservice/ordering/internal/ports/rest"
	"github.com/rezaAmiri123/microservice/pkg/di"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/rezaAmiri123/microservice/pkg/web"
	"net/http"
)

func (a *Agent) setupHttpServer() error {
	mux := chi.NewMux()
	mux.Use(middleware.Heartbeat("/liveness"))
	mux.Method("GET", "/metrics", promhttp.Handler())
	mux.Mount("/", http.FileServer(http.FS(web.WebUI)))
	//a.setupSwagger(mux)
	if err := rest.RegisterSwagger(mux); err != nil {
		return err
	}
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
