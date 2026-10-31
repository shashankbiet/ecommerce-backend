package service

import (
	"encoding/json"
	"inventory-service/app/dao"
	model "inventory-service/app/models"
	inventoryRequest "inventory-service/app/models/request/inventory"
	"inventory-service/app/producer"
	"inventory-service/pkg/logger"
)

type InventoryService struct {
	datastore dao.IInventoryDataStore
	producer  producer.IInventoryProducer
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
	updated, err := p.datastore.Update(request.Inventory)
	if err != nil {
		return false, err
	}
	if !updated {
		return false, nil
	}
	p.publishInventoryMessage(request.Inventory)
	return true, nil
}

func (p *InventoryService) GetByProductId(request *inventoryRequest.GetByIdRequest) (*model.Inventory, error) {
	inventory, err := p.datastore.GetByProductId(request.ProductId)
	if err != nil {
		return nil, err
	}

	return inventory, nil
}

func (p *InventoryService) GetAll(request *inventoryRequest.GetAllRequest) ([]*model.Inventory, error) {
	inventories, err := p.datastore.GetAll()
	if err != nil {
		return nil, err
	}
	return inventories, nil
}

func (p *InventoryService) publishInventoryMessage(inventory *model.Inventory) {
	arr, err := json.Marshal(inventory)
	if err != nil {
		logger.Error("unable to marshal inventory struct into []byte", "error", err)
		return
	}
	p.producer.Publish(arr)
}
