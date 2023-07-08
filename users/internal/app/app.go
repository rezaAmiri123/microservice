package app

import (
	"context"
	"github.com/rezaAmiri123/microservice/pkg/ddd"
	"github.com/rezaAmiri123/microservice/users/internal/domain"
	"github.com/stackus/errors"
)

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

func (a Application) RegisterUser(ctx context.Context, cmd RegisterUser) (err error) {
	defer func() {
		if err != nil {
			err = errors.Wrap(err, "app Application.RegisterUser")
		}
	}()

	user, err := domain.RegisterUser(cmd.ID, cmd.Username, cmd.Password, cmd.Email)
	if err != nil {
		return err
	}
	if err = a.users.Save(ctx, user); err != nil {
		return err
	}
	// publish domain events
	return a.publisher.Publish(ctx, user.Events()...)
}

func (a Application) AuthorizeUser(ctx context.Context, cmd AuthorizeUser) (err error) {
	defer func() {
		if err != nil {
			err = errors.Wrap(err, "app Application.AuthorizeUser")
		}
	}()

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
func (a Application) GetUser(ctx context.Context, cmd GetUser) (user *domain.User, err error) {
	defer func() {
		if err != nil {
			err = errors.Wrap(err, "app Application.GetUser")
		}
	}()

	return a.users.Find(ctx, cmd.ID)
}
func (a Application) EnableUser(ctx context.Context, cmd EnableUser) (err error) {
	defer func() {
		if err != nil {
			err = errors.Wrap(err, "app Application.EnableUser")
		}
	}()

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
func (a Application) DisableUser(ctx context.Context, cmd DisableUser) (err error) {
	defer func() {
		if err != nil {
			err = errors.Wrap(err, "app Application.DisableUser")
		}
	}()

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
