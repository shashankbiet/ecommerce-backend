package consumer

import (
	"encoding/json"
	"errors"
	model "search-service/app/models"
	searchpb "search-service/proto/search"
	"testing"
	"time"
)

type productStoreStub struct {
	upserted  *model.Product
	upsertErr error
}

func (*productStoreStub) GetProductById(string) (*searchpb.Product, error) {
	return nil, nil
}

func (*productStoreStub) FilterProduct(*searchpb.ProductSearchRequest) (*searchpb.ProductSearchResponse, error) {
	return nil, nil
}

func (s *productStoreStub) UpsertProduct(productID int64, product *model.Product) error {
	s.upserted = product
	return s.upsertErr
}

func (*productStoreStub) DeleteProduct(int64) error {
	return nil
}

func TestUpdateProductIndexesAllEventFields(t *testing.T) {
	store := &productStoreStub{}
	consumer := NewProductConsumerWithDataStore(store)
	createdAt := time.Date(2026, time.November, 22, 1, 2, 3, 0, time.UTC)
	event := model.Product{
		Id:          9007199254740993,
		Name:        "Mobile",
		Description: "Smart phone",
		Slug:        "mobile",
		Brand:       2,
		Category:    3,
		SubCategory: 4,
		ImageUrl:    "https://example.test/mobile.png",
		Status:      1,
		CreatedAt:   createdAt,
		UpdatedAt:   createdAt.Add(time.Hour),
	}
	payload, err := json.Marshal(event)
	if err != nil {
		t.Fatalf("marshal event: %v", err)
	}

	if err := consumer.updateProduct(payload); err != nil {
		t.Fatalf("updateProduct() error = %v", err)
	}
	if store.upserted == nil {
		t.Fatal("product was not upserted")
	}
	if store.upserted.Id != event.Id ||
		store.upserted.Slug != event.Slug ||
		store.upserted.Brand != event.Brand ||
		store.upserted.Category != event.Category ||
		store.upserted.SubCategory != event.SubCategory ||
		store.upserted.ImageUrl != event.ImageUrl ||
		store.upserted.Status != event.Status ||
		!store.upserted.CreatedAt.Equal(event.CreatedAt) ||
		!store.upserted.UpdatedAt.Equal(event.UpdatedAt) {
		t.Fatalf("upserted product = %#v, want %#v", store.upserted, event)
	}
}

func TestUpdateProductRejectsInvalidProductID(t *testing.T) {
	consumer := NewProductConsumerWithDataStore(&productStoreStub{})
	if err := consumer.updateProduct([]byte(`{"id":0,"name":"invalid"}`)); err == nil {
		t.Fatal("updateProduct() accepted an invalid product id")
	}
}

func TestUpdateProductRejectsUndefinedStatus(t *testing.T) {
	consumer := NewProductConsumerWithDataStore(&productStoreStub{})
	if err := consumer.updateProduct([]byte(`{"id":42,"name":"Phone","status":7}`)); err == nil {
		t.Fatal("updateProduct() accepted undefined status 7")
	}
}

func TestUpdateProductAcceptsLegacyCamelCasePayload(t *testing.T) {
	store := &productStoreStub{}
	consumer := NewProductConsumerWithDataStore(store)
	payload := []byte(`{"id":42,"subCategory":3,"imageUrl":"legacy.png","createdAt":"2026-11-22T01:02:03Z"}`)

	if err := consumer.updateProduct(payload); err != nil {
		t.Fatalf("updateProduct() legacy payload error = %v", err)
	}
	if store.upserted == nil ||
		store.upserted.SubCategory != 3 ||
		store.upserted.ImageUrl != "legacy.png" {
		t.Fatalf("upserted legacy product = %#v", store.upserted)
	}
}

func TestUpdateProductReturnsMalformedPayloadError(t *testing.T) {
	consumer := NewProductConsumerWithDataStore(&productStoreStub{})
	if err := consumer.updateProduct([]byte(`{"id":`)); err == nil {
		t.Fatal("updateProduct() accepted malformed JSON")
	}
}

func TestUpdateProductReturnsUpsertError(t *testing.T) {
	expected := errors.New("elasticsearch unavailable")
	consumer := NewProductConsumerWithDataStore(&productStoreStub{upsertErr: expected})
	err := consumer.updateProduct([]byte(`{"id":42,"name":"Phone"}`))
	if err == nil {
		t.Fatal("updateProduct() error = nil, want datastore error")
	}
	if !errors.Is(err, expected) {
		t.Fatalf("updateProduct() error = %v, want wrapped %v", err, expected)
	}
}
