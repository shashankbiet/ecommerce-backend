package inventory

import (
	"encoding/json"
	"fmt"
	model "inventory-service/app/models"
	"net/http"
	"strings"
)

type UpdateRequest struct {
	Product *model.Product
}

func (c UpdateRequest) Validate(r *http.Request) (interface{}, error) {
	var request UpdateRequest
	errors := make([]string, 0)

	if err := json.NewDecoder(r.Body).Decode(&request.Product); err != nil {
		return nil, err
	}

	if request.Product.Id <= 0 {
		errors = append(errors, "Invalid Product Id")
	}

	if request.Product.Name == "" {
		errors = append(errors, "Invalid Product Name")
	}

	if request.Product.Brand == "" {
		errors = append(errors, "Invalid Brand")
	}

	if request.Product.Category == "" {
		errors = append(errors, "Invalid Category")
	}

	if request.Product.SubCategory == "" {
		errors = append(errors, "Invalid SubCategory")
	}

	if request.Product.Weight <= 0 {
		errors = append(errors, "Invalid SubCategory")
	}

	if len(errors) > 0 {
		return nil, fmt.Errorf(strings.Join(errors, ","))
	}
	return &request, nil
}
