package handler

import (
	"inventory-service/app/constants"
	"inventory-service/app/models/request"
	productKeywordRequest "inventory-service/app/models/request/product_keyword"
	"inventory-service/app/service"
	"net/http"
)

type ProductKeywordHandler struct {
	service service.IProductKeywordService
}

func NewProductKeywordHandler(service service.IProductKeywordService) *ProductKeywordHandler {
	return &ProductKeywordHandler{service: service}
}

func (p ProductKeywordHandler) Add(w http.ResponseWriter, r *http.Request) {
	value, err := (&productKeywordRequest.AddRequest{}).Validate(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_, err = p.service.Add(value.(*productKeywordRequest.AddRequest))
	if err != nil {
		switch err.Error() {
		case constants.CONFLICT_ERROR_MESSAGE:
			w.WriteHeader(http.StatusConflict)
		case constants.UNPROCESSABLE_ENTITY_ERROR_MESSAGE:
			w.WriteHeader(http.StatusUnprocessableEntity)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (p ProductKeywordHandler) GetByProductId(w http.ResponseWriter, r *http.Request) {
	value, err := (&productKeywordRequest.GetByProductIdRequest{}).Validate(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	associations, err := p.service.GetByProductId(value.(*productKeywordRequest.GetByProductIdRequest))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, associations)
}

func (p ProductKeywordHandler) GetByKeywordId(w http.ResponseWriter, r *http.Request) {
	value, err := (&productKeywordRequest.GetByKeywordIdRequest{}).Validate(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	associations, err := p.service.GetByKeywordId(value.(*productKeywordRequest.GetByKeywordIdRequest))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, associations)
}

func (p ProductKeywordHandler) Delete(w http.ResponseWriter, r *http.Request) {
	value, err := (&productKeywordRequest.DeleteRequest{}).Validate(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := p.service.Delete(value.(*productKeywordRequest.DeleteRequest)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

var _ request.IValidateRequest = (*productKeywordRequest.AddRequest)(nil)
