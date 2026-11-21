package product

import (
	"fmt"
	model "inventory-service/app/models"
	"net/http"
	"strings"
)

type AddRequest struct {
	Product *model.Product
}

func (c AddRequest) Validate(r *http.Request) (interface{}, error) {
	var request AddRequest
	errors := make([]string, 0)

	if err := model.DecodeJSON(r.Body, &request.Product); err != nil {
		return nil, err
	}
	if request.Product == nil {
		return nil, fmt.Errorf("request body cannot be empty")
	}

	if request.Product.Name == "" {
		errors = append(errors, "Invalid Product Name")
	}

	if request.Product.Brand <= 0 {
		errors = append(errors, "Invalid Brand")
	}

	if request.Product.Category <= 0 {
		errors = append(errors, "Invalid Category")
	}

	if request.Product.SubCategory <= 0 {
		errors = append(errors, "Invalid SubCategory")
	}

	if !request.Product.Status.IsValid() {
		errors = append(errors, "Invalid Status")
	}

	if len(errors) > 0 {
		return nil, fmt.Errorf("%s", strings.Join(errors, ","))
	}
	return &request, nil
}
