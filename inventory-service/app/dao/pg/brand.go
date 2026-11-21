package pg

import (
	"database/sql"
	"errors"
	"inventory-service/app/dao"
	model "inventory-service/app/models"
	"inventory-service/pkg/db"
	"inventory-service/pkg/logger"
	"sync"
)

var singletonPgBrandStore *PgBrandStore
var onceBrand sync.Once
var brandStoreError error

type PgBrandStore struct {
	sqlDb *sql.DB
}

var _ dao.IBrandDataStore = (*PgBrandStore)(nil)

func GetPgBrandStore() (*PgBrandStore, error) {
	onceBrand.Do(func() {
		sqlDb, err := db.GetSqlConnection()
		if err != nil {
			brandStoreError = err
			return
		}
		brandStoreError = initPgBrandStore(sqlDb)
	})
	return singletonPgBrandStore, brandStoreError
}

func initPgBrandStore(db *sql.DB) error {
	if err := createBrandsTable(db); err != nil {
		return err
	}
	singletonPgBrandStore = &PgBrandStore{sqlDb: db}
	return nil
}

func (b PgBrandStore) GetAll() ([]*model.Brand, error) {
	if b.sqlDb == nil {
		logger.Error(DB_CONNECTION_FAILED)
		return nil, errors.New(DB_CONNECTION_FAILED)
	}

	rows, err := b.sqlDb.Query(BRAND_GET_QUERY)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	brands := make([]*model.Brand, 0)
	for rows.Next() {
		var brand model.Brand
		if err := rows.Scan(&brand.Id, &brand.Name, &brand.CreatedAt); err != nil {
			return nil, err
		}
		brands = append(brands, &brand)
	}
	return brands, rows.Err()
}
