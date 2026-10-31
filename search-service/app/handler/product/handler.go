package producthandler

import (
	"context"
	"search-service/app/service"
	searchpb "search-service/proto/search"
)

type productServer struct {
	productService service.IProductService
}

func NewProductServer(service service.IProductService) *productServer {
	return &productServer{
		productService: service,
	}
}

func (ps *productServer) GetProduct(ctx context.Context, req *searchpb.ProductSearchRequest) (*searchpb.ProductSearchResponse, error) {
	err := validateGetProductRequest(req)
	if err != nil {
		return nil, err
	}
	return ps.productService.FilterProduct(req)
}
