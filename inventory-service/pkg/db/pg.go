package db

import (
	"context"
	"database/sql"
	"fmt"
	"inventory-service/pkg/config"
	"inventory-service/pkg/logger"
	"net"
	"net/url"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

var singletonSqlDb *sql.DB
var once sync.Once
var connectionError error

func GetSqlConnection() (*sql.DB, error) {
	once.Do(func() {
		singletonSqlDb, connectionError = initSqlConnection()
	})
	return singletonSqlDb, connectionError
}

func initSqlConnection() (*sql.DB, error) {
	sqlConfig := config.GetConfig().SqlConfig
	dsn := postgresDSN(sqlConfig)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("open PostgreSQL connection: %w", err)
	}

	db.SetConnMaxLifetime(time.Minute * time.Duration(sqlConfig.MaxLifetime))
	db.SetMaxOpenConns(sqlConfig.MaxOpenConns)
	db.SetMaxIdleConns(sqlConfig.MaxIdleConns)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		db.Close()
		return nil, fmt.Errorf("ping PostgreSQL: %w", err)
	}

	logger.Info("connected to the postgres database!")

	return db, nil
}

func postgresDSN(sqlConfig config.SqlConfig) string {
	return (&url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(sqlConfig.Username, sqlConfig.Password),
		Host:   net.JoinHostPort(sqlConfig.Host, sqlConfig.Port),
		Path:   sqlConfig.Database,
		RawQuery: url.Values{
			"sslmode": {"disable"},
		}.Encode(),
	}).String()
}
