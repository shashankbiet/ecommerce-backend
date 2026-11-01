package producthandler

import (
	searchpb "search-service/proto/search"
	"strings"
	"unicode/utf8"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const maxCategoryNameLength = 255

func validateGetProductRequest(req *searchpb.ProductSearchRequest) error {
	if req == nil {
		return status.Error(codes.InvalidArgument, "request cannot be nil")
	}

	for _, field := range []struct {
		name  string
		value *string
	}{
		{name: "category", value: &req.Category},
		{name: "sub_category", value: &req.SubCategory},
	} {
		if *field.value == "" {
			continue
		}

		trimmed := strings.TrimSpace(*field.value)
		if trimmed == "" {
			return status.Errorf(codes.InvalidArgument, "%s must not be blank", field.name)
		}
		if utf8.RuneCountInString(trimmed) > maxCategoryNameLength {
			return status.Errorf(codes.InvalidArgument, "%s must not exceed %d characters", field.name, maxCategoryNameLength)
		}
		*field.value = trimmed
	}

	return nil
}
