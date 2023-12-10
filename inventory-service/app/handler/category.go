package handler

import (
	"encoding/json"
	"inventory-service/app/constants"
	intf "inventory-service/app/interface"
	categoryRequest "inventory-service/app/models/request/category"
	"net/http"
)

type CategoryHandler struct {
	service intf.CategoryService
}

func NewCategoryHandler(service intf.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		service: service,
	}
}

func (c CategoryHandler) Add(w http.ResponseWriter, r *http.Request) {
	var validateRequest intf.ValidateRequest = &categoryRequest.AddRequest{}
	val, err := validateRequest.Validate(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
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
	var validateRequest intf.ValidateRequest = &categoryRequest.GetByIdRequest{}
	val, err := validateRequest.Validate(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	request := val.(*categoryRequest.GetByIdRequest)

	categories, err := c.service.GetById(request)
	if err != nil {
		if err.Error() == constants.NOT_FOUND_ERROR_MESSAGE {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
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

func (c CategoryHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	var validateRequest intf.ValidateRequest = &categoryRequest.GetAllRequest{}
	val, err := validateRequest.Validate(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
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
