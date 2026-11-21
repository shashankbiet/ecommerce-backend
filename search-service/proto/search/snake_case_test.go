package search

import (
	"encoding/json"
	"testing"

	"google.golang.org/protobuf/encoding/protojson"
)

func TestProductProtoJSONUsesSnakeCaseFields(t *testing.T) {
	data, err := protojson.Marshal(&Product{
		SubCategory:   "Accessories",
		SubCategoryId: 4,
		ImageUrl:      "phone.png",
		CreatedAt:     "2026-11-22T01:02:03Z",
	})
	if err != nil {
		t.Fatalf("marshal product: %v", err)
	}

	var fields map[string]interface{}
	if err := json.Unmarshal(data, &fields); err != nil {
		t.Fatalf("decode product JSON: %v", err)
	}
	for _, field := range []string{"sub_category", "sub_category_id", "image_url", "created_at"} {
		if _, ok := fields[field]; !ok {
			t.Errorf("protobuf JSON missing %q: %s", field, data)
		}
	}
	for _, field := range []string{"subCategory", "subCategoryId", "imageUrl", "createdAt"} {
		if _, ok := fields[field]; ok {
			t.Errorf("protobuf JSON contains camelCase field %q: %s", field, data)
		}
	}
}
