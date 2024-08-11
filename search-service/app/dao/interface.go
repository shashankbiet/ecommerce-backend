package dao

import (
	searchpb "search-service/proto/search"
)

type IProductDataStore interface {
	GetProductById(id string) (*searchpb.Product, error)
	FilterProduct(req *searchpb.ProductSearchRequest) (*searchpb.ProductSearchResponse, error)
	UpsertProduct(productId int, product *searchpb.Product) error
	DeleteProduct(productId int) error
}
