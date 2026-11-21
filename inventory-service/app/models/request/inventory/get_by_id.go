package inventory

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

const (
	ID       = "product_id"
	LegacyID = "productId"
)

type GetByIdRequest struct {
	ProductId int64
}

func (c GetByIdRequest) Validate(r *http.Request) (interface{}, error) {
	var request GetByIdRequest
	errors := make([]string, 0)
	value := mux.Vars(r)[ID]
	if value == "" {
		value = mux.Vars(r)[LegacyID]
	}
	id, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return request, err
	}
	request.ProductId = id

	if request.ProductId <= 0 {
		errors = append(errors, "Invalid ProductId")
	}

	if len(errors) > 0 {
		return nil, fmt.Errorf("%s", strings.Join(errors, ","))
	}
	return &request, nil
}
