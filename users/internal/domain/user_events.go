package domain

const (
	UserRegisteredEvent = "users.UserRegistered"
	UserEnabledEvent    = "users.UserEnabled"
	UserDisabledEvent   = "users.UserDisabled"
	UserAuthorizedEvent = "users.UserAuthorized"
)

type UserRegistered struct {
	User *User
}

func (UserRegistered) Key() string { return UserRegisteredEvent }

type UserEnabled struct {
	User *User
}

func (UserEnabled) Key() string { return UserEnabledEvent }

type UserDisabled struct {
	User *User
}

func (UserDisabled) Key() string { return UserDisabledEvent }

type UserAuthorized struct {
	User *User
}

func (UserAuthorized) Key() string { return UserAuthorizedEvent }
