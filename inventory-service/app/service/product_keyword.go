package service

import (
	"inventory-service/app/dao"
	model "inventory-service/app/models"
	productKeywordRequest "inventory-service/app/models/request/product_keyword"
)

type ProductKeywordService struct {
	datastore dao.IProductKeywordDataStore
}

func NewProductKeywordService(datastore dao.IProductKeywordDataStore) IProductKeywordService {
	return &ProductKeywordService{datastore: datastore}
}

func (p *ProductKeywordService) Add(request *productKeywordRequest.AddRequest) (int64, error) {
	return p.datastore.Add(request.ProductKeyword)
}

func (p *ProductKeywordService) GetByProductId(request *productKeywordRequest.GetByProductIdRequest) ([]*model.ProductKeyword, error) {
	return p.datastore.GetByProductId(request.ProductId)
}

func (p *ProductKeywordService) GetByKeywordId(request *productKeywordRequest.GetByKeywordIdRequest) ([]*model.ProductKeyword, error) {
	return p.datastore.GetByKeywordId(request.KeywordId)
}

func (p *ProductKeywordService) Delete(request *productKeywordRequest.DeleteRequest) error {
	return p.datastore.DeleteById(request.ProductId, request.KeywordId)
}
