package service

import (
	"encoding/json"
	model "inventory-service/app/models"
	inventoryRequest "inventory-service/app/models/request/inventory"
	"testing"
)

type inventoryStoreStub struct {
	added *model.Inventory
}

func (s *inventoryStoreStub) Add(inventory *model.Inventory) (int64, error) {
	s.added = inventory
	return inventory.ProductId, nil
}

func (*inventoryStoreStub) Update(*model.Inventory) (bool, error) {
	return false, nil
}

func (*inventoryStoreStub) GetByProductId(int64) (*model.Inventory, error) {
	return nil, nil
}

func (*inventoryStoreStub) GetAll() ([]*model.Inventory, error) {
	return nil, nil
}

type inventoryProducerStub struct {
	payloads [][]byte
}

func (p *inventoryProducerStub) Publish(value []byte) {
	p.payloads = append(p.payloads, value)
}

func TestInventoryServiceAddPublishesMRP(t *testing.T) {
	store := &inventoryStoreStub{}
	producer := &inventoryProducerStub{}
	service := NewInventoryService(store, producer)
	inventory := &model.Inventory{
		ProductId:     12,
		SKU:           100,
		MRP:           999,
		PurchasePrice: 400,
		SalePrice:     750,
		Quantity:      3,
	}

	if _, err := service.Add(&inventoryRequest.AddRequest{Inventory: inventory}); err != nil {
		t.Fatalf("Add() error = %v", err)
	}
	if len(producer.payloads) != 1 {
		t.Fatalf("published payload count = %d, want 1", len(producer.payloads))
	}
	var fields map[string]interface{}
	if err := json.Unmarshal(producer.payloads[0], &fields); err != nil {
		t.Fatalf("decode inventory event fields: %v", err)
	}
	for _, field := range []string{"product_id", "purchase_price", "sale_price", "created_at", "updated_at"} {
		if _, ok := fields[field]; !ok {
			t.Errorf("published inventory event missing %q: %s", field, producer.payloads[0])
		}
	}
	for _, field := range []string{"productId", "purchasePrice", "salePrice", "createdAt", "updatedAt"} {
		if _, ok := fields[field]; ok {
			t.Errorf("published inventory event contains legacy field %q: %s", field, producer.payloads[0])
		}
	}
	var event model.Inventory
	if err := json.Unmarshal(producer.payloads[0], &event); err != nil {
		t.Fatalf("decode inventory event: %v", err)
	}
	if event.MRP != inventory.MRP || event.ProductId != inventory.ProductId {
		t.Fatalf("published event = %#v", event)
	}
}
