package subcategory

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

const (
	ID = "id"
)

type GetByIdRequest struct {
	Id int16
}

func (c GetByIdRequest) Validate(r *http.Request) (interface{}, error) {
	var request GetByIdRequest
	errors := make([]string, 0)
	id, err := strconv.ParseInt(mux.Vars(r)[ID], 10, 16)
	if err != nil {
		return request, err
	}
	request.Id = int16(id)

	if request.Id <= 0 {
		errors = append(errors, "Invalid Id")
	}

	if len(errors) > 0 {
		return nil, fmt.Errorf(strings.Join(errors, ","))
	}
	return &request, nil
}
