package jwt_test

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	pkgutils "github.com/rezaAmiri123/microservice/pkg/utils"
	"github.com/rezaAmiri123/microservice/service_user/pkg/token"
	jwtPkg "github.com/rezaAmiri123/microservice/service_user/pkg/token/jwt"
	"github.com/stretchr/testify/require"
)

func TestJWTMAker(t *testing.T) {
	maker, err := jwtPkg.NewJWTMaker(pkgutils.RandomString(32))
	require.NoError(t, err)

	username := pkgutils.RandomCharacter()
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, payload, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)

	payload, err = maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.NotZero(t, payload.ID)
	require.Equal(t, username, payload.Username)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}

func TestExpiredJWTMAker(t *testing.T) {
	maker, err := jwtPkg.NewJWTMaker(pkgutils.RandomString(32))
	require.NoError(t, err)

	jwtToken, payload, err := maker.CreateToken(pkgutils.RandomCharacter(), -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, jwtToken)
	require.NotEmpty(t, payload)

	payload, err = maker.VerifyToken(jwtToken)
	require.Error(t, err)
	require.EqualError(t, err, token.ErrExpiredToken.Error())
	require.Nil(t, payload)
}

func TestInvalidJWTTokenAlgNone(t *testing.T) {
	payload, err := token.NewPayload(pkgutils.RandomCharacter(), time.Minute)
	require.NoError(t, err)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, payload)
	resToken, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
	require.NoError(t, err)

	maker, err := jwtPkg.NewJWTMaker(pkgutils.RandomString(32))
	require.NoError(t, err)

	payload, err = maker.VerifyToken(resToken)
	require.Error(t, err)
	require.EqualError(t, err, token.ErrInvalidToken.Error())
	require.Nil(t, payload)
}
