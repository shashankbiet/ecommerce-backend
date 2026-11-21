package keyword

import (
	"fmt"
	model "inventory-service/app/models"
	"net/http"
	"strings"
)

type AddRequest struct {
	Keyword *model.Keyword
}

func (AddRequest) Validate(r *http.Request) (interface{}, error) {
	var request AddRequest
	if err := model.DecodeJSON(r.Body, &request.Keyword); err != nil {
		return nil, err
	}
	if request.Keyword == nil {
		return nil, fmt.Errorf("request body cannot be empty")
	}
	if strings.TrimSpace(request.Keyword.Name) == "" {
		return nil, fmt.Errorf("Invalid Keyword Name")
	}
	return &request, nil
}
