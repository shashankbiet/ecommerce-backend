package request

import "net/http"

type IValidateRequest interface {
	Validate(r *http.Request) (interface{}, error)
}
