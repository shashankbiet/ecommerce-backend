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

var singletonPgProductStore *PgProductStore
var onceProduct sync.Once
var productStoreError error

type PgProductStore struct {
	sqlDb *sql.DB
}

func GetPgProductStore() (*PgProductStore, error) {
	onceProduct.Do(func() {
		sqlDb, err := db.GetSqlConnection()
		if err != nil {
			productStoreError = err
			return
		}
		productStoreError = initPgProductStore(sqlDb)
	})
	return singletonPgProductStore, productStoreError
}

func initPgProductStore(db *sql.DB) error {
	if err := createProductsTable(db); err != nil {
		return err
	}
	singletonPgProductStore = &PgProductStore{
		sqlDb: db,
	}
	return nil
}

func (c PgProductStore) Add(product *model.Product) (int64, error) {
	if c.sqlDb == nil {
		logger.Error(DB_CONNECTION_FAILED)
		return 0, errors.New(DB_CONNECTION_FAILED)
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), SQL_TIMEOUT*time.Second)
	defer cancelFunc()

	var deletedAt interface{}
	if !product.DeletedAt.IsZero() {
		deletedAt = product.DeletedAt
	}

	var id int64
	now := time.Now().UTC()
	err := c.sqlDb.QueryRowContext(
		ctx,
		PRODUCT_INSERT_QUERY,
		product.Name,
		product.Description,
		product.Slug,
		product.Brand,
		product.Category,
		product.SubCategory,
		product.ImageUrl,
		product.Status,
		now,
		now,
		deletedAt,
	).Scan(&id)
	if err != nil {
		return 0, mapConstraintError(err)
	}
	product.Id = id
	product.CreatedAt = now
	product.UpdatedAt = now
	return id, nil
}

func (c PgProductStore) Update(product *model.Product) (bool, error) {
	if c.sqlDb == nil {
		logger.Error(DB_CONNECTION_FAILED)
		return false, errors.New(DB_CONNECTION_FAILED)
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), SQL_TIMEOUT*time.Second)
	defer cancelFunc()

	var deletedAt interface{}
	if !product.DeletedAt.IsZero() {
		deletedAt = product.DeletedAt
	}

	updatedAt := time.Now().UTC()
	res, err := c.sqlDb.ExecContext(
		ctx,
		PRODUCT_UPDATE_QUERY,
		product.Name,
		product.Description,
		product.Slug,
		product.Brand,
		product.Category,
		product.SubCategory,
		product.ImageUrl,
		product.Status,
		updatedAt,
		deletedAt,
		product.Id,
	)
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
	product.UpdatedAt = updatedAt
	return true, nil
}

func (c PgProductStore) GetById(id int64) (*model.Product, error) {
	if c.sqlDb == nil {
		logger.Error(DB_CONNECTION_FAILED)
		return nil, errors.New(DB_CONNECTION_FAILED)
	}

	var product model.Product
	var deletedAt sql.NullTime
	row := c.sqlDb.QueryRow(PRODUCT_GET_BY_ID_QUERY, id)
	err := row.Scan(
		&product.Id,
		&product.Name,
		&product.Description,
		&product.Slug,
		&product.Brand,
		&product.Category,
		&product.SubCategory,
		&product.ImageUrl,
		&product.Status,
		&product.CreatedAt,
		&product.UpdatedAt,
		&deletedAt,
	)
	if deletedAt.Valid {
		product.DeletedAt = deletedAt.Time
	}
	return &product, err
}

func (c PgProductStore) GetAll() ([]*model.Product, error) {
	if c.sqlDb == nil {
		logger.Error(DB_CONNECTION_FAILED)
		return nil, errors.New(DB_CONNECTION_FAILED)
	}

	rows, err := c.sqlDb.Query(PRODUCT_GET_QUERY)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	products := make([]*model.Product, 0)
	for rows.Next() {
		var product model.Product
		var deletedAt sql.NullTime
		err := rows.Scan(
			&product.Id,
			&product.Name,
			&product.Description,
			&product.Slug,
			&product.Brand,
			&product.Category,
			&product.SubCategory,
			&product.ImageUrl,
			&product.Status,
			&product.CreatedAt,
			&product.UpdatedAt,
			&deletedAt,
		)
		if err != nil {
			logger.Error(err.Error())
			continue
		}
		if deletedAt.Valid {
			product.DeletedAt = deletedAt.Time
		}
		products = append(products, &product)
	}
	return products, rows.Err()
}
