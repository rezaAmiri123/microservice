///go:build e2e

package e2e

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/cucumber/godog"
	"github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/rezaAmiri123/microservice/users/usersclient"
	"github.com/rezaAmiri123/microservice/users/usersclient/models"
	"github.com/rezaAmiri123/microservice/users/usersclient/user"
	"github.com/stackus/errors"
)

type userIDKey struct{}

type usersFeature struct {
	client *usersclient.UserServiceAPI
	//db     *sql.DB
}

var _ feature = (*usersFeature)(nil)

func (u *usersFeature) getDB() (*sql.DB, error) {
	return sql.Open("pgx", "postgres://users_user:users_pass@localhost:5432/users?sslmode=disable&search_path=users,public")
}

func (u *usersFeature) init(cfg featureConfig) (err error) {
	//if cfg.useMonoDB {
	//	u.db, err = sql.Open("pgx", "postgres://mallbots_user:mallbots_pass@localhost:5432/mallbots?sslmode=disable")
	//} else {
	//	u.db, err = sql.Open("pgx", "postgres://users_user:users_pass@localhost:5432/users?sslmode=disable&search_path=users,public")
	//}
	//if err != nil {
	//	return
	//}
	conn := client.New("localhost:8080", "/", nil)
	u.client = usersclient.New(conn, strfmt.Default)
	return
}

func (u *usersFeature) register(ctx *godog.ScenarioContext) {
	ctx.Step(`^I am a registered user$`, u.iAmARegisteredUser)
	ctx.Step(`^I register a new user as "([^"]*)"$`, u.iRegisterANewUserAs)
	ctx.Step(`^(?:I )?(?:ensure |expect )?the user (?:was|is) created$`, u.expectTheUserWasCreated)
	ctx.Step(`^(?:I )?(?:ensure |expect )?a user named "([^"]*)" (?:to )?exists?$`, u.expectAUserNamedToExist)
	ctx.Step(`^(?:I )?(?:ensure |expect )?no user named "([^"]*)" (?:to )?exists?$`, u.expectNoUserNamedToExist)
}

func (u *usersFeature) reset() {
	db, _ := u.getDB()
	defer db.Close()

	truncate := func(tableName string) {
		_, _ = db.Exec(fmt.Sprintf("TRUNCATE %s", tableName))
	}

	truncate("users")
	truncate("inbox")
	truncate("outbox")
}

func (u *usersFeature) iAmARegisteredUser(ctx context.Context) context.Context {
	resp, err := u.client.User.RegisterUser(user.NewRegisterUserParams().WithBody(&models.UserspbRegisterUserRequest{
		Username: withRandomString("RegisteredUser"),
		Email:    withRandomString("RegisteredUserEmail"),
		Password: withRandomString("RegisteredUserPassword"),
	}))
	ctx = setLastResponseAndError(ctx, resp, err)
	if err != nil {
		return ctx
	}
	ctx = context.WithValue(ctx, userIDKey{}, resp.Payload.ID)
	fmt.Println("user id = ", ctx.Value(userIDKey{}))
	return ctx

}

func (u *usersFeature) iRegisterANewUserAs(ctx context.Context, username string) context.Context {
	resp, err := u.client.User.RegisterUser(user.NewRegisterUserParams().WithBody(&models.UserspbRegisterUserRequest{
		Username: withRandomString(username),
		Email:    fmt.Sprintf("%s@eample.com.com", withRandomString(username)),
		Password: fmt.Sprintf("%s-password", withRandomString(username)),
	}))
	ctx = setLastResponseAndError(ctx, resp, err)
	if err != nil {
		return ctx
	}
	return context.WithValue(ctx, userIDKey{}, resp.Payload.ID)
}

func (u *usersFeature) expectAUserNamedToExist(username string) error {
	db, _ := u.getDB()
	defer db.Close()

	var userID string
	row := db.QueryRow("SELECT id FROM users WHERE username = $1", withRandomString(username))
	err := row.Scan(&userID)
	if err != nil {
		return errors.ErrNotFound.Msgf("the user `%s` does not exist", username)
	}
	return nil
}

func (u *usersFeature) expectNoUserNamedToExist(username string) error {
	db, _ := u.getDB()
	defer db.Close()

	var userID string
	row := db.QueryRow("SELECT id FROM users WHERE username = $1", withRandomString(username))
	err := row.Scan(&userID)
	if err == sql.ErrNoRows {
		return nil
	} else if err != nil {
		return err
	}
	return errors.ErrAlreadyExists.Msgf("the user `%s`does exist", username)
}

func (u *usersFeature) expectTheUserWasCreated(ctx context.Context) error {
	if err := lastResponseWas(ctx, &user.RegisterUserOK{}); err != nil {
		return err
	}
	return nil
}

func lastUserID(ctx context.Context) (string, error) {
	v := ctx.Value(userIDKey{})
	if v == nil {
		return "", errors.ErrNotFound.Msg("no user ID to work with")
	}
	return v.(string), nil
}
