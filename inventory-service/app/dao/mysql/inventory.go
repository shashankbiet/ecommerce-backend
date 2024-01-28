package mysql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"inventory-service/app/constants"
	model "inventory-service/app/models"
	"inventory-service/pkg/db"
	"inventory-service/pkg/logger"
	"sync"
	"time"

	"github.com/go-sql-driver/mysql"
)

var singletonMySqlInventoryStore *MySqlInventoryStore
var onceInventory sync.Once

type MySqlInventoryStore struct {
	sqlDb *sql.DB
}

func GetMySqlInventoryStore() *MySqlInventoryStore {
	onceInventory.Do(func() {
		sqlDb := db.GetSqlConnection()
		initMySqlInventoryStore(sqlDb)
	})
	return singletonMySqlInventoryStore
}

func initMySqlInventoryStore(db *sql.DB) {
	createInventoryTable(db)
	singletonMySqlInventoryStore = &MySqlInventoryStore{
		sqlDb: db,
	}
}

func (c MySqlInventoryStore) Add(inventory *model.Inventory) (int64, error) {
	if c.sqlDb == nil {
		logger.Error(DB_CONNECTION_FAILED)
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), MYSQL_TIMEOUT*time.Second)
	defer cancelFunc()

	res, err := c.sqlDb.ExecContext(ctx, INVENTORY_INSERT_QUERY, inventory.ProductId, inventory.SKU, inventory.PurchasePrice, inventory.SalePrice, time.Now(), time.Now())
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == MYSQL_CONFLICT_ERR_NUMBER {
			return 0, fmt.Errorf(constants.CONFLICT_ERROR_MESSAGE)
		} else if errors.As(err, &mysqlErr) && mysqlErr.Number == MYSQL_FOREIGN_KEY_ERR_NUMBER {
			return 0, fmt.Errorf(constants.UNPROCESSABLE_ENTITY_ERROR_MESSAGE)
		} else {
			return 0, err
		}
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (c MySqlInventoryStore) Update(inventory *model.Inventory) (bool, error) {
	if c.sqlDb == nil {
		logger.Error(DB_CONNECTION_FAILED)
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), MYSQL_TIMEOUT*time.Second)
	defer cancelFunc()

	res, err := c.sqlDb.ExecContext(ctx, INVENTORY_UPDATE_QUERY, inventory.SKU, inventory.PurchasePrice, inventory.SalePrice, time.Now(), inventory.ProductId)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == MYSQL_CONFLICT_ERR_NUMBER {
			return false, fmt.Errorf(constants.CONFLICT_ERROR_MESSAGE)
		} else if errors.As(err, &mysqlErr) && mysqlErr.Number == MYSQL_FOREIGN_KEY_ERR_NUMBER {
			return false, fmt.Errorf(constants.UNPROCESSABLE_ENTITY_ERROR_MESSAGE)
		} else {
			return false, err
		}
	}
	_, err = res.LastInsertId()
	if err != nil {
		return false, err
	}
	return true, nil
}

func (c MySqlInventoryStore) GetByProductId(productId int64) (*model.Inventory, error) {
	if c.sqlDb == nil {
		logger.Error(DB_CONNECTION_FAILED)
	}

	var inventory model.Inventory
	row := c.sqlDb.QueryRow(INVENTORY_GET_BY_ID_QUERY, productId)
	err := row.Scan(&inventory.ProductId, &inventory.SKU, &inventory.PurchasePrice, &inventory.SalePrice, &inventory.CreatedAt, &inventory.UpdatedAt)
	return &inventory, err
}

func (c MySqlInventoryStore) GetAll() ([]*model.Inventory, error) {
	if c.sqlDb == nil {
		logger.Error(DB_CONNECTION_FAILED)
		return nil, fmt.Errorf(DB_CONNECTION_FAILED)
	}

	rows, err := c.sqlDb.Query(INVENTORY_GET_QUERY)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	defer rows.Close()

	inventoryList := make([]*model.Inventory, 0)
	for rows.Next() {
		var inventory model.Inventory
		err := rows.Scan(&inventory.ProductId, &inventory.SKU, &inventory.PurchasePrice, &inventory.SalePrice, &inventory.CreatedAt, &inventory.UpdatedAt)
		if err != nil {
			logger.Error(err.Error())
		}
		inventoryList = append(inventoryList, &inventory)
	}
	return inventoryList, nil
}
