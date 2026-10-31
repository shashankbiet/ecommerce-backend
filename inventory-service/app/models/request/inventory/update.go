package inventory

import (
	"encoding/json"
	"fmt"
	model "inventory-service/app/models"
	"net/http"
	"strings"
)

type UpdateRequest struct {
	Inventory *model.Inventory
}

func (c UpdateRequest) Validate(r *http.Request) (interface{}, error) {
	var request UpdateRequest
	errors := make([]string, 0)

	if err := json.NewDecoder(r.Body).Decode(&request.Inventory); err != nil {
		return nil, err
	}
	if request.Inventory == nil {
		return nil, fmt.Errorf("request body cannot be empty")
	}

	if request.Inventory.ProductId <= 0 {
		errors = append(errors, "Invalid ProductId")
	}

	if request.Inventory.SKU < 0 {
		errors = append(errors, "Invalid SKU")
	}

	if request.Inventory.PurchasePrice < 0 {
		errors = append(errors, "Invalid PurchasePrice")
	}

	if request.Inventory.SalePrice < 0 {
		errors = append(errors, "Invalid SalePrice")
	}

	if request.Inventory.Quantity < 0 {
		errors = append(errors, "Invalid Quantity")
	}

	if len(errors) > 0 {
		return nil, fmt.Errorf("%s", strings.Join(errors, ","))
	}
	return &request, nil
}
