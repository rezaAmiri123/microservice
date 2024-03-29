// /go:build e2e
package e2e

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/cucumber/godog"
	"github.com/go-openapi/strfmt"
	"github.com/rezaAmiri123/microservice/payments/paymentsclient"
	"github.com/rezaAmiri123/microservice/payments/paymentsclient/models"
	"github.com/rezaAmiri123/microservice/payments/paymentsclient/payment"
	"github.com/rezaAmiri123/microservice/pkg/db/postgres"
	"github.com/stackus/errors"
)

type paymentsIDKey struct{}

type paymentsFeature struct {
	client *paymentsclient.PaymentServiceAPI
	db     *sql.DB
}

var _ feature = (*paymentsFeature)(nil)

func (c *paymentsFeature) init(cfg featureConfig) (err error) {
	c.client = paymentsclient.New(cfg.transport, strfmt.Default)
	c.db, err = postgres.NewDB(postgres.Config{
		PGDriver:     cfg.PGDriver,
		PGHost:       cfg.PGHost,
		PGPort:       cfg.PGPort,
		PGUser:       cfg.PGPaymentsUser,
		PGDBName:     cfg.PGPaymentsDBName,
		PGPassword:   cfg.PGPaymentsPassword,
		PGSearchPath: cfg.PGPaymentsSearchPath,
	})
	if err != nil {
		return fmt.Errorf("cannot load db: %w", err)
	}

	return
}

func (c *paymentsFeature) register(ctx *godog.ScenarioContext) {
	ctx.Step(`^I authorize the payment$`, c.iCreateThePayment)
	ctx.Step(`^(?:I )?(?:ensure |expect )?the payment (?:was|is) authorized$`, c.expectThePaymentWasAuthorized)
}

func (c *paymentsFeature) reset() {

}

func (c *paymentsFeature) iCreateThePayment(ctx context.Context) context.Context {
	return c.iCreateThePaymentWithAmount(ctx, 50.0)
}
func (c *paymentsFeature) iCreateThePaymentWithAmount(ctx context.Context, amount float64) context.Context {
	userID, err := lastUserID(ctx)
	if err != nil {
		return ctx
	}
	resp, err := c.client.Payment.AuthorizePayment(payment.NewAuthorizePaymentParams().WithBody(&models.PaymentspbAuthorizePaymentRequest{
		UserID: userID,
		Amount: amount,
	}))
	ctx = setLastResponseAndError(ctx, resp, err)
	if err != nil {
		return ctx
	}
	return context.WithValue(ctx, paymentsIDKey{}, resp.Payload.ID)
}

func (c *paymentsFeature) expectThePaymentWasAuthorized(ctx context.Context) error {
	if err := lastResponseWas(ctx, &payment.AuthorizePaymentOK{}); err != nil {
		return err
	}
	return nil
}

func lastPaymentID(ctx context.Context) (string, error) {
	v := ctx.Value(paymentsIDKey{})
	if v == nil {
		return "", errors.ErrNotFound.Msg("no payment Id to work with")
	}
	return v.(string), nil
}
