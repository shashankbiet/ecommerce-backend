package keyword

import "net/http"

type GetAllRequest struct{}

func (GetAllRequest) Validate(r *http.Request) (interface{}, error) {
	return &GetAllRequest{}, nil
}
