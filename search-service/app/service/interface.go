package service

import searchpb "search-service/proto/search"

type IProductService interface {
	FilterProduct(req *searchpb.ProductSearchRequest) (*searchpb.ProductSearchResponse, error)
}
