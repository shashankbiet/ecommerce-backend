package productkeyword

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type GetByKeywordIdRequest struct {
	KeywordId int64
}

func (GetByKeywordIdRequest) Validate(r *http.Request) (interface{}, error) {
	value := mux.Vars(r)["keyword_id"]
	if value == "" {
		value = mux.Vars(r)["keywordId"]
	}
	id, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return nil, err
	}
	if id <= 0 {
		return nil, fmt.Errorf("Invalid KeywordId")
	}
	return &GetByKeywordIdRequest{KeywordId: id}, nil
}
