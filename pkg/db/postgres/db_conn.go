package postgres

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/jackc/pgx/stdlib" // pgx driver
	"github.com/jmoiron/sqlx"
)

const (
	maxOpenConns    = 60
	connMaxLifetime = 120
	maxIdleConns    = 30
	connMaxIdleTime = 20
)

type Config struct {
	PGDriver     string `mapstructure:"POSTGRES_DRIVER"`
	PGHost       string `mapstructure:"POSTGRES_HOST"`
	PGPort       string `mapstructure:"POSTGRES_PORT"`
	PGUser       string `mapstructure:"POSTGRES_USER"`
	PGDBName     string `mapstructure:"POSTGRES_DB_NAME"`
	PGPassword   string `mapstructure:"POSTGRES_PASSWORD"`
	PGSearchPath string `mapstructure:"POSTGRES_SEARCH_PATH"`
}

// Return new Postgresql db instance
func NewPsqlDB(c Config) (*sqlx.DB, error) {
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		c.PGHost,
		c.PGPort,
		c.PGUser,
		c.PGDBName,
		c.PGPassword,
	)

	db, err := sqlx.Connect(c.PGDriver, dataSourceName)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenConns)
	db.SetConnMaxLifetime(connMaxLifetime * time.Second)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxIdleTime(connMaxIdleTime * time.Second)
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func NewDB(c Config) (*sql.DB, error) {
	// host=postgres dbname=stores user=stores_user password=stores_pass search_path=stores,public
	//dataSourceName := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s search_path=stores,public",
	dataSourceName := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s search_path=%s",
		c.PGHost,
		c.PGPort,
		c.PGDBName,
		c.PGUser,
		c.PGPassword,
		c.PGSearchPath,
	)

	db, err := sql.Open(c.PGDriver, dataSourceName)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenConns)
	db.SetConnMaxLifetime(connMaxLifetime * time.Second)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxIdleTime(connMaxIdleTime * time.Second)
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
