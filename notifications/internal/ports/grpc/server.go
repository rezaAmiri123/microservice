package grpc

import (
	"github.com/rezaAmiri123/microservice/notifications/internal/app"
	"github.com/rezaAmiri123/microservice/notifications/notificationspb"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"google.golang.org/grpc"
)

type (
	Config struct {
		App    app.App
		Logger logger.Logger
	}
	server struct {
		cfg *Config
		notificationspb.UnimplementedNotificationsServiceServer
	}
)

var _ notificationspb.NotificationsServiceServer = (*server)(nil)

func RegisterServer(application app.App, registrar grpc.ServiceRegistrar, logger logger.Logger) error {
	cfg := &Config{
		App:    application,
		Logger: logger,
	}
	notificationspb.RegisterNotificationsServiceServer(registrar, server{cfg: cfg})
	return nil
}
