package business

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // calls an init
)

type Config struct {
	User       string
	Password   string
	Host       string
	Port       int
	Name       string
	DisableTLS bool
}

type DBStore struct {
	DB     *sql.DB
	config Config
}

// NewDatabaseConnection a psql connection
func NewDatabaseConnection(cfg Config) (DBStore, error) {
	sslMode := "require"
	if cfg.DisableTLS {
		sslMode = "disable"
	}

	connURI := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name, sslMode)

	db, err := sql.Open("postgres", connURI)
	if err != nil {
		return DBStore{}, err
	}

	d := DBStore{
		DB:     db,
		config: cfg,
	}

	return d, nil
}
