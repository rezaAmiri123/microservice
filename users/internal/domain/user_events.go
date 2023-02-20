package domain

type UserRegistered struct {
	User *User
}

func (UserRegistered) EventName() string { return "users.UserRegistered" }

type UserEnabled struct {
	User *User
}

func (UserEnabled) EventName() string { return "users.UserEnabled" }

type UserDisabled struct {
	User *User
}

func (UserDisabled) EventName() string { return "users.UserDisabled" }
