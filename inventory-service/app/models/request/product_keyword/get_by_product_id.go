package productkeyword

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type GetByProductIdRequest struct {
	ProductId int64
}

func (GetByProductIdRequest) Validate(r *http.Request) (interface{}, error) {
	value := mux.Vars(r)["product_id"]
	if value == "" {
		value = mux.Vars(r)["productId"]
	}
	id, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return nil, err
	}
	if id <= 0 {
		return nil, fmt.Errorf("Invalid ProductId")
	}
	return &GetByProductIdRequest{ProductId: id}, nil
}
