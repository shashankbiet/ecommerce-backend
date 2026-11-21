package service

import (
	"errors"
	model "inventory-service/app/models"
	brandRequest "inventory-service/app/models/request/brand"
	"testing"
)

type brandStoreStub struct {
	brands []*model.Brand
	err    error
}

func (s *brandStoreStub) GetAll() ([]*model.Brand, error) {
	return s.brands, s.err
}

func TestBrandServiceGetAllReturnsBrands(t *testing.T) {
	want := []*model.Brand{{Id: 1, Name: "Acme"}}
	service := NewBrandService(&brandStoreStub{brands: want})

	got, err := service.GetAll(&brandRequest.GetAllRequest{})
	if err != nil {
		t.Fatalf("GetAll() error = %v", err)
	}
	if len(got) != 1 || got[0] != want[0] {
		t.Fatalf("GetAll() = %#v, want %#v", got, want)
	}
}

func TestBrandServiceGetAllReturnsStoreError(t *testing.T) {
	wantErr := errors.New("read failed")
	service := NewBrandService(&brandStoreStub{err: wantErr})

	_, err := service.GetAll(&brandRequest.GetAllRequest{})
	if !errors.Is(err, wantErr) {
		t.Fatalf("GetAll() error = %v, want %v", err, wantErr)
	}
}
