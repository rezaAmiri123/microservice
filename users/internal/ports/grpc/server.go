package grpc

import (
	"database/sql"
	"fmt"
	"github.com/rezaAmiri123/microservice/pkg/di"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/rezaAmiri123/microservice/users/internal/app"
	"github.com/rezaAmiri123/microservice/users/internal/constants"
	"github.com/rezaAmiri123/microservice/users/internal/domain"
	"github.com/rezaAmiri123/microservice/users/userspb"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type (
	Config struct {
		App    app.App
		Logger logger.Logger
	}
	server struct {
		cfg *Config
		userspb.UnimplementedUserServiceServer
	}
	serverTx struct {
		c di.Container
		userspb.UnimplementedUserServiceServer
	}
)

var _ userspb.UserServiceServer = (*server)(nil)
var _ userspb.UserServiceServer = (*serverTx)(nil)

func RegisterServerTx(container di.Container, registrar grpc.ServiceRegistrar) error {
	userspb.RegisterUserServiceServer(registrar, serverTx{
		c: container,
	})
	return nil
}
func (s serverTx) getNextServer() server {
	cfg := &Config{
		App:    s.c.Get(constants.ApplicationKey).(app.App),
		Logger: s.c.Get(constants.LoggerKey).(logger.Logger),
	}
	return server{cfg: cfg}
}

//func NewServer(app app.App, logger logger.Logger) *server {
//	cfg := &Config{
//		App:    app,
//		Logger: logger,
//	}
//	return &server{cfg: cfg}
//}

func (s server) userFromDomain(user *domain.User) *userspb.User {
	return &userspb.User{
		UserUuid:  user.ID(),
		Username:  user.Username,
		Email:     user.Email,
		Enabled:   user.Enabled,
		Bio:       user.Bio,
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}
}
func (s serverTx) closeTx(tx *sql.Tx, err error) error {
	if p := recover(); p != nil {
		_ = tx.Rollback()
		fmt.Println("rollback")
		panic(p)
	} else if err != nil {
		_ = tx.Rollback()
		fmt.Println("rollback")
		return err
	} else {
		fmt.Println("commit")
		return tx.Commit()
	}
}
