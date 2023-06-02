package e2e

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/cucumber/godog"
	"github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/rezaAmiri123/microservice/baskets/basketsclient"
	"github.com/rezaAmiri123/microservice/baskets/basketsclient/basket"
	"github.com/rezaAmiri123/microservice/baskets/basketsclient/item"
	"github.com/rezaAmiri123/microservice/baskets/basketsclient/models"
	"github.com/stackus/errors"
)

type basketIDKey = struct{}

type basketsFeature struct {
	client *basketsclient.BasketServiceAPI
	//db     *sql.DB
}

var _ feature = (*basketsFeature)(nil)

func (c *basketsFeature) getDB() (*sql.DB, error) {
	return sql.Open("pgx", "postgres://baskets_user:baskets_pass@localhost:5432/baskets?sslmode=disable&search_path=baskets,public")
}

func (c *basketsFeature) init(cfg featureConfig) (err error) {
	//if cfg.useMonoDB {
	//	c.db, err = sql.Open("pgx", "postgres://mallbots_user:mallbots_pass@localhost:5432/mallbots?sslmode=disable")
	//} else {
	//	c.db, err = sql.Open("pgx", "postgres://baskets_user:baskets_pass@localhost:5432/baskets?sslmode=disable&search_path=baskets,public")
	//}
	//if err != nil {
	//	return
	//}
	conn := client.New("localhost:8080", "/", nil)
	c.client = basketsclient.New(conn, strfmt.Default)

	return
}

func (c *basketsFeature) register(ctx *godog.ScenarioContext) {
	ctx.Step(`^I start a new basket$`, c.iStartANewBasket)
	ctx.Step(`^(?:I )?(?:ensure |expect )?the basket (?:was|is) started$`, c.expectTheBasketWasStarted)

	ctx.Step(`^I add the items$`, c.iAddTheItems)
	ctx.Step(`^(?:I )?(?:ensure |expect )?the items (?:were|are) added$`, c.expectTheItemsWereAdded)
}

func (c *basketsFeature) reset() {
	db, _ := c.getDB()
	defer db.Close()

	truncate := func(tableName string) {
		_, _ = db.Exec(fmt.Sprintf("TRUNCATE %s", tableName))
	}

	truncate("events")
	truncate("snapshots")
	truncate("inbox")
	truncate("outbox")
	truncate("products_cache")
	truncate("stores_cache")
}

func (c *basketsFeature) iStartANewBasket(ctx context.Context) (context.Context, error) {
	userID, err := lastUserID(ctx)
	fmt.Println("basket: user id: ", userID)
	if err != nil {
		return ctx, err
	}
	resp, err := c.client.Basket.StartBasket(basket.NewStartBasketParams().WithBody(&models.BasketspbStartBasketRequest{
		UserID: userID,
	}))
	ctx = setLastResponseAndError(ctx, resp, err)
	if err != nil {
		return ctx, nil
	}
	ctx = context.WithValue(ctx, basketIDKey{}, resp.Payload.ID)
	fmt.Println("basket id = ", ctx.Value(basketIDKey{}))
	return ctx, nil

}

func (c *basketsFeature) expectTheBasketWasStarted(ctx context.Context) error {
	if err := lastResponseWas(ctx, &basket.StartBasketOK{}); err != nil {
		return err
	}
	return nil
}

func (c *basketsFeature) iAddTheItems(ctx context.Context, table *godog.Table) (context.Context, error) {
	type Item struct {
		Name     string
		Quantity int
	}
	fmt.Println("iAddTheItems Tables")
	basketID, err := lastBasketID(ctx)
	fmt.Println("basket id: ", basketID)
	if err != nil {
		return ctx, err
	}
	items, err := assist.CreateSlice(new(Item), table)
	if err != nil {
		return ctx, err
	}
	for _, i := range items.([]*Item) {
		productID := getProductID(ctx, i.Name)
		fmt.Println("product id: ", productID)
		resp, err := c.client.Item.AddItem(item.NewAddItemParams().WithBody(&models.BasketspbAddItemRequest{
			ID:        basketID,
			ProductID: productID,
			Quantity:  int32(i.Quantity),
		}))
		if err != nil {
			fmt.Println("got error: ", err.Error())
		}
		ctx = setLastResponseAndError(ctx, resp, err)
		if err != nil {
			break
		}
	}

	return ctx, nil
}

func (c *basketsFeature) expectTheItemsWereAdded(ctx context.Context) error {
	if err := lastResponseWas(ctx, &item.AddItemOK{}); err != nil {
		return err
	}
	return nil
}

func lastBasketID(ctx context.Context) (string, error) {
	v := ctx.Value(basketIDKey{})
	if v == nil {
		return "", errors.ErrNotFound.Msg("no basket ID to work with")
	}
	return v.(string), nil
}
