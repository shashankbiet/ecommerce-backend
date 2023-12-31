package handler

import (
	"encoding/json"
	"inventory-service/app/constants"
	"inventory-service/app/models/request"
	productRequest "inventory-service/app/models/request/product"
	"inventory-service/app/service"
	"net/http"
)

type ProductHandler struct {
	service service.IProductService
}

func NewProductHandler(service service.IProductService) *ProductHandler {
	return &ProductHandler{
		service: service,
	}
}

func (p ProductHandler) Add(w http.ResponseWriter, r *http.Request) {
	var validateRequest request.IValidateRequest = &productRequest.AddRequest{}
	val, err := validateRequest.Validate(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	request := val.(*productRequest.AddRequest)

	_, err = p.service.Add(request)
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

func (p ProductHandler) Update(w http.ResponseWriter, r *http.Request) {
	var validateRequest request.IValidateRequest = &productRequest.UpdateRequest{}
	val, err := validateRequest.Validate(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	request := val.(*productRequest.UpdateRequest)

	isUpdated, err := p.service.Update(request)
	if err != nil || !isUpdated {
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

func (p ProductHandler) GetById(w http.ResponseWriter, r *http.Request) {
	var validateRequest request.IValidateRequest = &productRequest.GetByIdRequest{}
	val, err := validateRequest.Validate(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	request := val.(*productRequest.GetByIdRequest)

	product, err := p.service.GetById(request)
	if err != nil {
		if err.Error() == constants.NOT_FOUND_ERROR_MESSAGE {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		return
	}

	jsonBytes, err := json.Marshal(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (p ProductHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	var validateRequest request.IValidateRequest = &productRequest.GetAllRequest{}
	val, err := validateRequest.Validate(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	request := val.(*productRequest.GetAllRequest)

	products, err := p.service.GetAll(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	jsonBytes, err := json.Marshal(products)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}
