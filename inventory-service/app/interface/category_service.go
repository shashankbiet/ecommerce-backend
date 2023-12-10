package intf

import (
	model "inventory-service/app/models"
	categoryRequest "inventory-service/app/models/request/category"
)

type CategoryService interface {
	Add(request *categoryRequest.AddRequest) (int64, error)
	GetById(request *categoryRequest.GetByIdRequest) (*model.Category, error)
	GetAll(request *categoryRequest.GetAllRequest) (map[string]*model.Category, error)
}
