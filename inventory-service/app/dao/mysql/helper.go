package mysql

import (
	"context"
	"database/sql"
	"inventory-service/pkg/logger"
	"time"
)

const (
	MYSQL_TIMEOUT                = 5
	MYSQL_CONFLICT_ERR_NUMBER    = 1062
	MYSQL_FOREIGN_KEY_ERR_NUMBER = 1452
	DB_CONNECTION_FAILED         = "database connection not initialized"
)

func createCategoriesTable(db *sql.DB) error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), MYSQL_TIMEOUT*time.Second)
	defer cancelFunc()
	_, err := db.ExecContext(ctx, CATEGORY_CREATE_TABLE_QUERY)
	if err != nil {
		logger.Log.Error("error in creating categories table", "error", err)
		return err
	}
	return nil
}

func createSubCategoriesTable(db *sql.DB) error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), MYSQL_TIMEOUT*time.Second)
	defer cancelFunc()
	_, err := db.ExecContext(ctx, SUB_CATEGORY_CREATE_TABLE_QUERY)
	if err != nil {
		logger.Log.Error("error in creating subcategories table", "error", err)
		return err
	}
	return nil
}

func createProductsTable(db *sql.DB) error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), MYSQL_TIMEOUT*time.Second)
	defer cancelFunc()
	_, err := db.ExecContext(ctx, PRODUCT_CREATE_TABLE_QUERY)
	if err != nil {
		logger.Log.Error("error in creating products table", "error", err)
		return err
	}
	return nil
}

func createInventoryTable(db *sql.DB) error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), MYSQL_TIMEOUT*time.Second)
	defer cancelFunc()
	_, err := db.ExecContext(ctx, INVENTORY_CREATE_TABLE_QUERY)
	if err != nil {
		logger.Log.Error("error in creating inventory table", "error", err)
		return err
	}
	return nil
}
