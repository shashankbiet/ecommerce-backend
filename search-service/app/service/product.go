package service

import (
	"search-service/app/dao"
	searchpb "search-service/proto/search"
)

type ProductService struct {
	datastore dao.IProductDataStore
}

func NewProductService(datastore dao.IProductDataStore) *ProductService {
	return &ProductService{datastore: datastore}
}

func (s *ProductService) FilterProduct(req *searchpb.ProductSearchRequest) (*searchpb.ProductSearchResponse, error) {
	return s.datastore.FilterProduct(req)
}
