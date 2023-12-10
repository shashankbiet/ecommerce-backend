package datastore

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"inventory-service/app/constants"
	model "inventory-service/app/models"
	"inventory-service/pkg/db"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/go-sql-driver/mysql"
)

const (
	DB_CONNECTION_FAILED = "database connection not initialized"
	INSERT_QUERY         = "INSERT INTO categories(name, created_at) VALUES(?, ?)"
	GET_BY_ID_QUERY      = "SELECT id,name,created_at FROM categories WHERE id=?"
	GET_QUERY            = "SELECT id,name,created_at FROM categories"
	TIMEOUT              = 5
)

var singletonMySqlCategoryStore *MySqlCategoryStore
var once sync.Once

type MySqlCategoryStore struct {
	sqlDb *sql.DB
}

func GetMySqlCategoryStore() *MySqlCategoryStore {
	once.Do(func() {
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

func createCategoriesTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS categories(
		id INT AUTO_INCREMENT PRIMARY KEY,
        name VARCHAR(255) UNIQUE NOT NULL,
		created_at DATETIME
	)`
	ctx, cancelFunc := context.WithTimeout(context.Background(), TIMEOUT*time.Second)
	defer cancelFunc()
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when creating categories table", err)
		return err
	}
	return nil
}

func (c MySqlCategoryStore) Add(category *model.Category) (int64, error) {
	if c.sqlDb == nil {
		log.Fatal(DB_CONNECTION_FAILED)
	}

	ctx, canceFunc := context.WithTimeout(context.Background(), TIMEOUT*time.Second)
	defer canceFunc()

	res, err := c.sqlDb.ExecContext(ctx, INSERT_QUERY, strings.ToUpper(category.Name), time.Now())
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return 0, fmt.Errorf(constants.CONFLICT_ERROR_MESSAGE)
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

func (c MySqlCategoryStore) GetById(id int64) (*model.Category, error) {
	if c.sqlDb == nil {
		log.Fatal(DB_CONNECTION_FAILED)
	}

	var category model.Category
	row := c.sqlDb.QueryRow(GET_BY_ID_QUERY, id)
	err := row.Scan(&category.Id, &category.Name, &category.CreatedAt)
	return &category, err
}

func (c MySqlCategoryStore) GetAll() (map[string]*model.Category, error) {
	if c.sqlDb == nil {
		log.Fatal(DB_CONNECTION_FAILED)
		return nil, fmt.Errorf(DB_CONNECTION_FAILED)
	}

	rows, err := c.sqlDb.Query(GET_QUERY)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	defer rows.Close()

	categories := make(map[string]*model.Category, 0)
	for rows.Next() {
		var (
			id        int
			name      string
			createdAt time.Time
		)

		err := rows.Scan(&id, &name, &createdAt)
		if err != nil {
			log.Fatal(err.Error())
		}
		categories[name] = &model.Category{
			Id:        id,
			Name:      name,
			CreatedAt: createdAt,
		}
	}
	return categories, nil
}
