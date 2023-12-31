package service

import (
	"encoding/json"
	"fmt"
	"inventory-service/app/dao"
	model "inventory-service/app/models"
	inventoryRequest "inventory-service/app/models/request/inventory"
	"inventory-service/app/producer"
	"inventory-service/pkg/logger"
)

type InventoryService struct {
	datastore dao.IInventoryDataStore
	producer  producer.IProductProducer
}

func NewInventoryService(datastore dao.IInventoryDataStore, producer producer.IInventoryProducer) IInventoryService {
	return &InventoryService{
		datastore: datastore,
		producer:  producer,
	}
}

func (p *InventoryService) Add(request *inventoryRequest.AddRequest) (int64, error) {
	id, err := p.datastore.Add(request.Inventory)
	if err != nil {
		return 0, err
	}
	p.publishInventoryMessage(request.Inventory)
	return id, nil
}

func (p *InventoryService) Update(request *inventoryRequest.UpdateRequest) (bool, error) {
	_, err := p.datastore.Update(request.Inventory)
	if err != nil {
		return false, err
	}
	p.publishInventoryMessage(request.Inventory)
	return true, nil
}

func (p *InventoryService) GetByProductId(request *inventoryRequest.GetByIdRequest) (*model.Inventory, error) {
	inventory, err := p.datastore.GetByProductId(request.ProductId)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return inventory, nil
}

func (p *InventoryService) GetAll(request *inventoryRequest.GetAllRequest) ([]*model.Inventory, error) {
	inventories, err := p.datastore.GetAll()
	if err != nil {
		return nil, fmt.Errorf("")
	}
	return inventories, nil
}

func (p *InventoryService) publishInventoryMessage(inventory *model.Inventory) {
	arr, err := json.Marshal(inventory)
	if err != nil {
		logger.Log.Error("unable to marshal inventory struct into []byte")
	}
	p.producer.Publish(arr)
}
