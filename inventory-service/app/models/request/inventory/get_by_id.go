package inventory

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

const (
	ID = "productId"
)

type GetByIdRequest struct {
	ProductId int64
}

func (c GetByIdRequest) Validate(r *http.Request) (interface{}, error) {
	var request GetByIdRequest
	errors := make([]string, 0)
	id, err := strconv.ParseInt(mux.Vars(r)[ID], 10, 64)
	if err != nil {
		return request, err
	}
	request.ProductId = id

	if request.ProductId <= 0 {
		errors = append(errors, "Invalid ProductId")
	}

	if len(errors) > 0 {
		return nil, fmt.Errorf(strings.Join(errors, ","))
	}
	return &request, nil
}
