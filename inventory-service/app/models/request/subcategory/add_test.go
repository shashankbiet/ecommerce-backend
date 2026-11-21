package subcategory

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAddRequestValidateRequiresPositiveCategoryId(t *testing.T) {
	body := `{"name":"Mobiles","category":0}`
	req := httptest.NewRequest("POST", "/api/subcategory", strings.NewReader(body))

	_, err := AddRequest{}.Validate(req)
	if err == nil {
		t.Fatal("expected validation error for non-positive category id")
	}
}
