package handler

import (
	"encoding/json"
	"inventory-service/app/constants"
	"inventory-service/app/models/request"
	keywordRequest "inventory-service/app/models/request/keyword"
	"inventory-service/app/service"
	"net/http"
)

type KeywordHandler struct {
	service service.IKeywordService
}

func NewKeywordHandler(service service.IKeywordService) *KeywordHandler {
	return &KeywordHandler{service: service}
}

func (k KeywordHandler) Add(w http.ResponseWriter, r *http.Request) {
	value, err := (&keywordRequest.AddRequest{}).Validate(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_, err = k.service.Add(value.(*keywordRequest.AddRequest))
	if err != nil {
		if err.Error() == constants.CONFLICT_ERROR_MESSAGE {
			w.WriteHeader(http.StatusConflict)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (k KeywordHandler) GetById(w http.ResponseWriter, r *http.Request) {
	value, err := (&keywordRequest.GetByIdRequest{}).Validate(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	keyword, err := k.service.GetById(value.(*keywordRequest.GetByIdRequest))
	if err != nil {
		if err.Error() == constants.NOT_FOUND_ERROR_MESSAGE {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, keyword)
}

func (k KeywordHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	value, err := (&keywordRequest.GetAllRequest{}).Validate(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	keywords, err := k.service.GetAll(value.(*keywordRequest.GetAllRequest))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, keywords)
}

func writeJSON(w http.ResponseWriter, value interface{}) {
	jsonBytes, err := json.Marshal(value)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(jsonBytes)
}

var _ request.IValidateRequest = (*keywordRequest.AddRequest)(nil)
