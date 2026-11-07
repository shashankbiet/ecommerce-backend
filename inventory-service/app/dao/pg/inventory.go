package pg

import (
	"context"
	"database/sql"
	"errors"
	model "inventory-service/app/models"
	"inventory-service/pkg/db"
	"inventory-service/pkg/logger"
	"sync"
	"time"
)

var singletonPgInventoryStore *PgInventoryStore
var onceInventory sync.Once
var inventoryStoreError error

type PgInventoryStore struct {
	sqlDb *sql.DB
}

func GetPgInventoryStore() (*PgInventoryStore, error) {
	onceInventory.Do(func() {
		sqlDb, err := db.GetSqlConnection()
		if err != nil {
			inventoryStoreError = err
			return
		}
		inventoryStoreError = initPgInventoryStore(sqlDb)
	})
	return singletonPgInventoryStore, inventoryStoreError
}

func initPgInventoryStore(db *sql.DB) error {
	if err := createInventoryTable(db); err != nil {
		return err
	}
	singletonPgInventoryStore = &PgInventoryStore{
		sqlDb: db,
	}
	return nil
}

func (c PgInventoryStore) Add(inventory *model.Inventory) (int64, error) {
	if c.sqlDb == nil {
		logger.Error(DB_CONNECTION_FAILED)
		return 0, errors.New(DB_CONNECTION_FAILED)
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), SQL_TIMEOUT*time.Second)
	defer cancelFunc()

	_, err := c.sqlDb.ExecContext(ctx, INVENTORY_INSERT_QUERY, inventory.ProductId, inventory.SKU, inventory.PurchasePrice, inventory.SalePrice, inventory.Quantity, time.Now(), time.Now())
	if err != nil {
		return 0, mapConstraintError(err)
	}
	return inventory.ProductId, nil
}

func (c PgInventoryStore) Update(inventory *model.Inventory) (bool, error) {
	if c.sqlDb == nil {
		logger.Error(DB_CONNECTION_FAILED)
		return false, errors.New(DB_CONNECTION_FAILED)
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), SQL_TIMEOUT*time.Second)
	defer cancelFunc()

	res, err := c.sqlDb.ExecContext(ctx, INVENTORY_UPDATE_QUERY, inventory.SKU, inventory.PurchasePrice, inventory.SalePrice, inventory.Quantity, time.Now(), inventory.ProductId)
	if err != nil {
		return false, mapConstraintError(err)
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return false, err
	}
	if rowsAffected == 0 {
		return false, nil
	}
	return true, nil
}

func (c PgInventoryStore) GetByProductId(productId int64) (*model.Inventory, error) {
	if c.sqlDb == nil {
		logger.Error(DB_CONNECTION_FAILED)
		return nil, errors.New(DB_CONNECTION_FAILED)
	}

	var inventory model.Inventory
	row := c.sqlDb.QueryRow(INVENTORY_GET_BY_ID_QUERY, productId)
	err := row.Scan(&inventory.ProductId, &inventory.SKU, &inventory.PurchasePrice, &inventory.SalePrice, &inventory.Quantity, &inventory.CreatedAt, &inventory.UpdatedAt)
	return &inventory, err
}

func (c PgInventoryStore) GetAll() ([]*model.Inventory, error) {
	if c.sqlDb == nil {
		logger.Error(DB_CONNECTION_FAILED)
		return nil, errors.New(DB_CONNECTION_FAILED)
	}

	rows, err := c.sqlDb.Query(INVENTORY_GET_QUERY)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	inventoryList := make([]*model.Inventory, 0)
	for rows.Next() {
		var inventory model.Inventory
		err := rows.Scan(&inventory.ProductId, &inventory.SKU, &inventory.PurchasePrice, &inventory.SalePrice, &inventory.Quantity, &inventory.CreatedAt, &inventory.UpdatedAt)
		if err != nil {
			logger.Error(err.Error())
			continue
		}
		inventoryList = append(inventoryList, &inventory)
	}
	return inventoryList, rows.Err()
}
