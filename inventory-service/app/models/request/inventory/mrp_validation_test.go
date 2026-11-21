package inventory

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAddRequestValidateRejectsNegativeMRP(t *testing.T) {
	body := `{"productId":1,"sku":100,"mrp":-1,"purchasePrice":10,"salePrice":20,"quantity":1}`
	req := httptest.NewRequest("POST", "/api/inventory", strings.NewReader(body))

	_, err := AddRequest{}.Validate(req)
	if err == nil {
		t.Fatal("expected validation error for negative MRP")
	}
}

func TestUpdateRequestValidateAcceptsNonNegativeMRP(t *testing.T) {
	body := `{"productId":1,"sku":100,"mrp":0,"purchasePrice":10,"salePrice":20,"quantity":1}`
	req := httptest.NewRequest("PUT", "/api/inventory", strings.NewReader(body))

	_, err := UpdateRequest{}.Validate(req)
	if err != nil {
		t.Fatalf("expected request to be valid, got error: %v", err)
	}
}
