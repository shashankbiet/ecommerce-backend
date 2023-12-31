package handler

import (
	"encoding/json"
	"inventory-service/app/constants"
	"inventory-service/app/models/request"
	inventoryRequest "inventory-service/app/models/request/inventory"
	"inventory-service/app/service"
	"net/http"
)

type InventoryHandler struct {
	service service.IInventoryService
}

func NewInventoryHandler(service service.IInventoryService) *InventoryHandler {
	return &InventoryHandler{
		service: service,
	}
}

func (p InventoryHandler) Add(w http.ResponseWriter, r *http.Request) {
	var validateRequest request.IValidateRequest = &inventoryRequest.AddRequest{}
	val, err := validateRequest.Validate(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	request := val.(*inventoryRequest.AddRequest)

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

func (p InventoryHandler) Update(w http.ResponseWriter, r *http.Request) {
	var validateRequest request.IValidateRequest = &inventoryRequest.UpdateRequest{}
	val, err := validateRequest.Validate(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	request := val.(*inventoryRequest.UpdateRequest)

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

func (p InventoryHandler) GetById(w http.ResponseWriter, r *http.Request) {
	var validateRequest request.IValidateRequest = &inventoryRequest.GetByIdRequest{}
	val, err := validateRequest.Validate(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	request := val.(*inventoryRequest.GetByIdRequest)

	inventory, err := p.service.GetByProductId(request)
	if err != nil {
		if err.Error() == constants.NOT_FOUND_ERROR_MESSAGE {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		return
	}

	jsonBytes, err := json.Marshal(inventory)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (p InventoryHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	var validateRequest request.IValidateRequest = &inventoryRequest.GetAllRequest{}
	val, err := validateRequest.Validate(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	request := val.(*inventoryRequest.GetAllRequest)

	inventories, err := p.service.GetAll(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	jsonBytes, err := json.Marshal(inventories)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}
