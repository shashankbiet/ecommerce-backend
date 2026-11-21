package productkeyword

import (
	"fmt"
	model "inventory-service/app/models"
	"net/http"
	"strings"
)

type AddRequest struct {
	ProductKeyword *model.ProductKeyword
}

func (AddRequest) Validate(r *http.Request) (interface{}, error) {
	var request AddRequest
	if err := model.DecodeJSON(r.Body, &request.ProductKeyword); err != nil {
		return nil, err
	}
	if request.ProductKeyword == nil {
		return nil, fmt.Errorf("request body cannot be empty")
	}

	errors := make([]string, 0)
	if request.ProductKeyword.ProductId <= 0 {
		errors = append(errors, "Invalid ProductId")
	}
	if request.ProductKeyword.KeywordId <= 0 {
		errors = append(errors, "Invalid KeywordId")
	}
	if len(errors) > 0 {
		return nil, fmt.Errorf("%s", strings.Join(errors, ","))
	}
	return &request, nil
}
