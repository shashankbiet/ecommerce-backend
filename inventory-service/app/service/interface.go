package service

import (
	model "inventory-service/app/models"
	categoryRequest "inventory-service/app/models/request/category"
	inventoryRequest "inventory-service/app/models/request/inventory"
	productRequest "inventory-service/app/models/request/product"
	subCategoryRequest "inventory-service/app/models/request/subcategory"
)

type ICategoryService interface {
	Add(request *categoryRequest.AddRequest) (int16, error)
	GetById(request *categoryRequest.GetByIdRequest) (*model.Category, error)
	GetAll(request *categoryRequest.GetAllRequest) (map[string]*model.Category, error)
}

type ISubCategoryService interface {
	Add(request *subCategoryRequest.AddRequest) (int16, error)
	GetById(request *subCategoryRequest.GetByIdRequest) (*model.SubCategory, error)
	GetAll(request *subCategoryRequest.GetAllRequest) (map[string]*model.SubCategory, error)
}

type IProductService interface {
	Add(request *productRequest.AddRequest) (int64, error)
	Update(request *productRequest.UpdateRequest) (bool, error)
	GetById(request *productRequest.GetByIdRequest) (*model.Product, error)
	GetAll(request *productRequest.GetAllRequest) ([]*model.Product, error)
}

type IInventoryService interface {
	Add(request *inventoryRequest.AddRequest) (int64, error)
	Update(request *inventoryRequest.UpdateRequest) (bool, error)
	GetByProductId(request *inventoryRequest.GetByIdRequest) (*model.Inventory, error)
	GetAll(request *inventoryRequest.GetAllRequest) ([]*model.Inventory, error)
}
