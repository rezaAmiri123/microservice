///go:build e2e

package e2e

import (
	"context"
	"flag"
	"fmt"
	"github.com/cucumber/godog"
	"github.com/go-openapi/runtime/client"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"time"

	//_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/rdumont/assistdog"
	"github.com/spf13/viper"
	"github.com/stackus/errors"
)

var useMonoDB = flag.Bool("mono", false, "Use mono DB resources")

var assist = assistdog.NewDefault()
var randomString string

type lastResponseKey struct{}
type lastErrorKey struct{}

func withRandomString(str string) string {
	if str == "" {
		return ""
	}
	return str + randomString
}

type feature interface {
	init(cfg featureConfig) error
	register(ctx *godog.ScenarioContext)
	reset()
}

type featureConfig struct {
	transport    *client.Runtime
	useMonoDB    bool
	randomString string

	TransportURL string `mapstructure:"TRANSPORT_URL"`
	PGDriver     string `mapstructure:"POSTGRES_DRIVER"`
	PGHost       string `mapstructure:"POSTGRES_HOST"`
	PGPort       string `mapstructure:"POSTGRES_PORT"`

	PGUsersUser       string `mapstructure:"POSTGRES_USERS_USER"`
	PGUsersDBName     string `mapstructure:"POSTGRES_USERS_DB_NAME"`
	PGUsersPassword   string `mapstructure:"POSTGRES_USERS_PASSWORD"`
	PGUsersSearchPath string `mapstructure:"POSTGRES_USERS_SEARCH_PATH"`

	PGStoresUser       string `mapstructure:"POSTGRES_STORES_USER"`
	PGStoresDBName     string `mapstructure:"POSTGRES_STORES_DB_NAME"`
	PGStoresPassword   string `mapstructure:"POSTGRES_STORES_PASSWORD"`
	PGStoresSearchPath string `mapstructure:"POSTGRES_STORES_SEARCH_PATH"`

	PGBasketsUser       string `mapstructure:"POSTGRES_BASKETS_USER"`
	PGBasketsDBName     string `mapstructure:"POSTGRES_BASKETS_DB_NAME"`
	PGBasketsPassword   string `mapstructure:"POSTGRES_BASKETS_PASSWORD"`
	PGBasketsSearchPath string `mapstructure:"POSTGRES_BASKETS_SEARCH_PATH"`

	PGPaymentsUser       string `mapstructure:"POSTGRES_PAYMENTS_USER"`
	PGPaymentsDBName     string `mapstructure:"POSTGRES_PAYMENTS_DB_NAME"`
	PGPaymentsPassword   string `mapstructure:"POSTGRES_PAYMENTS_PASSWORD"`
	PGPaymentsSearchPath string `mapstructure:"POSTGRES_PAYMENTS_SEARCH_PATH"`
}

//func TestEndToEnd2(t *testing.T) {
//	opts := godog.Options{
//		Format:    "progress",
//		Paths:     []string{"features"},
//		Randomize: time.Now().UTC().UnixNano(), // randomize scenario execution order
//	}
//	fmt.Println(opts)
//
//}

func TestEndToEnd(t *testing.T) {
	assist.RegisterComparer(float64(0.0), func(raw string, actual interface{}) error {
		af, ok := actual.(float64)
		if !ok {
			return fmt.Errorf("%v is not a float64", actual)
		}
		ef, err := strconv.ParseFloat(raw, 64)
		if err != nil {
			return err
		}

		if ef != af {
			return fmt.Errorf("expected %v, but got %v", ef, af)
		}

		return nil
	})
	assist.RegisterParser(float64(0.0), func(raw string) (interface{}, error) {
		return strconv.ParseFloat(raw, 64)
	})

	cfg, err := LoadConfig(".")
	if err != nil {
		t.Error(err)
	}
	cfg.transport = client.New(cfg.TransportURL, "/", nil)
	cfg.useMonoDB = *useMonoDB
	//cfg := featureConfig{
	//	transport: client.New("localhost:8080", "/", nil),
	//	useMonoDB: *useMonoDB,
	//}

	features, err := func(fs ...feature) ([]feature, error) {
		features := make([]feature, len(fs))
		for i, f := range fs {
			err := f.init(cfg)
			if err != nil {
				return features, err
			}
			features[i] = f
		}
		return features, nil
	}(
		&basketsFeature{},
		&usersFeature{},
		&storesFeature{},
		&paymentsFeature{},
	)
	if err != nil {
		t.Fatal(err)
	}

	featurePaths := []string{
		"features/baskets",
		"features/users",
		"features/kiosk",
		"features/orders",
		"features/stores",
		"features/payments",
	}

	suite := godog.TestSuite{
		Name: "mallbots-e2e",
		ScenarioInitializer: func(ctx *godog.ScenarioContext) {
			ctx.Step(`^I receive a "([^"]*)" error$`, iReceiveAError)
			ctx.Step(`^(?:ensure |expect )?the returned error message is "([^"]*)"$`, iReceiveAError)
			for _, f := range features {
				f.register(ctx)
			}
			ctx.Before(func(ctx context.Context, s *godog.Scenario) (context.Context, error) {
				//for _, f := range features {
				//f.reset()
				//}
				randomString = fmt.Sprintf("%d", time.Now().UTC().UnixNano())

				return ctx, nil
			})

		},
		Options: &godog.Options{
			Format:    "pretty",
			Paths:     featurePaths,
			Randomize: time.Now().UTC().UnixNano(), // randomize scenario execution order,
		},
	}

	if status := suite.Run(); status != 0 {
		t.Error("end to end feature test failed with status:", status)
	}
	//suite.TestSuiteInitializer
}

func iReceiveAError(ctx context.Context, msg string) error {
	err := lastError(ctx)
	if err == nil {
		return errors.Wrap(errors.ErrUnknown, "expected error to not be nil")
	}
	fmt.Println("got error: ", err.Error())
	fmt.Println("want error: ", "Message:"+msg)
	if !strings.Contains(err.Error(), msg) {
		fmt.Println("return error")
		return errors.Wrapf(errors.ErrInvalidArgument, "expected: %s: got: %s", msg, err.Error())
	}
	fmt.Println("return nil")
	return nil
}

func setLastResponseAndError(ctx context.Context, resp any, err error) context.Context {
	if err != nil {
		fmt.Println("setLastResponseAndError: ", err.Error())
	}
	return context.WithValue(
		context.WithValue(ctx, lastResponseKey{}, resp),
		lastErrorKey{}, err,
	)
}

func lastResponseWas(ctx context.Context, resp any) error {
	r := ctx.Value(lastResponseKey{})
	if reflect.ValueOf(r).Kind() == reflect.Ptr && reflect.ValueOf(r).IsNil() {
		e := ctx.Value(lastErrorKey{})
		if e == nil {
			return errors.ErrUnknown.Msg("no last response or error")
		}
		return e.(error)
	}
	if reflect.TypeOf(r) == reflect.TypeOf(resp) {
		return nil
	}
	return errors.ErrBadRequest.Msgf("last request was `%v`", r)
}

func lastResponse(ctx context.Context) any {
	return ctx.Value(lastResponseKey{})
}

func lastError(ctx context.Context) error {
	e := ctx.Value(lastErrorKey{})
	if e == nil {
		return nil
	}
	return e.(error)
}

func LoadConfig(path string) (config featureConfig, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
