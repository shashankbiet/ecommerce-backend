package db

import (
	"database/sql"
	"fmt"
	"inventory-service/pkg/config"
	"inventory-service/pkg/logger"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// DB is a global variable representing the database connection
var singletonSqlDb *sql.DB
var once sync.Once

// Initialize the database connection
func GetSqlConnection() *sql.DB {
	once.Do(func() {
		initSqlConnection()
	})
	return singletonSqlDb
}

func initSqlConnection() error {
	// Build the DSN (Data Source Name) for the MySQL connection
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", config.GetConfig().SqlConfig.Username, config.GetConfig().SqlConfig.Password, config.GetConfig().SqlConfig.Host, config.GetConfig().SqlConfig.Port, config.GetConfig().SqlConfig.Database)

	// Open a connection to the MySQL database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	db.SetConnMaxLifetime(time.Minute * time.Duration(config.GetConfig().SqlConfig.MaxLifetime))
	db.SetMaxOpenConns(config.GetConfig().SqlConfig.MaxOpenConns)
	db.SetMaxIdleConns(config.GetConfig().SqlConfig.MaxIdleConns)

	// Test the connection
	if err := db.Ping(); err != nil {
		return err
	}

	// Set the global DB variable to the opened database connection
	singletonSqlDb = db
	logger.Info("connected to the mysql database!")

	return nil
}
