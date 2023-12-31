package service

import (
	"fmt"
	"inventory-service/app/dao"
	model "inventory-service/app/models"
	subCategoryRequest "inventory-service/app/models/request/subcategory"
)

type SubCategoryService struct {
	datastore dao.ISubCategoryDataStore
}

func NewSubCategoryService(datastore dao.ISubCategoryDataStore) ISubCategoryService {
	return &SubCategoryService{
		datastore: datastore,
	}
}

func (s *SubCategoryService) Add(request *subCategoryRequest.AddRequest) (int16, error) {
	setSubCategory(&request.SubCategory.Name)
	setCategory(&request.SubCategory.Category)
	id, err := s.datastore.Add(request.SubCategory)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *SubCategoryService) GetById(request *subCategoryRequest.GetByIdRequest) (*model.SubCategory, error) {
	subcategory, err := s.datastore.GetById(request.Id)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return subcategory, nil
}

func (s *SubCategoryService) GetAll(request *subCategoryRequest.GetAllRequest) (map[string]*model.SubCategory, error) {
	subcategories, err := s.datastore.GetAll()
	if err != nil {
		return nil, fmt.Errorf("")
	}
	return subcategories, nil
}
