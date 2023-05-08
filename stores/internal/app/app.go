package app

import (
	"context"
	"github.com/rezaAmiri123/microservice/pkg/ddd"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/rezaAmiri123/microservice/stores/internal/app/commands"
	"github.com/rezaAmiri123/microservice/stores/internal/app/queries"
	"github.com/rezaAmiri123/microservice/stores/internal/domain"
)

//type Application struct {
//	Commands Commands
//	Queries  Queries
//}
//
//type Queries struct {
//	GetProduct *queries.GetProductHandler
//	GetStore   *queries.GetStoreHandler
//}
//
//type Commands struct {
//	AddProduct  *commands.AddProductHandler
//	CreateStore *commands.CreateStoreHandler
//}

type (
	App interface {
		Commands
		Queries
	}
	Commands interface {
		CreateStore(ctx context.Context, cmd commands.CreateStore) error
		EnableParticipation(ctx context.Context, cmd commands.EnableParticipation) error
		DisableParticipation(ctx context.Context, cmd commands.DisableParticipation) error
		RebrandStore(ctx context.Context, cmd commands.RebrandStore) error
		AddProduct(ctx context.Context, cmd commands.AddProduct) error
		RebrandProduct(ctx context.Context, cmd commands.RebrandProduct) error
		IncreaseProductPrice(ctx context.Context, cmd commands.IncreaseProductPrice) error
		DecreaseProductPrice(ctx context.Context, cmd commands.DecreaseProductPrice) error
		RemoveProduct(ctx context.Context, cmd commands.RemoveProduct) error
	}
	Queries interface {
		GetStore(ctx context.Context, query queries.GetStore) (*domain.MallStore, error)
		GetStores(ctx context.Context, query queries.GetStores) ([]*domain.MallStore, error)
		GetParticipatingStores(ctx context.Context, query queries.GetParticipatingStores) ([]*domain.MallStore, error)
		GetCatalog(ctx context.Context, query queries.GetCatalog) ([]*domain.CatalogProduct, error)
		GetProduct(ctx context.Context, query queries.GetProduct) (*domain.CatalogProduct, error)
	}

	Application struct {
		appCommands
		appQueries
	}

	appCommands struct {
		commands.CreateStoreHandler
		commands.EnableParticipationHandler
		commands.DisableParticipationHandler
		commands.RebrandStoreHandler
		commands.AddProductHandler
		commands.RebrandProductHandler
		commands.IncreaseProductPriceHandler
		commands.DecreaseProductPriceHandler
		commands.RemoveProductHandler
	}
	appQueries struct {
		queries.GetStoreHandler
		queries.GetStoresHandler
		queries.GetParticipatingStoresHandler
		queries.GetProductHandler
		queries.GetCatalogHandler
	}
)

var _ App = (*Application)(nil)

func New(
	stores domain.StoreRepository,
	products domain.ProductRepository,
	catalog domain.CatalogRepository,
	mall domain.MallRepository,
	publisher ddd.EventPublisher[ddd.Event],
	logger logger.Logger,
) Application {
	return Application{
		appCommands: appCommands{
			CreateStoreHandler:          commands.NewCreateStoreHandler(stores, publisher, logger),
			EnableParticipationHandler:  commands.NewEnableParticipationHandler(stores, publisher, logger),
			DisableParticipationHandler: commands.NewDisableParticipationHandler(stores, publisher, logger),
			RebrandStoreHandler:         commands.NewRebrandStoreHandler(stores, publisher, logger),
			AddProductHandler:           commands.NewAddProductHandler(products, publisher, logger),
			RebrandProductHandler:       commands.NewRebrandProductHandler(products, publisher, logger),
			IncreaseProductPriceHandler: commands.NewIncreaseProductPriceHandler(products, publisher, logger),
			DecreaseProductPriceHandler: commands.NewDecreaseProductPriceHandler(products, publisher, logger),
			RemoveProductHandler:        commands.NewRemoveProductHandler(products, publisher, logger),
		},
		appQueries: appQueries{
			GetStoreHandler:               queries.NewGetStoreHandler(mall, logger),
			GetStoresHandler:              queries.NewGetStoresHandler(mall, logger),
			GetParticipatingStoresHandler: queries.NewGetParticipatingStoresHandler(mall, logger),
			GetProductHandler:             queries.NewGetProductHandler(catalog, logger),
			GetCatalogHandler:             queries.NewGetCatalogHandler(catalog, logger),
		},
	}
}
