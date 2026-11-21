package service

import (
	model "inventory-service/app/models"
	keywordRequest "inventory-service/app/models/request/keyword"
	"testing"
)

type keywordStoreStub struct {
	added *model.Keyword
}

func (s *keywordStoreStub) Add(keyword *model.Keyword) (int64, error) {
	s.added = keyword
	return 1, nil
}

func (*keywordStoreStub) GetById(int64) (*model.Keyword, error) {
	return nil, nil
}

func (*keywordStoreStub) GetAll() (map[string]*model.Keyword, error) {
	return nil, nil
}

func TestKeywordServiceAddNormalizesName(t *testing.T) {
	store := &keywordStoreStub{}
	service := NewKeywordService(store)
	request := &keywordRequest.AddRequest{Keyword: &model.Keyword{Name: " best seller "}}

	id, err := service.Add(request)
	if err != nil {
		t.Fatalf("Add() error = %v", err)
	}
	if id != 1 {
		t.Fatalf("Add() = %d, want 1", id)
	}
	if store.added.Name != "BEST_SELLER" {
		t.Fatalf("keyword name = %q, want %q", store.added.Name, "BEST_SELLER")
	}
}
