package pg

import (
	"context"
	"database/sql"
	"fmt"
	model "inventory-service/app/models"
	"inventory-service/pkg/db"
	"inventory-service/pkg/logger"
	"sync"
	"time"
)

var singletonPgSubCategoryStore *PgSubCategoryStore
var onceSubcategory sync.Once
var subcategoryStoreError error

type PgSubCategoryStore struct {
	sqlDb *sql.DB
}

func GetPgSubCategoryStore() (*PgSubCategoryStore, error) {
	onceSubcategory.Do(func() {
		sqlDb, err := db.GetSqlConnection()
		if err != nil {
			subcategoryStoreError = err
			return
		}
		subcategoryStoreError = initPgSubCategoryStore(sqlDb)
	})
	return singletonPgSubCategoryStore, subcategoryStoreError
}

func initPgSubCategoryStore(db *sql.DB) error {
	if err := createSubCategoriesTable(db); err != nil {
		return err
	}
	singletonPgSubCategoryStore = &PgSubCategoryStore{
		sqlDb: db,
	}
	return nil
}

func (c PgSubCategoryStore) Add(subcategory *model.SubCategory) (int16, error) {
	if c.sqlDb == nil {
		logger.Error(DB_CONNECTION_FAILED)
		return 0, fmt.Errorf(DB_CONNECTION_FAILED)
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), SQL_TIMEOUT*time.Second)
	defer cancelFunc()

	var id int16
	err := c.sqlDb.QueryRowContext(ctx, SUB_CATEGORY_INSERT_QUERY, subcategory.Name, subcategory.Category, time.Now()).Scan(&id)
	if err != nil {
		return 0, mapConstraintError(err)
	}
	return id, nil
}

func (c PgSubCategoryStore) GetById(id int16) (*model.SubCategory, error) {
	if c.sqlDb == nil {
		logger.Error(DB_CONNECTION_FAILED)
		return nil, fmt.Errorf(DB_CONNECTION_FAILED)
	}

	var subcategory model.SubCategory
	row := c.sqlDb.QueryRow(SUB_CATEGORY_GET_BY_ID_QUERY, id)
	err := row.Scan(&subcategory.Id, &subcategory.Name, &subcategory.Category, &subcategory.CreatedAt)
	return &subcategory, err
}

func (c PgSubCategoryStore) GetAll() (map[string]*model.SubCategory, error) {
	if c.sqlDb == nil {
		logger.Error(DB_CONNECTION_FAILED)
		return nil, fmt.Errorf(DB_CONNECTION_FAILED)
	}

	rows, err := c.sqlDb.Query(SUB_CATEGORY_GET_QUERY)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	defer rows.Close()

	subcategories := make(map[string]*model.SubCategory, 0)
	for rows.Next() {
		var subcategory model.SubCategory
		err := rows.Scan(&subcategory.Id, &subcategory.Name, &subcategory.Category, &subcategory.CreatedAt)
		if err != nil {
			logger.Error(err.Error())
			continue
		}
		subcategories[subcategory.Name] = &subcategory
	}
	return subcategories, rows.Err()
}
