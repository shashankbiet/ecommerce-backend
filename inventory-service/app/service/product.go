package service

import (
	"encoding/json"
	"fmt"
	"inventory-service/app/dao"
	model "inventory-service/app/models"
	productRequest "inventory-service/app/models/request/product"
	"inventory-service/app/producer"
	"inventory-service/pkg/logger"
)

type ProductService struct {
	datastore dao.IProductDataStore
	producer  producer.IProductProducer
}

func NewProductService(datastore dao.IProductDataStore, producer producer.IProductProducer) IProductService {
	return &ProductService{
		datastore: datastore,
		producer:  producer,
	}
}

func (p *ProductService) Add(request *productRequest.AddRequest) (int64, error) {
	setCategory(&request.Product.Category)
	setSubCategory(&request.Product.SubCategory)
	id, err := p.datastore.Add(request.Product)
	if err != nil {
		return 0, err
	}

	p.publishProductMessage(request.Product)
	return id, nil
}

func (p *ProductService) Update(request *productRequest.UpdateRequest) (bool, error) {
	_, err := p.datastore.Update(request.Product)
	if err != nil {
		return false, err
	}
	p.publishProductMessage(request.Product)
	return true, nil
}

func (p *ProductService) GetById(request *productRequest.GetByIdRequest) (*model.Product, error) {
	product, err := p.datastore.GetById(request.Id)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return product, nil
}

func (p *ProductService) GetAll(request *productRequest.GetAllRequest) ([]*model.Product, error) {
	products, err := p.datastore.GetAll()
	if err != nil {
		return nil, fmt.Errorf("")
	}
	return products, nil
}

func (p *ProductService) publishProductMessage(product *model.Product) {
	arr, err := json.Marshal(product)
	if err != nil {
		logger.Log.Error("unable to marshal product struct into []byte")
	}
	p.producer.Publish(arr)
}
