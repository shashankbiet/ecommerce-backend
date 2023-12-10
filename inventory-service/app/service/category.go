package service

import (
	"fmt"
	intf "inventory-service/app/interface"
	model "inventory-service/app/models"
	categoryRequest "inventory-service/app/models/request/category"
)

type CategoryService struct {
	datastore intf.CategoryDataStore
}

func NewCategoryService(datastore intf.CategoryDataStore) intf.CategoryService {
	return &CategoryService{
		datastore: datastore,
	}
}

func (c CategoryService) Add(request *categoryRequest.AddRequest) (int64, error) {
	id, err := c.datastore.Add(request.Category)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (c CategoryService) GetById(request *categoryRequest.GetByIdRequest) (*model.Category, error) {
	category, err := c.datastore.GetById(request.Id)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return category, nil
}

func (c CategoryService) GetAll(request *categoryRequest.GetAllRequest) (map[string]*model.Category, error) {
	categories, err := c.datastore.GetAll()
	if err != nil {
		return nil, fmt.Errorf("")
	}
	return categories, nil
}
