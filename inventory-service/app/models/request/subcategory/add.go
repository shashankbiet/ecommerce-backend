package subcategory

import (
	"fmt"
	model "inventory-service/app/models"
	"net/http"
	"strings"
)

type AddRequest struct {
	SubCategory *model.SubCategory
}

func (c AddRequest) Validate(r *http.Request) (interface{}, error) {
	var request AddRequest
	errors := make([]string, 0)

	if err := model.DecodeJSON(r.Body, &request.SubCategory); err != nil {
		return nil, err
	}
	if request.SubCategory == nil {
		return nil, fmt.Errorf("request body cannot be empty")
	}

	if request.SubCategory.Name == "" {
		errors = append(errors, "Invalid SubCategory Name")
	}

	if request.SubCategory.Category <= 0 {
		errors = append(errors, "Invalid Category Name")
	}

	if len(errors) > 0 {
		return nil, fmt.Errorf("%s", strings.Join(errors, ","))
	}
	return &request, nil
}
