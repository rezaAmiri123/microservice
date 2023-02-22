package domain

const (
	UserRegisteredEvent = "users.UserRegistered"
	UserEnabledEvent    = "users.UserEnabled"
	UserDisabledEvent   = "users.UserDisabled"
)

type UserRegistered struct {
	User *User
}

func (UserRegistered) EventName() string { return UserRegisteredEvent }

type UserEnabled struct {
	User *User
}

func (UserEnabled) EventName() string { return UserEnabledEvent }

type UserDisabled struct {
	User *User
}

func (UserDisabled) EventName() string { return UserDisabledEvent }
