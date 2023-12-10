package intf

import "net/http"

type ValidateRequest interface {
	Validate(r *http.Request) (interface{}, error)
}
