package dao

import (
	searchpb "search-service/proto/search"
)

type IProductDataStore interface {
	GetProductById(id string) (*searchpb.Product, error)
	FilterProduct(req *searchpb.ProductSearchRequest) (*searchpb.ProductSearchResponse, error)
	AddProduct(product *searchpb.Product) error
	UpdateProduct(productId int, product *searchpb.Product) error
	DeleteProduct(productId int) error
}
