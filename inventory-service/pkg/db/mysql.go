package db

import (
	"database/sql"
	"fmt"
	"inventory-service/app/config"
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
	config := config.GetConfig()
	// Build the DSN (Data Source Name) for the MySQL connection
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", config.SqlConfig.Username, config.SqlConfig.Password, config.SqlConfig.Host, config.SqlConfig.Port, config.SqlConfig.Database)

	// Open a connection to the MySQL database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	db.SetConnMaxLifetime(time.Minute * time.Duration(config.SqlConfig.MaxLifetime))
	db.SetMaxOpenConns(config.SqlConfig.MaxOpenConns)
	db.SetMaxIdleConns(config.SqlConfig.MaxIdleConns)

	// Test the connection
	if err := db.Ping(); err != nil {
		return err
	}

	// Set the global DB variable to the opened database connection
	singletonSqlDb = db
	fmt.Println("Connected to the mysql database!")

	return nil
}
