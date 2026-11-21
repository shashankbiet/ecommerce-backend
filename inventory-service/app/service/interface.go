package service

import (
	model "inventory-service/app/models"
	brandRequest "inventory-service/app/models/request/brand"
	categoryRequest "inventory-service/app/models/request/category"
	inventoryRequest "inventory-service/app/models/request/inventory"
	keywordRequest "inventory-service/app/models/request/keyword"
	productRequest "inventory-service/app/models/request/product"
	productKeywordRequest "inventory-service/app/models/request/product_keyword"
	subCategoryRequest "inventory-service/app/models/request/subcategory"
)

type IBrandService interface {
	GetAll(request *brandRequest.GetAllRequest) ([]*model.Brand, error)
}

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

type IKeywordService interface {
	Add(request *keywordRequest.AddRequest) (int64, error)
	GetById(request *keywordRequest.GetByIdRequest) (*model.Keyword, error)
	GetAll(request *keywordRequest.GetAllRequest) (map[string]*model.Keyword, error)
}

type IProductKeywordService interface {
	Add(request *productKeywordRequest.AddRequest) (int64, error)
	GetByProductId(request *productKeywordRequest.GetByProductIdRequest) ([]*model.ProductKeyword, error)
	GetByKeywordId(request *productKeywordRequest.GetByKeywordIdRequest) ([]*model.ProductKeyword, error)
	Delete(request *productKeywordRequest.DeleteRequest) error
}
