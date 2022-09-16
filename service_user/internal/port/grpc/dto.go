package grpc

import (
	"github.com/rezaAmiri123/microservice/service_user/internal/domain/user"
	userService "github.com/rezaAmiri123/microservice/service_user/proto/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func UserToGrpc(u *user.User) *userService.User {
	res := &userService.User{}
	res.UserUuid = u.UserID[:]
	res.Username = u.Username
	res.Email = u.Email
	res.Bio = u.Bio
	res.Image = u.Image
	res.CreatedAt = timestamppb.New(u.CreatedAt)
	res.UpdatedAt = timestamppb.New(u.UpdatedAt)
	return res
}
