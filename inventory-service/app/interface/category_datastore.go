package intf

import model "inventory-service/app/models"

type CategoryDataStore interface {
	Add(category *model.Category) (int64, error)
	GetById(id int64) (*model.Category, error)
	GetAll() (map[string]*model.Category, error)
}
