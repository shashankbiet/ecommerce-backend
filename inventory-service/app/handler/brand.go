package handler

import (
	"encoding/json"
	"inventory-service/app/models/request"
	brandRequest "inventory-service/app/models/request/brand"
	"inventory-service/app/service"
	"net/http"
)

type BrandHandler struct {
	service service.IBrandService
}

func NewBrandHandler(service service.IBrandService) *BrandHandler {
	return &BrandHandler{service: service}
}

func (b BrandHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	var validateRequest request.IValidateRequest = &brandRequest.GetAllRequest{}
	value, err := validateRequest.Validate(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	brands, err := b.service.GetAll(value.(*brandRequest.GetAllRequest))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(brands); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var _ request.IValidateRequest = (*brandRequest.GetAllRequest)(nil)
