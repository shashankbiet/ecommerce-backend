package producthandler

import (
	"strings"
	"testing"

	searchpb "search-service/proto/search"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestValidateGetProductRequestOptionalCategoryFilters(t *testing.T) {
	tests := []struct {
		name            string
		category        string
		subCategory     string
		wantCode        codes.Code
		wantCategory    string
		wantSubCategory string
	}{
		{
			name: "both omitted",
		},
		{
			name:            "provided values are trimmed",
			category:        " Electronics ",
			subCategory:     " Accessories ",
			wantCategory:    "Electronics",
			wantSubCategory: "Accessories",
		},
		{
			name:     "whitespace category is invalid",
			category: "  ",
			wantCode: codes.InvalidArgument,
		},
		{
			name:        "whitespace subcategory is invalid",
			subCategory: "  ",
			wantCode:    codes.InvalidArgument,
		},
		{
			name:     "category over max length is invalid",
			category: strings.Repeat("a", maxCategoryNameLength+1),
			wantCode: codes.InvalidArgument,
		},
		{
			name:        "subcategory over max length is invalid",
			subCategory: strings.Repeat("a", maxCategoryNameLength+1),
			wantCode:    codes.InvalidArgument,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := &searchpb.ProductSearchRequest{
				Category:    tt.category,
				SubCategory: tt.subCategory,
			}
			err := validateGetProductRequest(req)
			if tt.wantCode == codes.OK {
				if err != nil {
					t.Fatalf("validateGetProductRequest() error = %v, want nil", err)
				}
				if req.GetCategory() != tt.wantCategory {
					t.Errorf("category = %q, want %q", req.GetCategory(), tt.wantCategory)
				}
				if req.GetSubCategory() != tt.wantSubCategory {
					t.Errorf("sub_category = %q, want %q", req.GetSubCategory(), tt.wantSubCategory)
				}
				return
			}

			if status.Code(err) != tt.wantCode {
				t.Errorf("validateGetProductRequest() code = %v, want %v", status.Code(err), tt.wantCode)
			}
		})
	}
}
