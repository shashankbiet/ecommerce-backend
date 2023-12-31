package handler

import (
	"encoding/json"
	"inventory-service/app/constants"
	"inventory-service/app/models/request"
	categoryRequest "inventory-service/app/models/request/category"
	"inventory-service/app/service"
	"net/http"
)

type CategoryHandler struct {
	service service.ICategoryService
}

func NewCategoryHandler(service service.ICategoryService) *CategoryHandler {
	return &CategoryHandler{
		service: service,
	}
}

func (c CategoryHandler) Add(w http.ResponseWriter, r *http.Request) {
	var validateRequest request.IValidateRequest = &categoryRequest.AddRequest{}
	val, err := validateRequest.Validate(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	request := val.(*categoryRequest.AddRequest)

	_, err = c.service.Add(request)
	if err != nil {
		if err.Error() == constants.CONFLICT_ERROR_MESSAGE {
			w.WriteHeader(http.StatusConflict)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (c CategoryHandler) GetById(w http.ResponseWriter, r *http.Request) {
	var validateRequest request.IValidateRequest = &categoryRequest.GetByIdRequest{}
	val, err := validateRequest.Validate(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	request := val.(*categoryRequest.GetByIdRequest)

	category, err := c.service.GetById(request)
	if err != nil {
		if err.Error() == constants.NOT_FOUND_ERROR_MESSAGE {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		return
	}

	jsonBytes, err := json.Marshal(category)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (c CategoryHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	var validateRequest request.IValidateRequest = &categoryRequest.GetAllRequest{}
	val, err := validateRequest.Validate(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	request := val.(*categoryRequest.GetAllRequest)

	categories, err := c.service.GetAll(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	jsonBytes, err := json.Marshal(categories)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}
