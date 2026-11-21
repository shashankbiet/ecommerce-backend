package keyword

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAddRequestRejectsBlankKeyword(t *testing.T) {
	request := httptest.NewRequest("POST", "/api/keyword", strings.NewReader(`{"name":"  "}`))
	if _, err := (AddRequest{}).Validate(request); err == nil {
		t.Fatal("Validate() accepted blank keyword")
	}
}

func TestAddRequestAcceptsKeyword(t *testing.T) {
	request := httptest.NewRequest("POST", "/api/keyword", strings.NewReader(`{"name":"best seller"}`))
	if _, err := (AddRequest{}).Validate(request); err != nil {
		t.Fatalf("Validate() error = %v", err)
	}
}
