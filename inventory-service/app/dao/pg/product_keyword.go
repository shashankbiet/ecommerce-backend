package pg

import (
	"context"
	"database/sql"
	"errors"
	"inventory-service/app/dao"
	model "inventory-service/app/models"
	"inventory-service/pkg/db"
	"inventory-service/pkg/logger"
	"sync"
	"time"
)

var singletonPgProductKeywordStore *PgProductKeywordStore
var onceProductKeyword sync.Once
var productKeywordStoreError error

type PgProductKeywordStore struct {
	sqlDb *sql.DB
}

var _ dao.IProductKeywordDataStore = (*PgProductKeywordStore)(nil)

func GetPgProductKeywordStore() (*PgProductKeywordStore, error) {
	onceProductKeyword.Do(func() {
		sqlDb, err := db.GetSqlConnection()
		if err != nil {
			productKeywordStoreError = err
			return
		}
		productKeywordStoreError = initPgProductKeywordStore(sqlDb)
	})
	return singletonPgProductKeywordStore, productKeywordStoreError
}

func initPgProductKeywordStore(db *sql.DB) error {
	if err := createProductKeywordsTable(db); err != nil {
		return err
	}
	singletonPgProductKeywordStore = &PgProductKeywordStore{sqlDb: db}
	return nil
}

func (p PgProductKeywordStore) Add(productKeyword *model.ProductKeyword) (int64, error) {
	if p.sqlDb == nil {
		logger.Error(DB_CONNECTION_FAILED)
		return 0, errors.New(DB_CONNECTION_FAILED)
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), SQL_TIMEOUT*time.Second)
	defer cancelFunc()

	_, err := p.sqlDb.ExecContext(
		ctx,
		PRODUCT_KEYWORD_INSERT_QUERY,
		productKeyword.ProductId,
		productKeyword.KeywordId,
	)
	if err != nil {
		return 0, mapConstraintError(err)
	}
	return productKeyword.ProductId, nil
}

func (p PgProductKeywordStore) GetByProductId(productId int64) ([]*model.ProductKeyword, error) {
	return p.getAll(PRODUCT_KEYWORD_GET_BY_PRODUCT_ID_QUERY, productId)
}

func (p PgProductKeywordStore) GetByKeywordId(keywordId int64) ([]*model.ProductKeyword, error) {
	return p.getAll(PRODUCT_KEYWORD_GET_BY_KEYWORD_ID_QUERY, keywordId)
}

func (p PgProductKeywordStore) getAll(query string, id int64) ([]*model.ProductKeyword, error) {
	if p.sqlDb == nil {
		logger.Error(DB_CONNECTION_FAILED)
		return nil, errors.New(DB_CONNECTION_FAILED)
	}

	rows, err := p.sqlDb.Query(query, id)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	productKeywords := make([]*model.ProductKeyword, 0)
	for rows.Next() {
		var productKeyword model.ProductKeyword
		if err := rows.Scan(&productKeyword.ProductId, &productKeyword.KeywordId); err != nil {
			logger.Error(err.Error())
			continue
		}
		productKeywords = append(productKeywords, &productKeyword)
	}
	return productKeywords, rows.Err()
}

func (p PgProductKeywordStore) DeleteById(productId int64, keywordId int64) error {
	if p.sqlDb == nil {
		logger.Error(DB_CONNECTION_FAILED)
		return errors.New(DB_CONNECTION_FAILED)
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), SQL_TIMEOUT*time.Second)
	defer cancelFunc()

	_, err := p.sqlDb.ExecContext(ctx, PRODUCT_KEYWORD_DELETE_QUERY, productId, keywordId)
	if err != nil {
		return mapConstraintError(err)
	}
	return nil
}
