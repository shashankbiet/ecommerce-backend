package handler

import (
	"encoding/json"
	"inventory-service/app/constants"
	"inventory-service/app/models/request"
	subCategoryRequest "inventory-service/app/models/request/subcategory"
	"inventory-service/app/service"
	"net/http"
)

type SubCategoryHandler struct {
	service service.ISubCategoryService
}

func NewSubCategoryHandler(service service.ISubCategoryService) *SubCategoryHandler {
	return &SubCategoryHandler{
		service: service,
	}
}

func (s *SubCategoryHandler) Add(w http.ResponseWriter, r *http.Request) {
	var validateRequest request.IValidateRequest = &subCategoryRequest.AddRequest{}
	val, err := validateRequest.Validate(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	request := val.(*subCategoryRequest.AddRequest)
	_, err = s.service.Add(request)
	if err != nil {
		if err.Error() == constants.CONFLICT_ERROR_MESSAGE {
			w.WriteHeader(http.StatusConflict)
		} else if err.Error() == constants.UNPROCESSABLE_ENTITY_ERROR_MESSAGE {
			w.WriteHeader(http.StatusUnprocessableEntity)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (s *SubCategoryHandler) GetById(w http.ResponseWriter, r *http.Request) {
	var validateRequest request.IValidateRequest = &subCategoryRequest.GetByIdRequest{}
	val, err := validateRequest.Validate(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	request := val.(*subCategoryRequest.GetByIdRequest)

	subcategory, err := s.service.GetById(request)
	if err != nil {
		if err.Error() == constants.NOT_FOUND_ERROR_MESSAGE {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		return
	}

	jsonBytes, err := json.Marshal(subcategory)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (s *SubCategoryHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	var validateRequest request.IValidateRequest = &subCategoryRequest.GetAllRequest{}
	val, err := validateRequest.Validate(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	request := val.(*subCategoryRequest.GetAllRequest)

	subcategories, err := s.service.GetAll(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	jsonBytes, err := json.Marshal(subcategories)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}
