package service

import (
	model "inventory-service/app/models"
	productKeywordRequest "inventory-service/app/models/request/product_keyword"
	"testing"
)

type productKeywordStoreStub struct {
	added     *model.ProductKeyword
	deleted   [2]int64
	productID int64
	keywordID int64
}

func (s *productKeywordStoreStub) Add(value *model.ProductKeyword) (int64, error) {
	s.added = value
	return value.ProductId, nil
}

func (s *productKeywordStoreStub) GetByProductId(id int64) ([]*model.ProductKeyword, error) {
	s.productID = id
	return []*model.ProductKeyword{}, nil
}

func (s *productKeywordStoreStub) GetByKeywordId(id int64) ([]*model.ProductKeyword, error) {
	s.keywordID = id
	return []*model.ProductKeyword{}, nil
}

func (s *productKeywordStoreStub) DeleteById(productID int64, keywordID int64) error {
	s.deleted = [2]int64{productID, keywordID}
	return nil
}

func TestProductKeywordServiceDelegatesOperations(t *testing.T) {
	store := &productKeywordStoreStub{}
	service := NewProductKeywordService(store)

	association := &model.ProductKeyword{ProductId: 4, KeywordId: 7}
	_, err := service.Add(&productKeywordRequest.AddRequest{ProductKeyword: association})
	if err != nil {
		t.Fatalf("Add() error = %v", err)
	}
	_, err = service.GetByProductId(&productKeywordRequest.GetByProductIdRequest{ProductId: 4})
	if err != nil {
		t.Fatalf("GetByProductId() error = %v", err)
	}
	_, err = service.GetByKeywordId(&productKeywordRequest.GetByKeywordIdRequest{KeywordId: 7})
	if err != nil {
		t.Fatalf("GetByKeywordId() error = %v", err)
	}
	err = service.Delete(&productKeywordRequest.DeleteRequest{ProductId: 4, KeywordId: 7})
	if err != nil {
		t.Fatalf("Delete() error = %v", err)
	}

	if store.added != association || store.productID != 4 || store.keywordID != 7 || store.deleted != [2]int64{4, 7} {
		t.Fatalf("service did not delegate the expected values: %#v", store)
	}
}
