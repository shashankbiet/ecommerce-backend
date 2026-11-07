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

var singletonPgCategoryStore *PgCategoryStore
var onceCategory sync.Once
var categoryStoreError error

type PgCategoryStore struct {
	sqlDb *sql.DB
}

func GetPgCategoryStore() (*PgCategoryStore, error) {
	onceCategory.Do(func() {
		sqlDb, err := db.GetSqlConnection()
		if err != nil {
			categoryStoreError = err
			return
		}
		categoryStoreError = initPgCategoryStore(sqlDb)
	})
	return singletonPgCategoryStore, categoryStoreError
}

func initPgCategoryStore(db *sql.DB) error {
	if err := createCategoriesTable(db); err != nil {
		return err
	}
	singletonPgCategoryStore = &PgCategoryStore{
		sqlDb: db,
	}
	return nil
}

func (c PgCategoryStore) Add(category *model.Category) (int16, error) {
	if c.sqlDb == nil {
		logger.Error(DB_CONNECTION_FAILED)
		return 0, errors.New(DB_CONNECTION_FAILED)
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), SQL_TIMEOUT*time.Second)
	defer cancelFunc()

	var id int16
	err := c.sqlDb.QueryRowContext(ctx, CATEGORY_INSERT_QUERY, category.Name, time.Now()).Scan(&id)
	if err != nil {
		return 0, mapConstraintError(err)
	}
	return id, nil
}

func (c PgCategoryStore) GetById(id int16) (*model.Category, error) {
	if c.sqlDb == nil {
		logger.Error(DB_CONNECTION_FAILED)
		return nil, errors.New(DB_CONNECTION_FAILED)
	}

	var category model.Category
	row := c.sqlDb.QueryRow(CATEGORY_GET_BY_ID_QUERY, id)
	err := row.Scan(&category.Id, &category.Name, &category.CreatedAt)
	return &category, err
}

func (c PgCategoryStore) GetAll() (map[string]*model.Category, error) {
	if c.sqlDb == nil {
		logger.Error(DB_CONNECTION_FAILED)
		return nil, errors.New(DB_CONNECTION_FAILED)
	}

	rows, err := c.sqlDb.Query(CATEGORY_GET_QUERY)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	categories := make(map[string]*model.Category, 0)
	for rows.Next() {
		var category model.Category
		err := rows.Scan(&category.Id, &category.Name, &category.CreatedAt)
		if err != nil {
			logger.Error(err.Error())
			continue
		}
		categories[category.Name] = &category
	}
	return categories, rows.Err()
}
