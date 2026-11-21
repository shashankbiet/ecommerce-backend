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
	return createTable(db, "sub categories", SUB_CATEGORY_CREATE_TABLE_QUERY)
}

func createBrandsTable(db *sql.DB) error {
	return createTable(db, "brands", BRAND_CREATE_TABLE_QUERY)
}

func createProductsTable(db *sql.DB) error {
	if err := createBrandsTable(db); err != nil {
		return err
	}
	if err := createTable(db, "products", PRODUCT_CREATE_TABLE_QUERY); err != nil {
		return err
	}
	if err := createTable(db, "products brand constraint", PRODUCT_BRAND_CONSTRAINT_QUERY); err != nil {
		return err
	}
	if err := createTable(db, "products status constraint", PRODUCT_STATUS_CONSTRAINT_QUERY); err != nil {
		return err
	}
	if err := createKeywordsTable(db); err != nil {
		return err
	}
	return createProductKeywordsTable(db)
}

func createKeywordsTable(db *sql.DB) error {
	return createTable(db, "keywords", KEYWORD_CREATE_TABLE_QUERY)
}

func createProductKeywordsTable(db *sql.DB) error {
	return createTable(db, "product keywords", PRODUCT_KEYWORD_CREATE_TABLE_QUERY)
}

func createInventoryTable(db *sql.DB) error {
	return createTable(db, "inventory", INVENTORY_CREATE_TABLE_QUERY)
}
