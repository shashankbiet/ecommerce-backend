package producthandler

import (
	"context"
	"search-service/app/service"
	"search-service/pkg/logger"
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
	logger.Info("received GetProduct request", "request", req)
	err := validateGetProductRequest(req)
	if err != nil {
		return nil, err
	}
	result, err := ps.productService.FilterProduct(req)
	if err != nil {
		return nil, err
	}
	return result, nil
}
