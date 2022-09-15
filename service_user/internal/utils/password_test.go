package utils_test

import (
	"testing"

	pkgutils "github.com/rezaAmiri123/microservice/pkg/utils"
	"github.com/rezaAmiri123/microservice/service_user/internal/utils"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	password := pkgutils.RandomString(6)

	hashedPassword1, err := utils.HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword1)

	err = utils.CheckPassword(password, hashedPassword1)
	require.NoError(t, err)

	wrongPassword := pkgutils.RandomString(6)
	err = utils.CheckPassword(wrongPassword, hashedPassword1)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	hashedPassword2, err := utils.HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword2)
	require.NotEqual(t, hashedPassword1, hashedPassword2)
}
