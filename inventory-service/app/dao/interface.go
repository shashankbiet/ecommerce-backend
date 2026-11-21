package dao

import model "inventory-service/app/models"

type ICategoryDataStore interface {
	Add(category *model.Category) (int16, error)
	GetById(id int16) (*model.Category, error)
	GetAll() (map[string]*model.Category, error)
}

type IBrandDataStore interface {
	GetAll() ([]*model.Brand, error)
}

type ISubCategoryDataStore interface {
	Add(subcategory *model.SubCategory) (int16, error)
	GetById(id int16) (*model.SubCategory, error)
	GetAll() (map[string]*model.SubCategory, error)
}

type IProductDataStore interface {
	Add(product *model.Product) (int64, error)
	Update(product *model.Product) (bool, error)
	GetById(id int64) (*model.Product, error)
	GetAll() ([]*model.Product, error)
}

type IInventoryDataStore interface {
	Add(inventory *model.Inventory) (int64, error)
	Update(inventory *model.Inventory) (bool, error)
	GetByProductId(productId int64) (*model.Inventory, error)
	GetAll() ([]*model.Inventory, error)
}

type IKeywordDataStore interface {
	Add(keyword *model.Keyword) (int64, error)
	GetById(id int64) (*model.Keyword, error)
	GetAll() (map[string]*model.Keyword, error)
}

type IProductKeywordDataStore interface {
	Add(productKeyword *model.ProductKeyword) (int64, error)
	GetByProductId(productId int64) ([]*model.ProductKeyword, error)
	GetByKeywordId(keywordId int64) ([]*model.ProductKeyword, error)
	DeleteById(productId int64, keywordId int64) error
}
