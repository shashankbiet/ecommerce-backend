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

var singletonMySqlSubCategoryStore *MySqlSubCategoryStore
var onceSubcategory sync.Once

type MySqlSubCategoryStore struct {
	sqlDb *sql.DB
}

func GetMySqlSubCategoryStore() *MySqlSubCategoryStore {
	onceSubcategory.Do(func() {
		sqlDb := db.GetSqlConnection()
		initMySqlSubCategoryStore(sqlDb)
	})
	return singletonMySqlSubCategoryStore
}

func initMySqlSubCategoryStore(db *sql.DB) {
	createSubCategoriesTable(db)
	singletonMySqlSubCategoryStore = &MySqlSubCategoryStore{
		sqlDb: db,
	}
}

func (c MySqlSubCategoryStore) Add(subcategory *model.SubCategory) (int16, error) {
	if c.sqlDb == nil {
		logger.Log.Error(DB_CONNECTION_FAILED)
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), MYSQL_TIMEOUT*time.Second)
	defer cancelFunc()

	res, err := c.sqlDb.ExecContext(ctx, SUB_CATEGORY_INSERT_QUERY, subcategory.Name, subcategory.Category, time.Now())
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
	return int16(id), nil
}

func (c MySqlSubCategoryStore) GetById(id int16) (*model.SubCategory, error) {
	if c.sqlDb == nil {
		logger.Log.Error(DB_CONNECTION_FAILED)
	}

	var subcategory model.SubCategory
	row := c.sqlDb.QueryRow(SUB_CATEGORY_GET_BY_ID_QUERY, id)
	err := row.Scan(&subcategory.Id, &subcategory.Name, &subcategory.Category, &subcategory.CreatedAt)
	return &subcategory, err
}

func (c MySqlSubCategoryStore) GetAll() (map[string]*model.SubCategory, error) {
	if c.sqlDb == nil {
		logger.Log.Error(DB_CONNECTION_FAILED)
		return nil, fmt.Errorf(DB_CONNECTION_FAILED)
	}

	rows, err := c.sqlDb.Query(SUB_CATEGORY_GET_QUERY)
	if err != nil {
		logger.Log.Error(err.Error())
		return nil, err
	}
	defer rows.Close()

	subcategories := make(map[string]*model.SubCategory, 0)
	for rows.Next() {
		var subcategory model.SubCategory
		err := rows.Scan(&subcategory.Id, &subcategory.Name, &subcategory.Category, &subcategory.CreatedAt)
		if err != nil {
			logger.Log.Error(err.Error())
		}
		subcategories[subcategory.Name] = &subcategory
	}
	return subcategories, nil
}
