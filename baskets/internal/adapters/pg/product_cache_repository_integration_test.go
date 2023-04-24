///go:build integration || database

package pg

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/docker/go-connections/nat"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/rezaAmiri123/microservice/baskets/internal/constants"
	"github.com/rezaAmiri123/microservice/baskets/internal/domain"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"path/filepath"
	"testing"
	"time"
)

type productCacheSuite struct {
	container testcontainers.Container
	db        *sql.DB
	mock      *domain.MockProductRepository
	repo      ProductCacheRepository
	suite.Suite
}

func TestProductCacheRepository(t *testing.T) {
	if testing.Short() {
		t.Skip("short mode: skipping")
	}
	suite.Run(t, &productCacheSuite{})
}

func (s *productCacheSuite) SetupSuite() {
	var err error

	ctx := context.Background()
	initDir, err := filepath.Abs("./../../../../docker/database")
	if err != nil {
		s.T().Fatal(err)
	}
	const dbUrl = "postgres://mallbots_user:mallbots_pass@%s:%s/mallbots?sslmode=disable"
	s.container, err = testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image:        "postgres:12-alpine",
			Hostname:     "postgres",
			ExposedPorts: []string{"5432/tcp"},
			Env: map[string]string{
				"POSTGRES_PASSWORD": "itsasecret",
			},
			Mounts: []testcontainers.ContainerMount{
				testcontainers.BindMount(initDir, "/docker-entrypoint-initdb.d"),
			},
			WaitingFor: wait.ForSQL("5432/tcp", "pgx", func(host string, port nat.Port) string {
				return fmt.Sprintf(dbUrl, host, port.Port())
			}).WithStartupTimeout(5 * time.Second),
		},
		Started: true,
	})
	if err != nil {
		s.T().Fatal(err)
	}

	endpoint, err := s.container.Endpoint(ctx, "")
	if err != nil {
		s.T().Fatal(err)
	}

	dataSource := fmt.Sprintf("postgres://mallbots_user:mallbots_pass@%s/mallbots?sslmode=disable", endpoint)
	s.db, err = sql.Open("pgx", dataSource)
	if err != nil {
		s.T().Fatal(err)
	}
	driver, err := postgres.WithInstance(s.db, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://../migrations",
		"postgres", driver)
	if err != nil {
		s.T().Fatal(err)
	}
	err = m.Up() // or m.Step(2) if you want to explicitly set the number of migrations to run
	if err != nil && err != migrate.ErrNoChange {
		s.T().Fatal(err)
	}
}
func (s *productCacheSuite) TearDownSuite() {
	err := s.db.Close()
	if err != nil {
		s.T().Fatal(err)
	}
	if err := s.container.Terminate(context.Background()); err != nil {
		s.T().Fatal(err)
	}
}

func (s *productCacheSuite) SetupTest() {
	s.mock = domain.NewMockProductRepository(s.T())
	s.repo = NewProductCacheRepository(constants.ProductsCacheTableName, s.db, s.mock)
}

func (s *productCacheSuite) TearDownTest() {
	_, err := s.db.ExecContext(context.Background(), fmt.Sprintf("TRUNCATE %s", constants.ProductsCacheTableName))
	if err != nil {
		s.T().Fatal(err)
	}
}

func (s *productCacheSuite) TestProductCacheRepository_Add() {
	s.NoError(s.repo.Add(context.Background(), "product-id", "store-id", "product-name", 10.00))
	query := fmt.Sprintf("SELECT name FROM %s WHERE id = $1", constants.ProductsCacheTableName)
	row := s.db.QueryRow(query, "product-id")
	if s.NoError(row.Err()) {
		var name string
		s.NoError(row.Scan(&name))
		s.Equal("product-name", name)
	}
}
func (s *productCacheSuite) TestProductCacheRepository_AddDupe() {
	s.NoError(s.repo.Add(context.Background(), "product-id", "store-id", "product-name", 10.00))
	s.NoError(s.repo.Add(context.Background(), "product-id", "store-id", "dupe-product-name", 10.00))
	query := fmt.Sprintf("SELECT name FROM %s WHERE id = $1", constants.ProductsCacheTableName)
	row := s.db.QueryRow(query, "product-id")
	if s.NoError(row.Err()) {
		var name string
		s.NoError(row.Scan(&name))
		s.Equal("product-name", name)
	}
}
func (s *productCacheSuite) TestProductCacheRepository_Rebrand() {
	// Arrange
	query := fmt.Sprintf("INSERT INTO %s (id, store_id, name, price) VALUES ('product-id', 'store-id', 'product-name', 10.00)", constants.ProductsCacheTableName)
	_, err := s.db.Exec(query)
	s.NoError(err)

	// Act
	s.NoError(s.repo.Rebrand(context.Background(), "product-id", "new-product-name"))

	// Assert
	query = fmt.Sprintf("SELECT name FROM %s WHERE id = $1", constants.ProductsCacheTableName)
	row := s.db.QueryRow(query, "product-id")
	if s.NoError(row.Err()) {
		var name string
		s.NoError(row.Scan(&name))
		s.Equal("new-product-name", name)
	}
}

func (s *productCacheSuite) TestProductCacheRepository_UpdatePrice() {
	query := fmt.Sprintf("INSERT INTO %s (id, store_id, name, price) VALUES ('product-id', 'store-id', 'product-name', 10.00)", constants.ProductsCacheTableName)
	_, err := s.db.Exec(query)
	s.NoError(err)

	s.NoError(s.repo.UpdatePrice(context.Background(), "product-id", 2.00))
	query = fmt.Sprintf("SELECT price FROM %s WHERE id = $1", constants.ProductsCacheTableName)
	row := s.db.QueryRow(query, "product-id")
	if s.NoError(row.Err()) {
		var price float64
		s.NoError(row.Scan(&price))
		s.Equal(12.00, price)
	}
}

func (s *productCacheSuite) TestProductCacheRepository_Remove() {
	query := fmt.Sprintf("INSERT INTO %s (id, store_id, name, price) VALUES ('product-id', 'store-id', 'product-name', 10.00)", constants.ProductsCacheTableName)
	_, err := s.db.Exec(query)
	s.NoError(err)

	s.NoError(s.repo.Remove(context.Background(), "product-id"))
	query = fmt.Sprintf("SELECT price FROM %s WHERE id = $1", constants.ProductsCacheTableName)
	row := s.db.QueryRow(query, "product-id")
	if s.NoError(row.Err()) {
		var name string
		s.Error(row.Scan(&name))
	}
}

func (s *productCacheSuite) TestProductCacheRepository_Find() {
	query := fmt.Sprintf("INSERT INTO %s (id, store_id, name, price) VALUES ('product-id', 'store-id', 'product-name', 10.00)", constants.ProductsCacheTableName)
	_, err := s.db.Exec(query)
	s.NoError(err)

	product, err := s.repo.Find(context.Background(), "product-id")
	if s.NoError(err) {
		s.NotNil(product)
		s.Equal("product-name", product.Name)
	}
}

func (s *productCacheSuite) TestProductCacheRepository_FindFromFallback() {
	s.mock.On("Find", mock.Anything, "product-id").Return(&domain.Product{
		ID:      "product-id",
		StoreID: "store-id",
		Name:    "product-name",
		Price:   10.00,
	}, nil)

	product, err := s.repo.Find(context.Background(), "product-id")
	if s.NoError(err) {
		s.NotNil(product)
		s.Equal("product-name", product.Name)
	}
}

//func runDBMigration(migrationURL string, dbSource string) {
//	driver, err := iofs.New(migrations.FS, "../migrations")
//	if err != nil {
//		log.Fatal(err)
//	}
//	m, err := migrate.NewWithSourceInstance("iofs", driver, "postgres://postgres@localhost/postgres?sslmode=disable")
//	if err != nil {
//		log.Fatal(err)
//	}
//	err = m.Up()
//	if err != nil && err != migrate.ErrNoChange {
//		// ...
//	}
//
//	//migration, err := migrate.New(migrationURL, dbSource)
//	//if err != nil {
//	//	//log.Fatal().Err(err).Msg("cannot create new migrate instance")
//	//}
//	//
//	//if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
//	//	//log.Fatal().Err(err).Msg("failed to run migrate up")
//	//}
//
//	//log.Info().Msg("db migrated successfully")
//}
