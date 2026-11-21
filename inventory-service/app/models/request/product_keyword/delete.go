package productkeyword

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type DeleteRequest struct {
	ProductId int64
	KeywordId int64
}

func (DeleteRequest) Validate(r *http.Request) (interface{}, error) {
	productValue := mux.Vars(r)["product_id"]
	if productValue == "" {
		productValue = mux.Vars(r)["productId"]
	}
	keywordValue := mux.Vars(r)["keyword_id"]
	if keywordValue == "" {
		keywordValue = mux.Vars(r)["keywordId"]
	}
	productId, productErr := strconv.ParseInt(productValue, 10, 64)
	keywordId, keywordErr := strconv.ParseInt(keywordValue, 10, 64)
	errors := make([]string, 0)
	if productErr != nil || productId <= 0 {
		errors = append(errors, "Invalid ProductId")
	}
	if keywordErr != nil || keywordId <= 0 {
		errors = append(errors, "Invalid KeywordId")
	}
	if len(errors) > 0 {
		return nil, fmt.Errorf("%s", strings.Join(errors, ","))
	}
	return &DeleteRequest{ProductId: productId, KeywordId: keywordId}, nil
}
