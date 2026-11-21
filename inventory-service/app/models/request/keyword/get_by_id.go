package keyword

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type GetByIdRequest struct {
	Id int64
}

func (GetByIdRequest) Validate(r *http.Request) (interface{}, error) {
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		return nil, err
	}
	if id <= 0 {
		return nil, fmt.Errorf("Invalid Id")
	}
	return &GetByIdRequest{Id: id}, nil
}
