package pg

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"inventory-service/app/constants"
	"inventory-service/pkg/logger"
	"time"

	"github.com/lib/pq"
)

const (
	SQL_TIMEOUT                    = 5
	POSTGRES_UNIQUE_VIOLATION      = "23505"
	POSTGRES_FOREIGN_KEY_VIOLATION = "23503"
	DB_CONNECTION_FAILED           = "database connection not initialized"
)

func mapConstraintError(err error) error {
	var pqErr *pq.Error
	if errors.As(err, &pqErr) {
		switch pqErr.Code {
		case POSTGRES_UNIQUE_VIOLATION:
			return fmt.Errorf(constants.CONFLICT_ERROR_MESSAGE)
		case POSTGRES_FOREIGN_KEY_VIOLATION:
			return fmt.Errorf(constants.UNPROCESSABLE_ENTITY_ERROR_MESSAGE)
		}
	}
	return err
}

func createTable(db *sql.DB, name, query string) error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), SQL_TIMEOUT*time.Second)
	defer cancelFunc()
	if _, err := db.ExecContext(ctx, query); err != nil {
		logger.Error("error in creating "+name+" table", "error", err)
		return err
	}
	return nil
}

func createCategoriesTable(db *sql.DB) error {
	return createTable(db, "categories", CATEGORY_CREATE_TABLE_QUERY)
}

func createSubCategoriesTable(db *sql.DB) error {
	return createTable(db, "subcategories", SUB_CATEGORY_CREATE_TABLE_QUERY)
}

func createProductsTable(db *sql.DB) error {
	return createTable(db, "products", PRODUCT_CREATE_TABLE_QUERY)
}

func createInventoryTable(db *sql.DB) error {
	if err := createTable(db, "inventory", INVENTORY_CREATE_TABLE_QUERY); err != nil {
		return err
	}
	return createTable(db, "inventory quantity", INVENTORY_ADD_QUANTITY_COLUMN_QUERY)
}
