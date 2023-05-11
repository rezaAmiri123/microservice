package app

import (
	"context"
	"fmt"
	"github.com/rezaAmiri123/microservice/pkg/ddd"
	"github.com/rezaAmiri123/microservice/users/internal/domain"
)

//type Application struct {
//	Commands Commands
//	Queries  Queries
//}
//
//type Queries struct {
//	AuthorizeUser *queries.AuthorizeUserHandler
//}
//
//type Commands struct {
//	RegisterUser *commands.RegisterUserHandler
//	EnableUser   *commands.EnableUserHandler
//}

type (
	RegisterUser struct {
		ID       string
		Username string
		Password string
		Email    string
	}

	AuthorizeUser struct {
		ID string
	}

	GetUser struct {
		ID string
	}

	EnableUser struct {
		ID string
	}

	DisableUser struct {
		ID string
	}
	App interface {
		RegisterUser(ctx context.Context, cmd RegisterUser) error
		AuthorizeUser(ctx context.Context, cmd AuthorizeUser) error
		GetUser(ctx context.Context, cmd GetUser) (*domain.User, error)
		EnableUser(ctx context.Context, cmd EnableUser) error
		DisableUser(ctx context.Context, cmd DisableUser) error
	}

	Application struct {
		users     domain.UserRepository
		publisher ddd.EventPublisher[ddd.AggregateEvent]
	}
)

var _ App = (*Application)(nil)

func New(
	users domain.UserRepository,
	publisher ddd.EventPublisher[ddd.AggregateEvent],
) Application {
	return Application{
		users:     users,
		publisher: publisher,
	}
}

func (a Application) RegisterUser(ctx context.Context, cmd RegisterUser) error {
	user, err := domain.RegisterUser(cmd.ID, cmd.Username, cmd.Password, cmd.Email)
	if err != nil {
		return err
	}
	fmt.Println("register user app:", user.Username)
	if err = a.users.Save(ctx, user); err != nil {
		return err
	}
	// publish domain events
	if err = a.publisher.Publish(ctx, user.Events()...); err != nil {
		return err
	}
	return nil
}
func (a Application) AuthorizeUser(ctx context.Context, cmd AuthorizeUser) error {
	user, err := a.users.Find(ctx, cmd.ID)
	if err != nil {
		return err
	}

	if err = user.Authorize(); err != nil {
		return err
	}

	// publish domain events
	if err = a.publisher.Publish(ctx, user.Events()...); err != nil {
		return err
	}

	return nil

}
func (a Application) GetUser(ctx context.Context, cmd GetUser) (*domain.User, error) {
	return a.users.Find(ctx, cmd.ID)
}
func (a Application) EnableUser(ctx context.Context, cmd EnableUser) error {
	user, err := a.users.Find(ctx, cmd.ID)
	if err != nil {
		return err
	}

	if err = user.Enable(); err != nil {
		return err
	}

	if err = a.users.Update(ctx, user); err != nil {
		return err
	}

	// publish domain events
	if err = a.publisher.Publish(ctx, user.Events()...); err != nil {
		return err
	}

	return nil

}
func (a Application) DisableUser(ctx context.Context, cmd DisableUser) error {
	user, err := a.users.Find(ctx, cmd.ID)
	if err != nil {
		return err
	}

	if err = user.Disable(); err != nil {
		return err
	}

	if err = a.users.Update(ctx, user); err != nil {
		return err
	}

	// publish domain events
	if err = a.publisher.Publish(ctx, user.Events()...); err != nil {
		return err
	}
	return nil
}
