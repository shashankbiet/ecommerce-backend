package productkeyword

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

func TestAddRequestRejectsMissingReferences(t *testing.T) {
	request := httptest.NewRequest("POST", "/api/product-keyword", strings.NewReader(`{"productId":0,"keywordId":-1}`))
	if _, err := (AddRequest{}).Validate(request); err == nil {
		t.Fatal("Validate() accepted invalid references")
	}
}

func TestDeleteRequestRejectsInvalidRouteIDs(t *testing.T) {
	request := httptest.NewRequest("DELETE", "/api/product-keyword/bad/0", nil)
	request = mux.SetURLVars(request, map[string]string{"productId": "bad", "keywordId": "0"})
	if _, err := (DeleteRequest{}).Validate(request); err == nil {
		t.Fatal("Validate() accepted invalid route ids")
	}
}
