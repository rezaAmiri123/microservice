package grpc_test

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	pkgUtils "github.com/rezaAmiri123/microservice/pkg/utils"
	"github.com/rezaAmiri123/microservice/service_user/internal/domain/user"
	"github.com/rezaAmiri123/microservice/service_user/internal/utils"
	userService "github.com/rezaAmiri123/microservice/service_user/proto/grpc"
	"github.com/stretchr/testify/require"
)

type eqCreateUserParamsMatcher struct {
	arg      *user.CreateUserParams
	password string
}

func (e *eqCreateUserParamsMatcher) Matches(x interface{}) bool {
	arg, ok := x.(user.CreateUserParams)
	if !ok {
		return false
	}

	err := utils.CheckPassword(e.password, arg.Password)
	if err != nil {
		return false
	}

	e.arg.Password = arg.Password
	return reflect.DeepEqual(e.arg, arg)
}

func (e *eqCreateUserParamsMatcher) String() string {
	return fmt.Sprintf("matches arg %v and password %v", e.arg, e.password)
}

func EqCreateUserParams(arg *user.CreateUserParams, password string) gomock.Matcher {
	return &eqCreateUserParamsMatcher{arg, password}
}

func XTestUserGRPCServer_CreateUser(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	u, password := randomUser(t)
	testServer := NewTestGrpcServer(t, ctrl)
	// userReq := &userService.CreateUserRequest{
	// 	Username: pkgUtils.RandomCharacter(),
	// 	Password: pkgUtils.RandomString(10),
	// 	Email:    pkgUtils.RandomEmail(),
	// 	Bio:      pkgUtils.RandomString(15),
	// 	Image:    pkgUtils.RandomString(15),
	// }
	testCases := []struct {
		name          string
		body          *userService.CreateUserRequest
		buildStubs    func(server *TestGrpcServer)
		checkResponse func(res *userService.CreateUserResponse, err error)
	}{
		{
			name: "OK",
			body: &userService.CreateUserRequest{
				Username: u.Username,
				Password: u.Password,
				Email:    u.Email,
				Bio:      u.Bio,
				Image:    u.Image,
			},
			buildStubs: func(server *TestGrpcServer) {
				arg := &user.CreateUserParams{
					Username: u.Username,
					Email:    u.Email,
					Bio:      u.Bio,
					Image:    u.Image,
				}
				server.repoMock.EXPECT().CreateUser(gomock.Any(), EqCreateUserParams(arg, password)).
					Times(1).
					Return(&u, nil)
			},
			checkResponse: func(res *userService.CreateUserResponse, err error) {
				require.Nil(t, err)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			tc.buildStubs(testServer)
			ctx := context.Background()
			res, err := testServer.grpcServer.CreateUser(ctx, tc.body)
			tc.checkResponse(res, err)
		})
	}
}

func randomUser(t *testing.T) (user.User, string) {
	t.Helper()

	password := pkgUtils.RandomString(6)
	hashPassword, err := utils.HashPassword(password)
	require.NoError(t, err)

	u := user.User{
		Username: pkgUtils.RandomCharacter(),
		Password: hashPassword,
		Email:    pkgUtils.RandomEmail(),
		Bio:      pkgUtils.RandomString(15),
		Image:    pkgUtils.RandomString(15),
	}
	return u, password
}
