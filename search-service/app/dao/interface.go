package dao

import (
	model "search-service/app/models"
	searchpb "search-service/proto/search"
)

type IProductDataStore interface {
	GetProductById(id string) (*searchpb.Product, error)
	FilterProduct(req *searchpb.ProductSearchRequest) (*searchpb.ProductSearchResponse, error)
	UpsertProduct(productId int64, product *model.Product) error
	DeleteProduct(productId int64) error
}
