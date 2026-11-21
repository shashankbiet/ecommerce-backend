package service

import (
	"encoding/json"
	"errors"
	model "inventory-service/app/models"
	productRequest "inventory-service/app/models/request/product"
	"testing"
	"time"
)

type productStoreStub struct {
	addID       int64
	addErr      error
	updated     bool
	updateErr   error
	createdAt   time.Time
	updatedAt   time.Time
	added       *model.Product
	updateValue *model.Product
}

func (s *productStoreStub) Add(product *model.Product) (int64, error) {
	s.added = product
	if s.addErr == nil {
		product.CreatedAt = s.createdAt
		product.UpdatedAt = s.updatedAt
	}
	return s.addID, s.addErr
}

func (s *productStoreStub) Update(product *model.Product) (bool, error) {
	s.updateValue = product
	return s.updated, s.updateErr
}

func (*productStoreStub) GetById(int64) (*model.Product, error) {
	return nil, nil
}

func (*productStoreStub) GetAll() ([]*model.Product, error) {
	return nil, nil
}

type productProducerStub struct {
	payloads [][]byte
}

func (p *productProducerStub) Publish(value []byte) {
	p.payloads = append(p.payloads, value)
}

func TestProductServiceAddPublishesPersistedModel(t *testing.T) {
	createdAt := time.Date(2026, time.November, 22, 1, 2, 3, 0, time.UTC)
	updatedAt := createdAt.Add(time.Minute)
	store := &productStoreStub{
		addID:     42,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
	producer := &productProducerStub{}
	service := NewProductService(store, producer)
	product := &model.Product{
		Name:        "Phone",
		Description: "Smart phone",
		Slug:        "phone",
		Brand:       2,
		Category:    3,
		SubCategory: 4,
		ImageUrl:    "phone.png",
		Status:      1,
	}

	id, err := service.Add(&productRequest.AddRequest{Product: product})
	if err != nil {
		t.Fatalf("Add() error = %v", err)
	}
	if id != 42 || product.Id != 42 {
		t.Fatalf("Add() id = %d, product id = %d; want 42", id, product.Id)
	}
	if len(producer.payloads) != 1 {
		t.Fatalf("published payload count = %d, want 1", len(producer.payloads))
	}

	var fields map[string]interface{}
	if err := json.Unmarshal(producer.payloads[0], &fields); err != nil {
		t.Fatalf("decode product event fields: %v", err)
	}
	for _, field := range []string{"sub_category", "image_url", "created_at", "updated_at", "deleted_at"} {
		if _, ok := fields[field]; !ok {
			t.Errorf("published product event missing %q: %s", field, producer.payloads[0])
		}
	}
	for _, field := range []string{"subCategory", "imageUrl", "createdAt", "updatedAt", "deletedAt"} {
		if _, ok := fields[field]; ok {
			t.Errorf("published product event contains legacy field %q: %s", field, producer.payloads[0])
		}
	}

	var event model.Product
	if err := json.Unmarshal(producer.payloads[0], &event); err != nil {
		t.Fatalf("decode product event: %v", err)
	}
	if event.Id != 42 || event.Slug != "phone" || event.Brand != 2 ||
		event.Category != 3 || event.SubCategory != 4 ||
		event.ImageUrl != "phone.png" || event.Status != 1 ||
		!event.CreatedAt.Equal(createdAt) || !event.UpdatedAt.Equal(updatedAt) {
		t.Fatalf("published event = %#v", event)
	}
}

func TestProductServiceDoesNotPublishFailedAdd(t *testing.T) {
	store := &productStoreStub{addErr: errors.New("write failed")}
	producer := &productProducerStub{}
	service := NewProductService(store, producer)

	_, err := service.Add(&productRequest.AddRequest{Product: &model.Product{Name: "Phone"}})
	if err == nil {
		t.Fatal("Add() error = nil, want write error")
	}
	if len(producer.payloads) != 0 {
		t.Fatalf("published payload count = %d, want 0", len(producer.payloads))
	}
}

func TestProductServicePublishesOnlySuccessfulUpdates(t *testing.T) {
	t.Run("publishes updated model", func(t *testing.T) {
		store := &productStoreStub{updated: true}
		producer := &productProducerStub{}
		service := NewProductService(store, producer)
		product := &model.Product{Id: 7, Name: "Updated phone", Slug: "updated-phone"}

		updated, err := service.Update(&productRequest.UpdateRequest{Product: product})
		if err != nil {
			t.Fatalf("Update() error = %v", err)
		}
		if !updated || len(producer.payloads) != 1 {
			t.Fatalf("updated = %v, payload count = %d; want true, 1", updated, len(producer.payloads))
		}
	})

	t.Run("does not publish missing product", func(t *testing.T) {
		store := &productStoreStub{updated: false}
		producer := &productProducerStub{}
		service := NewProductService(store, producer)

		updated, err := service.Update(&productRequest.UpdateRequest{Product: &model.Product{Id: 404}})
		if err != nil {
			t.Fatalf("Update() error = %v", err)
		}
		if updated || len(producer.payloads) != 0 {
			t.Fatalf("updated = %v, payload count = %d; want false, 0", updated, len(producer.payloads))
		}
	})
}
