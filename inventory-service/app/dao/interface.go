package dao

import model "inventory-service/app/models"

type ICategoryDataStore interface {
	Add(category *model.Category) (int16, error)
	GetById(id int16) (*model.Category, error)
	GetAll() (map[string]*model.Category, error)
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
