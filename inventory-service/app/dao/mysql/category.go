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

var singletonMySqlCategoryStore *MySqlCategoryStore
var onceCategory sync.Once

type MySqlCategoryStore struct {
	sqlDb *sql.DB
}

func GetMySqlCategoryStore() *MySqlCategoryStore {
	onceCategory.Do(func() {
		sqlDb := db.GetSqlConnection()
		initMySqlCategoryStore(sqlDb)
	})
	return singletonMySqlCategoryStore
}

func initMySqlCategoryStore(db *sql.DB) {
	createCategoriesTable(db)
	singletonMySqlCategoryStore = &MySqlCategoryStore{
		sqlDb: db,
	}
}

func (c MySqlCategoryStore) Add(category *model.Category) (int16, error) {
	if c.sqlDb == nil {
		logger.Log.Error(DB_CONNECTION_FAILED)
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), MYSQL_TIMEOUT*time.Second)
	defer cancelFunc()

	res, err := c.sqlDb.ExecContext(ctx, CATEGORY_INSERT_QUERY, category.Name, time.Now())
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == MYSQL_CONFLICT_ERR_NUMBER {
			return 0, fmt.Errorf(constants.CONFLICT_ERROR_MESSAGE)
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

func (c MySqlCategoryStore) GetById(id int16) (*model.Category, error) {
	if c.sqlDb == nil {
		logger.Log.Error(DB_CONNECTION_FAILED)
	}

	var category model.Category
	row := c.sqlDb.QueryRow(CATEGORY_GET_BY_ID_QUERY, id)
	err := row.Scan(&category.Id, &category.Name, &category.CreatedAt)
	return &category, err
}

func (c MySqlCategoryStore) GetAll() (map[string]*model.Category, error) {
	if c.sqlDb == nil {
		logger.Log.Error(DB_CONNECTION_FAILED)
		return nil, fmt.Errorf(DB_CONNECTION_FAILED)
	}

	rows, err := c.sqlDb.Query(CATEGORY_GET_QUERY)
	if err != nil {
		logger.Log.Error(err.Error())
		return nil, err
	}
	defer rows.Close()

	categories := make(map[string]*model.Category, 0)
	for rows.Next() {
		var category model.Category
		err := rows.Scan(&category.Id, &category.Name, &category.CreatedAt)
		if err != nil {
			logger.Log.Error(err.Error())
		}
		categories[category.Name] = &category
	}
	return categories, nil
}
