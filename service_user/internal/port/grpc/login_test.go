package grpc_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	userService "github.com/rezaAmiri123/microservice/service_user/proto/grpc"
	"github.com/stretchr/testify/require"
)

func TestUserGRPCServer_Login(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// u, password := randomUser(t)
	u, _ := randomUser(t)
	testServer := NewTestGrpcServer(t, ctrl)
	testCases := []struct {
		name          string
		body          *userService.LoginUserRequest
		buildStubs    func(server *TestGrpcServer)
		checkResponse func(res *userService.LoginUserResponse, err error)
	}{
		{
			name: "OK",
			body: &userService.LoginUserRequest{
				Username: u.Username,
				Password: u.Password,
			},
			buildStubs: func(server *TestGrpcServer) {

				server.repoMock.EXPECT().GetUserByUsername(gomock.Any(), gomock.Eq(u.Username)).
					Times(1).
					Return(&u, nil)
				server.repoMock.EXPECT().
					CreateSession(gomock.Any(), gomock.Any()).
					Times(1)
			},
			checkResponse: func(res *userService.LoginUserResponse, err error) {
				require.Nil(t, err)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			tc.buildStubs(testServer)
			ctx := context.Background()
			res, err := testServer.grpcServer.Login(ctx, tc.body)
			tc.checkResponse(res, err)
		})
	}
}
