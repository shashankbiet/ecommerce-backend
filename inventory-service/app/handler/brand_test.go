package handler

import (
	"encoding/json"
	model "inventory-service/app/models"
	brandRequest "inventory-service/app/models/request/brand"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type brandServiceStub struct {
	brands []*model.Brand
}

func (s *brandServiceStub) GetAll(request *brandRequest.GetAllRequest) ([]*model.Brand, error) {
	return s.brands, nil
}

func TestBrandHandlerGetAllReturnsJSONList(t *testing.T) {
	createdAt := time.Date(2026, time.November, 22, 0, 0, 0, 0, time.UTC)
	handler := NewBrandHandler(&brandServiceStub{
		brands: []*model.Brand{{Id: 1, Name: "Acme", CreatedAt: createdAt}},
	})
	recorder := httptest.NewRecorder()

	handler.GetAll(recorder, httptest.NewRequest(http.MethodGet, "/api/brand", nil))

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", recorder.Code, http.StatusOK)
	}
	if got := recorder.Header().Get("Content-Type"); got != "application/json" {
		t.Fatalf("Content-Type = %q, want application/json", got)
	}
	var response []*model.Brand
	if err := json.Unmarshal(recorder.Body.Bytes(), &response); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if len(response) != 1 || response[0].Id != 1 || response[0].Name != "Acme" {
		t.Fatalf("response = %#v", response)
	}
}
