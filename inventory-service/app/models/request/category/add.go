package category

import (
	"encoding/json"
	"fmt"
	model "inventory-service/app/models"
	"net/http"
	"strings"
)

type AddRequest struct {
	Category *model.Category
}

func (c AddRequest) Validate(r *http.Request) (interface{}, error) {
	var request AddRequest
	errors := make([]string, 0)

	if err := json.NewDecoder(r.Body).Decode(&request.Category); err != nil {
		return nil, err
	}

	if request.Category.Name == "" {
		errors = append(errors, "Invalid Category Name")
	}

	if len(errors) > 0 {
		return nil, fmt.Errorf(strings.Join(errors, ","))
	}
	return &request, nil
}
