///go:build e2e

package e2e

import (
	"context"
	"flag"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/cucumber/godog"
	"github.com/go-openapi/runtime/client"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/rdumont/assistdog"
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
}

func TestEndToEnd2(t *testing.T) {
	opts := godog.Options{
		Format:    "progress",
		Paths:     []string{"features"},
		Randomize: time.Now().UTC().UnixNano(), // randomize scenario execution order
	}
	fmt.Println(opts)

}
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

	cfg := featureConfig{
		transport: client.New("localhost:8080", "/", nil),
		useMonoDB: *useMonoDB,
	}

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
	)
	if err != nil {
		t.Fatal(err)
	}

	featurePaths := []string{
		"features/baskets",
		"features/users",
		//"features/kiosk",
		//"features/orders",
		"features/stores",
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
