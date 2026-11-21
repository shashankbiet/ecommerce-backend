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

var singletonPgKeywordStore *PgKeywordStore
var onceKeyword sync.Once
var keywordStoreError error

type PgKeywordStore struct {
	sqlDb *sql.DB
}

var _ dao.IKeywordDataStore = (*PgKeywordStore)(nil)

func GetPgKeywordStore() (*PgKeywordStore, error) {
	onceKeyword.Do(func() {
		sqlDb, err := db.GetSqlConnection()
		if err != nil {
			keywordStoreError = err
			return
		}
		keywordStoreError = initPgKeywordStore(sqlDb)
	})
	return singletonPgKeywordStore, keywordStoreError
}

func initPgKeywordStore(db *sql.DB) error {
	if err := createKeywordsTable(db); err != nil {
		return err
	}
	singletonPgKeywordStore = &PgKeywordStore{sqlDb: db}
	return nil
}

func (k PgKeywordStore) Add(keyword *model.Keyword) (int64, error) {
	if k.sqlDb == nil {
		logger.Error(DB_CONNECTION_FAILED)
		return 0, errors.New(DB_CONNECTION_FAILED)
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), SQL_TIMEOUT*time.Second)
	defer cancelFunc()

	var id int64
	err := k.sqlDb.QueryRowContext(ctx, KEYWORD_INSERT_QUERY, keyword.Name).Scan(&id)
	if err != nil {
		return 0, mapConstraintError(err)
	}
	return id, nil
}

func (k PgKeywordStore) GetById(id int64) (*model.Keyword, error) {
	if k.sqlDb == nil {
		logger.Error(DB_CONNECTION_FAILED)
		return nil, errors.New(DB_CONNECTION_FAILED)
	}

	var keyword model.Keyword
	err := k.sqlDb.QueryRow(KEYWORD_GET_BY_ID_QUERY, id).Scan(&keyword.Id, &keyword.Name)
	return &keyword, err
}

func (k PgKeywordStore) GetAll() (map[string]*model.Keyword, error) {
	if k.sqlDb == nil {
		logger.Error(DB_CONNECTION_FAILED)
		return nil, errors.New(DB_CONNECTION_FAILED)
	}

	rows, err := k.sqlDb.Query(KEYWORD_GET_QUERY)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	keywords := make(map[string]*model.Keyword)
	for rows.Next() {
		var keyword model.Keyword
		if err := rows.Scan(&keyword.Id, &keyword.Name); err != nil {
			logger.Error(err.Error())
			continue
		}
		keywords[keyword.Name] = &keyword
	}
	return keywords, rows.Err()
}
