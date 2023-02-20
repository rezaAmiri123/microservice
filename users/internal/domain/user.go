package domain

import (
	"github.com/rezaAmiri123/microservice/pkg/ddd"
	"github.com/stackus/errors"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var (
	ErrIDCannotBeBlank       = errors.Wrap(errors.ErrBadRequest, "the id cannot be blank")
	ErrUsernameCannotBeBlank = errors.Wrap(errors.ErrBadRequest, "the username cannot be blank")
	ErrPasswordCannotBeBlank = errors.Wrap(errors.ErrBadRequest, "the password cannot be blank")
	ErrEmailCannotBeBlank    = errors.Wrap(errors.ErrBadRequest, "the email cannot be blank")
	ErrUserAlreadyEnabled    = errors.Wrap(errors.ErrBadRequest, "the user is already enabled")
	ErrUserAlreadyDisabled   = errors.Wrap(errors.ErrBadRequest, "the user is already disabled")
)

// User is user model
type User struct {
	ddd.AggregateBase
	//ID        string    `json:"id" db:"id"`
	Username  string    `json:"username" db:"username" validate:"required,min=6,max=30"`
	Password  string    `json:"password" db:"password" validate:"required,min=8,max=15"`
	Email     string    `json:"email" db:"email" validate:"required,min=3,max=250,email"`
	Enabled   bool      `json:"enabled" db:"enabled"`
	Bio       string    `json:"bio" db:"bio"`
	Image     string    `json:"image" db:"image"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func NewUser(id string) *User {
	return &User{
		AggregateBase: ddd.AggregateBase{
			ID: id,
		},
	}
}
func RegisterUser(id, username, password, email string) (*User, error) {
	if id == "" {
		return nil, ErrIDCannotBeBlank
	}
	if username == "" {
		return nil, ErrUsernameCannotBeBlank
	}
	if password == "" {
		return nil, ErrPasswordCannotBeBlank
	}
	if email == "" {
		return nil, ErrEmailCannotBeBlank
	}
	user := NewUser(id)
	user.Username = username
	user.Password = password
	user.Email = email
	user.Enabled = true

	if err := user.hashPassword(); err != nil {
		return nil, errors.ErrInternal.Msgf("failed to hash password: %w", err.Error())
	}
	user.AddEvent(&UserRegistered{
		User: user,
	})
	return user, nil
}

func (u *User) Enable() error {
	if u.Enabled {
		return ErrUserAlreadyEnabled
	}

	u.Enabled = true

	u.AddEvent(&UserEnabled{
		User: u,
	})
	return nil
}
func (u *User) Disable() error {
	if !u.Enabled {
		return ErrUserAlreadyDisabled
	}

	u.Enabled = false

	u.AddEvent(&UserDisabled{
		User: u,
	})

	return nil
}

func (u *User) hashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}
