package service

import (
	"inventory-service/app/dao"
	model "inventory-service/app/models"
	brandRequest "inventory-service/app/models/request/brand"
)

type BrandService struct {
	datastore dao.IBrandDataStore
}

func NewBrandService(datastore dao.IBrandDataStore) IBrandService {
	return &BrandService{datastore: datastore}
}

func (b *BrandService) GetAll(request *brandRequest.GetAllRequest) ([]*model.Brand, error) {
	return b.datastore.GetAll()
}
