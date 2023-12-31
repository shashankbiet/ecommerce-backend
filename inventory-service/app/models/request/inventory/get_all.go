package inventory

import "net/http"

type GetAllRequest struct {
}

func (c GetAllRequest) Validate(r *http.Request) (interface{}, error) {
	return &GetAllRequest{}, nil
}
