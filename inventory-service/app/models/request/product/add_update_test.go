package product

import (
	"fmt"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAddRequestValidateAcceptsIdBasedProductFields(t *testing.T) {
	body := `{"name":"Phone","description":"desc","slug":"phone","brand":1,"category":2,"subCategory":3,"imageUrl":"img","status":1}`
	req := httptest.NewRequest("POST", "/api/product", strings.NewReader(body))

	_, err := AddRequest{}.Validate(req)
	if err != nil {
		t.Fatalf("expected request to be valid, got error: %v", err)
	}
}

func TestAddRequestValidateAcceptsSnakeCaseFields(t *testing.T) {
	body := `{"name":"Phone","brand":1,"category":2,"sub_category":3,"image_url":"img"}`
	req := httptest.NewRequest("POST", "/api/product", strings.NewReader(body))

	value, err := AddRequest{}.Validate(req)
	if err != nil {
		t.Fatalf("expected request to be valid, got error: %v", err)
	}
	product := value.(*AddRequest).Product
	if product.SubCategory != 3 || product.ImageUrl != "img" {
		t.Fatalf("decoded product = %#v", product)
	}
}

func TestAddRequestValidateRejectsNonPositiveReferences(t *testing.T) {
	body := `{"name":"Phone","brand":0,"category":0,"subCategory":0}`
	req := httptest.NewRequest("POST", "/api/product", strings.NewReader(body))

	_, err := AddRequest{}.Validate(req)
	if err == nil {
		t.Fatal("expected validation error for invalid references")
	}
}

func TestProductRequestsAcceptDefinedStatuses(t *testing.T) {
	for _, status := range []int{0, 1, 2} {
		t.Run(fmt.Sprintf("status_%d", status), func(t *testing.T) {
			addBody := fmt.Sprintf(`{"name":"Phone","brand":1,"category":2,"sub_category":3,"status":%d}`, status)
			addRequest := httptest.NewRequest("POST", "/api/product", strings.NewReader(addBody))
			if _, err := (AddRequest{}).Validate(addRequest); err != nil {
				t.Fatalf("AddRequest.Validate() accepted status %d with error: %v", status, err)
			}

			updateBody := fmt.Sprintf(`{"id":1,"name":"Phone","brand":1,"category":2,"sub_category":3,"status":%d}`, status)
			updateRequest := httptest.NewRequest("PUT", "/api/product", strings.NewReader(updateBody))
			if _, err := (UpdateRequest{}).Validate(updateRequest); err != nil {
				t.Fatalf("UpdateRequest.Validate() accepted status %d with error: %v", status, err)
			}
		})
	}
}

func TestProductRequestsRejectUndefinedStatuses(t *testing.T) {
	for _, status := range []int{-1, 3, 7} {
		t.Run(fmt.Sprintf("status_%d", status), func(t *testing.T) {
			addBody := fmt.Sprintf(`{"name":"Phone","brand":1,"category":2,"sub_category":3,"status":%d}`, status)
			addRequest := httptest.NewRequest("POST", "/api/product", strings.NewReader(addBody))
			if _, err := (AddRequest{}).Validate(addRequest); err == nil {
				t.Fatalf("AddRequest.Validate() accepted undefined status %d", status)
			}

			updateBody := fmt.Sprintf(`{"id":1,"name":"Phone","brand":1,"category":2,"sub_category":3,"status":%d}`, status)
			updateRequest := httptest.NewRequest("PUT", "/api/product", strings.NewReader(updateBody))
			if _, err := (UpdateRequest{}).Validate(updateRequest); err == nil {
				t.Fatalf("UpdateRequest.Validate() accepted undefined status %d", status)
			}
		})
	}
}
