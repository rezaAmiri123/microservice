///go:build e2e

package e2e

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-openapi/runtime/client"
	"github.com/rezaAmiri123/microservice/stores/storesclient"
	"github.com/rezaAmiri123/microservice/stores/storesclient/models"
	"github.com/rezaAmiri123/microservice/stores/storesclient/product"
	"github.com/rezaAmiri123/microservice/stores/storesclient/store"

	"github.com/cucumber/godog"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/stackus/errors"
)

type storeIDKey struct{}
type productIDKey struct{}
type productMapKey struct{}

type storesFeature struct {
	//db     *sql.DB
	client *storesclient.StoreServiceAPI
}

var _ feature = (*storesFeature)(nil)

func (c *storesFeature) getDB() (*sql.DB, error) {
	return sql.Open("pgx", "postgres://stores_user:stores_pass@localhost:5432/stores?sslmode=disable&search_path=stores,public")
}
func (c *storesFeature) init(cfg featureConfig) (err error) {
	conn := client.New("localhost:8080", "/", nil)
	c.client = storesclient.New(conn, strfmt.Default)
	return err
}

func (c *storesFeature) register(ctx *godog.ScenarioContext) {
	ctx.Step(`^a valid store owner$`, c.noop)
	ctx.Step(`^I create (?:the|a) store called "([^"]*)"$`, c.iCreateTheStoreCalled)
	ctx.Step(`^the store (?:called )?"([^"]*)" already exists$`, c.iCreateTheStoreCalled)
	ctx.Step(`^(?:I )?(?:ensure |expect )?the store (?:was|is) created$`, c.expectTheStoreWasCreated)
	ctx.Step(`^(?:I )?(?:ensure |expect )?a store called "([^"]*)" (?:to )?exists?$`, c.expectAStoreCalledToExist)
	ctx.Step(`^(?:I )?(?:ensure |expect )?no store called "([^"]*)" (?:to )?exists?$`, c.expectNoStoreCalledToExist)
	//
	ctx.Step(`^I create the product called "([^"]*)"$`, c.iCreateTheProductCalled)
	ctx.Step(`^I create the product called "([^"]*)" with price "([^"]*)"$`, c.iCreateTheProductCalledWithPrice)
	ctx.Step(`^(?:I )?(?:ensure |expect )?the product (?:was|is) created$`, c.expectTheProductWasCreated)
	ctx.Step(`^(?:I )?(?:ensure |expect )?a product called "([^"]*)" (?:to )?exists?$`, c.expectAProductCalledToExist)
	ctx.Step(`^(?:I )?(?:ensure |expect )?no product called "([^"]*)" (?:to )?exists?$`, c.expectNoProductCalledToExist)
	//
	ctx.Step(`^a store has the following items$`, c.aStoreHasTheFollowingItems)
	//ctx.Step(`^I add the items$`, c.iAddTheItems)
	//ctx.Step(`^the items are added$`, c.theItemsAreAdded)

	//  Scenario: Cannot create stores without a name
	//    Given a valid store owner
	//    When I create a store called ""
	//    Then I receive a "the store name cannot be blank" error

	//  Scenario: Creating a store called "Waldorf Books"
	//    Given a valid store owner
	//    And no store called "Waldorf Books" exists
	//    When I create the store called "Waldorf Books"
	//    Then the store is created

}

func (c *storesFeature) reset() {
	db, _ := c.getDB()
	defer db.Close()
	truncate := func(tableName string) {
		_, err := db.Exec(fmt.Sprintf("TRUNCATE %s", tableName))
		if err != nil {
			fmt.Errorf("reset error: %v", err)
		}
	}

	truncate("stores")
	truncate("products")
	truncate("events")
	truncate("snapshots")
	truncate("inbox")
	truncate("outbox")

	//   Scenario: Cannot create stores without a name
	//    Given a valid store owner
	//    When I create a store called ""
	//    Then I receive a "the store name cannot be blank" error
}

func (c *storesFeature) noop() {
	// noop
}

func (c *storesFeature) expectAStoreCalledToExist(ctx context.Context, name string) error {
	db, _ := c.getDB()
	defer db.Close()

	var storeID string
	row := db.QueryRow("SELECT id FROM stores.stores WHERE name = $1", withRandomString(name))
	err := row.Scan(&storeID)
	if err != nil {
		return errors.ErrNotFound.Msgf("the store `%s` does not exist", name)
	}

	return nil
}

func (c *storesFeature) expectNoStoreCalledToExist(name string) error {
	db, _ := c.getDB()
	defer db.Close()

	var storeID string
	row := db.QueryRow("SELECT id FROM stores WHERE name = $1", withRandomString(name))
	err := row.Scan(&storeID)
	if err == sql.ErrNoRows {
		return nil
	} else if err != nil {
		return err
	}

	return errors.ErrAlreadyExists.Msgf("the store `%s` does exist", name)
}

func (c *storesFeature) iCreateTheStoreCalled(ctx context.Context, name string) context.Context {
	resp, err := c.client.Store.CreateStore(store.NewCreateStoreParams().WithBody(&models.StorespbCreateStoreRequest{
		Location: "anywhere",
		Name:     withRandomString(name),
	}))

	ctx = setLastResponseAndError(ctx, resp, err)
	if err != nil {
		return ctx
	}
	ctx = context.WithValue(ctx, storeIDKey{}, resp.Payload.ID)
	return ctx
}

func (c *storesFeature) expectTheStoreWasCreated(ctx context.Context) error {
	if err := lastResponseWas(ctx, &store.CreateStoreOK{}); err != nil {
		return err
	}

	return nil
}

func (c *storesFeature) expectAProductCalledToExist(ctx context.Context, name string) error {
	db, _ := c.getDB()
	defer db.Close()

	var productID string
	row := db.QueryRow("SELECT id FROM stores.products WHERE name = $1", withRandomString(name))
	err := row.Scan(&productID)
	if err != nil {
		return errors.ErrNotFound.Msgf("the product `%s` does not exist", name)
	}

	return nil
}

func (c *storesFeature) expectNoProductCalledToExist(name string) error {
	db, _ := c.getDB()
	defer db.Close()

	var productID string
	row := db.QueryRow("SELECT id FROM stores.products WHERE name = $1", withRandomString(name))
	err := row.Scan(&productID)
	if err == sql.ErrNoRows {
		return nil
	} else if err != nil {
		return err
	}

	return errors.ErrAlreadyExists.Msgf("the product `%s` does exist", name)
}

func (c *storesFeature) iCreateTheProductCalled(ctx context.Context, name string) (context.Context, error) {
	return c.iCreateTheProductCalledWithPrice(ctx, name, 9.99)
}

func (c *storesFeature) iCreateTheProductCalledWithPrice(ctx context.Context, name string, price float64) (context.Context, error) {
	storeID, err := lastStoreID(ctx)
	if err != nil {
		return ctx, err
	}
	resp, err := c.client.Product.AddProduct(product.NewAddProductParams().WithBody(&models.StorespbAddProductRequest{
		StoreID: storeID,
		Name:    withRandomString(name),
		Price:   price,
	}))
	ctx = setLastResponseAndError(ctx, resp, err)
	if err != nil {
		return ctx, nil
	}

	ctx = addProduct(ctx, resp.Payload.ID, name)
	ctx = context.WithValue(ctx, productIDKey{}, resp.Payload.ID)
	return ctx, nil
}

func (c *storesFeature) expectTheProductWasCreated(ctx context.Context) error {
	if err := lastResponseWas(ctx, &product.AddProductOK{}); err != nil {
		return err
	}

	return nil
}
func (c *storesFeature) theItemsAreAdded(ctx context.Context) error {
	//if err := lastResponseWas(ctx, &product.AddProductOK{}); err != nil {
	//	return err
	//}

	return nil
}

func (c *storesFeature) aStoreHasTheFollowingItems(ctx context.Context, table *godog.Table) (context.Context, error) {
	type Item struct {
		Name  string
		Price float64
	}
	ctx = c.iCreateTheStoreCalled(ctx, withRandomString("AnyStore"))

	if err := lastError(ctx); err != nil {
		return ctx, err
	}
	items, err := assist.CreateSlice(new(Item), table)
	if err != nil {
		return ctx, err
	}

	for _, i := range items.([]*Item) {
		ctx, err = c.iCreateTheProductCalledWithPrice(ctx, i.Name, i.Price)
		if err != nil {
			return ctx, err
		}
	}

	return ctx, nil
}

func (c *storesFeature) iAddTheItems(ctx context.Context, table *godog.Table) (context.Context, error) {
	//type Item struct {
	//	Name     string
	//	Quantity int
	//}
	//ctx = c.iCreateTheStoreCalled(ctx, "AnyStore")
	//
	//if err := lastError(ctx); err != nil {
	//	return ctx, err
	//}
	//items, err := assist.CreateSlice(new(Item), table)
	//if err != nil {
	//	return ctx, err
	//}
	//
	//for _, i := range items.([]*Item) {
	//	ctx, err = c.iCreateTheProductCalledWithPrice(ctx, i.Name, i.Price)
	//	if err != nil {
	//		return ctx, err
	//	}
	//}

	return ctx, nil
}

func lastStoreID(ctx context.Context) (string, error) {
	v := ctx.Value(storeIDKey{})
	if v == nil {
		return "", errors.ErrNotFound.Msg("no store ID to work with")
	}
	return v.(string), nil
}

func lastProductID(ctx context.Context) (string, error) {
	v := ctx.Value(productIDKey{})
	if v == nil {
		return "", errors.ErrNotFound.Msg("no product ID to work with")
	}
	return v.(string), nil
}

func addProduct(ctx context.Context, id, name string) context.Context {
	var products map[string]string
	v := ctx.Value(productMapKey{})
	if v == nil {
		products = make(map[string]string)
		ctx = context.WithValue(ctx, productMapKey{}, products)
	} else {
		products = v.(map[string]string)
	}

	products[withRandomString(name)] = id
	return ctx
}

func getProductID(ctx context.Context, name string) string {
	v := ctx.Value(productMapKey{})
	if v == nil {
		return uuid.NewString()
	}
	products := v.(map[string]string)

	id, exists := products[withRandomString(name)]
	if !exists {
		return uuid.NewString()
	}

	return id
}
