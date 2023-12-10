package intf

import model "inventory-service/app/models"

type SubCategoryDataStore interface {
	Add(category *model.SubCategory) error
	Get(id int) (*model.SubCategory, error)
	List() (map[string]*model.SubCategory, error)
}
