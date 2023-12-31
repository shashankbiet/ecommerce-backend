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

var singletonMySqlProductStore *MySqlProductStore
var onceProduct sync.Once

type MySqlProductStore struct {
	sqlDb *sql.DB
}

func GetMySqlProductStore() *MySqlProductStore {
	onceProduct.Do(func() {
		sqlDb := db.GetSqlConnection()
		initMySqlProductStore(sqlDb)
	})
	return singletonMySqlProductStore
}

func initMySqlProductStore(db *sql.DB) {
	createProductsTable(db)
	singletonMySqlProductStore = &MySqlProductStore{
		sqlDb: db,
	}
}

func (c MySqlProductStore) Add(product *model.Product) (int64, error) {
	if c.sqlDb == nil {
		logger.Log.Error(DB_CONNECTION_FAILED)
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), MYSQL_TIMEOUT*time.Second)
	defer cancelFunc()

	res, err := c.sqlDb.ExecContext(ctx, PRODUCT_INSERT_QUERY, product.Name, product.Description, product.Brand, product.Category, product.SubCategory, product.ImageId, product.Weight, time.Now(), time.Now())
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

func (c MySqlProductStore) Update(product *model.Product) (bool, error) {
	if c.sqlDb == nil {
		logger.Log.Error(DB_CONNECTION_FAILED)
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), MYSQL_TIMEOUT*time.Second)
	defer cancelFunc()

	res, err := c.sqlDb.ExecContext(ctx, PRODUCT_UPDATE_QUERY, product.Name, product.Description, product.Brand, product.Category, product.SubCategory, product.ImageId, product.Weight, time.Now(), product.Id)
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

func (c MySqlProductStore) GetById(id int64) (*model.Product, error) {
	if c.sqlDb == nil {
		logger.Log.Error(DB_CONNECTION_FAILED)
	}

	var product model.Product
	row := c.sqlDb.QueryRow(PRODUCT_GET_BY_ID_QUERY, id)
	err := row.Scan(&product.Id, &product.Name, &product.Description, &product.Brand, &product.Category, &product.SubCategory, &product.ImageId, &product.Weight, &product.CreatedAt, &product.UpdatedAt)
	return &product, err
}

func (c MySqlProductStore) GetAll() ([]*model.Product, error) {
	if c.sqlDb == nil {
		logger.Log.Error(DB_CONNECTION_FAILED)
		return nil, fmt.Errorf(DB_CONNECTION_FAILED)
	}

	rows, err := c.sqlDb.Query(PRODUCT_GET_QUERY)
	if err != nil {
		logger.Log.Error(err.Error())
		return nil, err
	}
	defer rows.Close()

	products := make([]*model.Product, 0)
	for rows.Next() {
		var product model.Product
		err := rows.Scan(&product.Id, &product.Name, &product.Description, &product.Brand, &product.Category, &product.SubCategory, &product.ImageId, &product.Weight, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			logger.Log.Error(err.Error())
		}
		products = append(products, &product)
	}
	return products, nil
}
