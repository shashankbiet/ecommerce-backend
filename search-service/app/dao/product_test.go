package dao

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	model "search-service/app/models"
	"strings"
	"testing"
	"time"

	searchpb "search-service/proto/search"

	"github.com/elastic/go-elasticsearch/v8"
)

func TestFilterProductKeywordWildcardIsCaseInsensitive(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var request map[string]interface{}
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			t.Errorf("decode search request: %v", err)
			http.Error(w, "invalid request", http.StatusBadRequest)
			return
		}

		query := request["query"].(map[string]interface{})
		boolQuery := query["bool"].(map[string]interface{})
		must := boolQuery["must"].([]interface{})
		wildcard := must[0].(map[string]interface{})["wildcard"].(map[string]interface{})
		nameQuery := wildcard["name"].(map[string]interface{})
		if got := nameQuery["value"]; got != "*Accessories*" {
			t.Errorf("wildcard value = %v, want *Accessories*", got)
		}
		if got := nameQuery["case_insensitive"]; got != true {
			t.Errorf("case_insensitive = %v, want true", got)
		}

		w.Header().Set("X-Elastic-Product", "Elasticsearch")
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"hits":{"total":{"value":1},"hits":[{"_source":{"id":4,"name":"Mobile Phone Accessories","category":"Electronics","subCategory":"Accessories"}}]}}`))
	}))
	defer server.Close()

	client, err := elasticsearch.NewClient(elasticsearch.Config{Addresses: []string{server.URL}})
	if err != nil {
		t.Fatalf("create Elasticsearch client: %v", err)
	}

	response, err := NewProductDataStore(client, "products").FilterProduct(&searchpb.ProductSearchRequest{
		Keywords:    "Accessories",
		Category:    "Electronics",
		SubCategory: "Accessories",
	})
	if err != nil {
		t.Fatalf("FilterProduct() error = %v", err)
	}
	if response.GetTotalResults() != 1 || len(response.GetProducts()) != 1 {
		t.Fatalf("FilterProduct() returned %d results; want 1", response.GetTotalResults())
	}
	if !strings.EqualFold(response.GetProducts()[0].GetName(), "Mobile Phone Accessories") {
		t.Errorf("product name = %q, want Mobile Phone Accessories", response.GetProducts()[0].GetName())
	}
	if response.GetProducts()[0].GetCategory() != "Electronics" ||
		response.GetProducts()[0].GetSubCategory() != "Accessories" {
		t.Errorf("legacy category fields were not preserved: %#v", response.GetProducts()[0])
	}
}

func TestFilterProductOmitsEmptyCategoryFilters(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var request map[string]interface{}
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			t.Errorf("decode search request: %v", err)
			http.Error(w, "invalid request", http.StatusBadRequest)
			return
		}

		query := request["query"].(map[string]interface{})
		boolQuery := query["bool"].(map[string]interface{})
		if _, ok := boolQuery["filter"]; ok {
			t.Errorf("search query contains filters for omitted category fields: %v", boolQuery["filter"])
		}

		w.Header().Set("X-Elastic-Product", "Elasticsearch")
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"hits":{"total":{"value":0},"hits":[]}}`))
	}))
	defer server.Close()

	client, err := elasticsearch.NewClient(elasticsearch.Config{Addresses: []string{server.URL}})
	if err != nil {
		t.Fatalf("create Elasticsearch client: %v", err)
	}

	if _, err := NewProductDataStore(client, "products").FilterProduct(&searchpb.ProductSearchRequest{}); err != nil {
		t.Fatalf("FilterProduct() with omitted category fields: %v", err)
	}
}

func TestUpsertProductIndexesUpdatedProductModel(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/products/_update/9007199254740993" {
			t.Errorf("update path = %q, want int64 product id path", r.URL.Path)
		}
		if got := r.URL.Query().Get("refresh"); got != "true" {
			t.Errorf("refresh = %q, want true", got)
		}
		var body struct {
			Document      json.RawMessage `json:"doc"`
			DocumentIsNew bool            `json:"doc_as_upsert"`
		}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatalf("decode update body: %v", err)
		}
		if !body.DocumentIsNew {
			t.Error("doc_as_upsert = false, want true")
		}
		var fields map[string]interface{}
		if err := json.Unmarshal(body.Document, &fields); err != nil {
			t.Fatalf("decode indexed fields: %v", err)
		}
		for _, field := range []string{"sub_category", "image_url", "created_at", "updated_at", "deleted_at"} {
			if _, ok := fields[field]; !ok {
				t.Errorf("indexed document missing %q: %s", field, body.Document)
			}
		}
		var document model.Product
		if err := json.Unmarshal(body.Document, &document); err != nil {
			t.Fatalf("decode indexed product: %v", err)
		}
		if document.Slug != "phone" ||
			document.Brand != 2 ||
			document.Category != 3 ||
			document.SubCategory != 4 ||
			document.ImageUrl != "phone.png" ||
			document.Status != 1 {
			t.Errorf("indexed product = %#v", document)
		}
		w.Header().Set("X-Elastic-Product", "Elasticsearch")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"result":"updated"}`))
	}))
	defer server.Close()

	client, err := elasticsearch.NewClient(elasticsearch.Config{Addresses: []string{server.URL}})
	if err != nil {
		t.Fatalf("create Elasticsearch client: %v", err)
	}
	product := &model.Product{
		Id:          9007199254740993,
		Slug:        "phone",
		Brand:       2,
		Category:    3,
		SubCategory: 4,
		ImageUrl:    "phone.png",
		Status:      1,
	}
	if err := NewProductDataStore(client, "products").UpsertProduct(product.Id, product); err != nil {
		t.Fatalf("UpsertProduct() error = %v", err)
	}
}

func TestDecodeIndexedProductSupportsCurrentAndLegacyDocuments(t *testing.T) {
	createdAt := "2026-11-22T01:02:03Z"
	tests := []struct {
		name                  string
		document              string
		wantBrand             int16
		wantCategory          int16
		wantSubCategory       int16
		wantLegacyBrand       string
		wantLegacyCategory    string
		wantLegacySubCategory string
		wantImageURL          string
	}{
		{
			name: "current numeric references",
			document: `{
				"id":12,"name":"Phone","description":"Smart phone","slug":"phone",
				"brand":2,"category":3,"sub_category":4,"image_url":"current.png",
				"status":1,"created_at":"2026-11-22T01:02:03Z",
				"updated_at":"2026-11-22T02:02:03Z","deleted_at":"0001-01-01T00:00:00Z"
			}`,
			wantBrand: 2, wantCategory: 3, wantSubCategory: 4,
			wantLegacyBrand: "2", wantLegacyCategory: "3", wantLegacySubCategory: "4",
			wantImageURL: "current.png",
		},
		{
			name: "legacy string references",
			document: `{
				"id":13,"name":"Accessories","brand":"Acme","category":"Electronics",
				"subCategory":"Accessories","imageId":"legacy.png",
				"createdAt":"2026-11-22T01:02:03Z","updatedAt":"2026-11-22T01:02:03Z"
			}`,
			wantLegacyBrand: "Acme", wantLegacyCategory: "Electronics",
			wantLegacySubCategory: "Accessories", wantImageURL: "legacy.png",
		},
		{
			name: "numeric string references",
			document: `{
				"id":14,"name":"Tablet","brand":"2","category":"3","subCategory":"4",
				"createdAt":"2026-11-22T01:02:03Z","updatedAt":"2026-11-22T01:02:03Z"
			}`,
			wantBrand: 2, wantCategory: 3, wantSubCategory: 4,
			wantLegacyBrand: "2", wantLegacyCategory: "3", wantLegacySubCategory: "4",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			product, brand, category, subCategory, err := decodeIndexedProduct([]byte(test.document))
			if err != nil {
				t.Fatalf("decodeIndexedProduct() error = %v", err)
			}
			if product.Brand != test.wantBrand || product.Category != test.wantCategory ||
				product.SubCategory != test.wantSubCategory {
				t.Errorf("decoded references = %d/%d/%d, want %d/%d/%d",
					product.Brand, product.Category, product.SubCategory,
					test.wantBrand, test.wantCategory, test.wantSubCategory)
			}
			if brand != test.wantLegacyBrand || category != test.wantLegacyCategory ||
				subCategory != test.wantLegacySubCategory {
				t.Errorf("legacy references = %q/%q/%q", brand, category, subCategory)
			}
			if product.ImageUrl != test.wantImageURL {
				t.Errorf("image URL = %q, want %q", product.ImageUrl, test.wantImageURL)
			}
			if product.CreatedAt.Format(time.RFC3339) != createdAt {
				t.Errorf("createdAt = %s, want %s", product.CreatedAt, createdAt)
			}
		})
	}
}

func TestGetProductByIDConvertsCurrentDocumentToProto(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Elastic-Product", "Elasticsearch")
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"_source":{
			"id":9007199254740993,"name":"Phone","description":"Smart phone",
			"slug":"phone","brand":2,"category":3,"sub_category":4,
			"image_url":"phone.png","status":1,
			"created_at":"2026-11-22T01:02:03Z","updated_at":"2026-11-22T02:02:03Z"
		}}`))
	}))
	defer server.Close()

	client, err := elasticsearch.NewClient(elasticsearch.Config{Addresses: []string{server.URL}})
	if err != nil {
		t.Fatalf("create Elasticsearch client: %v", err)
	}

	product, err := NewProductDataStore(client, "products").GetProductById("9007199254740993")
	if err != nil {
		t.Fatalf("GetProductById() error = %v", err)
	}
	if product.Id != 9007199254740993 || product.BrandId != 2 ||
		product.CategoryId != 3 || product.SubCategoryId != 4 ||
		product.Slug != "phone" || product.ImageUrl != "phone.png" ||
		product.Status != 1 || product.CreatedAt != "2026-11-22T01:02:03Z" {
		t.Fatalf("converted product = %#v", product)
	}
	if product.DeletedAt != "" {
		t.Errorf("DeletedAt = %q, want empty", product.DeletedAt)
	}
}
