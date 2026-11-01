package dao

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

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
