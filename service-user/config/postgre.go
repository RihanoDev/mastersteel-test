package config

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Config struct {
	ConnectionURI string
	Host          string
	Port          int
	Username      string
	Password      string
	Database      string
	Table         string
	SSLMode       string
	Reset         bool
	GCInterval    time.Duration
}

var ConfigDefault = Config{
	ConnectionURI: "",
	Host:          "postgresql",
	Port:          5432,
	Database:      "api_gateway",
	Username:      "postgres",
	Password:      "123",
	Table:         "service_users",
	SSLMode:       "disable",
	Reset:         false,
	GCInterval:    10 * time.Second,
}

var postgresDB *pgxpool.Pool

func init() {
	var err error

	config := ConfigDefault
	connConfig := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", config.Host, config.Port, config.Username, config.Password, config.Database, config.SSLMode)

	dbpool, err := pgxpool.Connect(context.Background(), connConfig)
	if err != nil {
		panic(err)
	}

	err = dbpool.Ping(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to PostgreSQL!")

	postgresDB = dbpool
}

func NewPostgresContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}

// func GetPostgresPool() *pgxpool.Pool {
// 	return postgresDB
// }

func GetPostgresDatabase() *pgxpool.Pool {
	return postgresDB
}
