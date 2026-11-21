package dao

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	model "search-service/app/models"
	searchpb "search-service/proto/search"
	"strconv"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

type ProductDataStore struct {
	client *elasticsearch.Client
	index  string
}

func NewProductDataStore(client *elasticsearch.Client, index string) IProductDataStore {
	return &ProductDataStore{
		client: client,
		index:  index,
	}
}

func (es *ProductDataStore) GetProductById(id string) (*searchpb.Product, error) {
	res, err := es.client.Get(es.index, id)
	if err != nil {
		return nil, err
	}
	defer func() { _ = res.Body.Close() }()

	if res.IsError() {
		return nil, fmt.Errorf("error getting product: %s", res.String())
	}

	var doc struct {
		Source json.RawMessage `json:"_source"`
	}
	if err := json.NewDecoder(res.Body).Decode(&doc); err != nil {
		return nil, err
	}
	if len(doc.Source) == 0 || string(doc.Source) == "null" {
		return nil, fmt.Errorf("product not found in response")
	}

	product, legacyBrand, legacyCategory, legacySubCategory, err := decodeIndexedProduct(doc.Source)
	if err != nil {
		return nil, err
	}
	return toProtoProduct(product, legacyBrand, legacyCategory, legacySubCategory), nil
}

func (es *ProductDataStore) FilterProduct(req *searchpb.ProductSearchRequest) (*searchpb.ProductSearchResponse, error) {
	var filters []map[string]interface{}
	if req.Category != "" {
		filters = append(filters, map[string]interface{}{
			"match": map[string]interface{}{
				"category": req.Category,
			},
		})
	}
	if req.SubCategory != "" {
		filters = append(filters, map[string]interface{}{
			"bool": map[string]interface{}{
				"should": []map[string]interface{}{
					{"match": map[string]interface{}{"sub_category": req.SubCategory}},
					{"match": map[string]interface{}{"subCategory": req.SubCategory}},
				},
				"minimum_should_match": 1,
			},
		})
	}

	boolQuery := map[string]interface{}{}
	if len(filters) > 0 {
		boolQuery["filter"] = filters
	}

	if req.Keywords != "" {
		boolQuery["must"] = []map[string]interface{}{
			{
				"wildcard": map[string]interface{}{
					"name": map[string]interface{}{
						"value":            "*" + req.Keywords + "*",
						"case_insensitive": true,
					},
				},
			},
		}
	} else {
		boolQuery["must"] = []map[string]interface{}{
			{
				"match_all": map[string]interface{}{},
			},
		}
	}

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": boolQuery,
		},
	}

	// Marshal the query to JSON
	body, err := json.Marshal(query)
	if err != nil {
		return nil, err
	}

	// Perform the search query
	res, err := es.client.Search(
		es.client.Search.WithContext(context.Background()),
		es.client.Search.WithIndex(es.index),
		es.client.Search.WithBody(bytes.NewReader(body)),
		es.client.Search.WithTrackTotalHits(true),
	)
	if err != nil {
		return nil, err
	}
	defer func() { _ = res.Body.Close() }()

	if res.IsError() {
		return nil, fmt.Errorf("error filtering products: %s", res.String())
	}

	// Parse the search response
	var searchResponse struct {
		Hits struct {
			Total struct {
				Value int `json:"value"`
			} `json:"total"`
			Hits []struct {
				Source json.RawMessage `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}
	if err := json.NewDecoder(res.Body).Decode(&searchResponse); err != nil {
		return nil, err
	}

	// Construct the ProductSearchResponse
	var response searchpb.ProductSearchResponse
	for _, hit := range searchResponse.Hits.Hits {
		if len(hit.Source) > 0 {
			product, legacyBrand, legacyCategory, legacySubCategory, err := decodeIndexedProduct(hit.Source)
			if err != nil {
				return nil, fmt.Errorf("decode indexed product: %w", err)
			}
			response.Products = append(response.Products, toProtoProduct(product, legacyBrand, legacyCategory, legacySubCategory))
		}
	}
	response.TotalResults = uint32(searchResponse.Hits.Total.Value)
	response.Keywords = req.Keywords
	response.Category = req.Category
	response.SubCategory = req.SubCategory

	return &response, nil
}

func (es *ProductDataStore) UpsertProduct(productId int64, product *model.Product) error {
	// Serialize the product as JSON
	body, err := json.Marshal(map[string]interface{}{
		"doc":           product, // Document to update
		"doc_as_upsert": true,    // Create if the document doesn't exist
	})
	if err != nil {
		return err
	}

	// Create the update request
	req := esapi.UpdateRequest{
		Index:      es.index,
		DocumentID: strconv.FormatInt(productId, 10), // Ensure this matches the document ID in Elasticsearch
		Body:       bytes.NewReader(body),
		Refresh:    "true", // Optionally set to "true" to immediately refresh the index
	}

	// Execute the update request
	res, err := req.Do(context.Background(), es.client)
	if err != nil {
		return err
	}
	defer func() { _ = res.Body.Close() }()

	// Check for errors in the response
	if res.IsError() {
		return fmt.Errorf("error upserting product: %s", res.String())
	}

	return nil
}

func (es *ProductDataStore) DeleteProduct(productId int64) error {
	req := esapi.DeleteRequest{
		Index:      es.index,
		DocumentID: strconv.FormatInt(productId, 10),
		Refresh:    "true",
	}

	res, err := req.Do(context.Background(), es.client)
	if err != nil {
		return err
	}
	defer func() { _ = res.Body.Close() }()

	if res.IsError() {
		return fmt.Errorf("error deleting product: %s", res.String())
	}

	return nil
}

func decodeIndexedProduct(data []byte) (*model.Product, string, string, string, error) {
	normalized, err := model.NormalizeJSON(data)
	if err != nil {
		return nil, "", "", "", err
	}
	var document struct {
		Id          int64               `json:"id"`
		Name        string              `json:"name"`
		Description string              `json:"description"`
		Slug        string              `json:"slug"`
		Brand       json.RawMessage     `json:"brand"`
		Category    json.RawMessage     `json:"category"`
		SubCategory json.RawMessage     `json:"sub_category"`
		ImageUrl    string              `json:"image_url"`
		ImageId     string              `json:"image_id"`
		Status      model.ProductStatus `json:"status"`
		CreatedAt   time.Time           `json:"created_at"`
		UpdatedAt   time.Time           `json:"updated_at"`
		DeletedAt   time.Time           `json:"deleted_at"`
	}
	if err := json.Unmarshal(normalized, &document); err != nil {
		return nil, "", "", "", err
	}

	brand, legacyBrand := decodeReference(document.Brand)
	category, legacyCategory := decodeReference(document.Category)
	subCategory, legacySubCategory := decodeReference(document.SubCategory)
	imageURL := document.ImageUrl
	if imageURL == "" {
		imageURL = document.ImageId
	}

	return &model.Product{
		Id:          document.Id,
		Name:        document.Name,
		Description: document.Description,
		Slug:        document.Slug,
		Brand:       int16(brand),
		Category:    int16(category),
		SubCategory: int16(subCategory),
		ImageUrl:    imageURL,
		Status:      document.Status,
		CreatedAt:   document.CreatedAt,
		UpdatedAt:   document.UpdatedAt,
		DeletedAt:   document.DeletedAt,
	}, legacyBrand, legacyCategory, legacySubCategory, nil
}

func decodeReference(data json.RawMessage) (int64, string) {
	if len(data) == 0 || string(data) == "null" {
		return 0, ""
	}
	var numeric int64
	if err := json.Unmarshal(data, &numeric); err == nil {
		return numeric, strconv.FormatInt(numeric, 10)
	}
	var text string
	if err := json.Unmarshal(data, &text); err != nil {
		return 0, ""
	}
	numeric, err := strconv.ParseInt(text, 10, 64)
	if err == nil {
		return numeric, text
	}
	return 0, text
}

func toProtoProduct(product *model.Product, legacyBrand string, legacyCategory string, legacySubCategory string) *searchpb.Product {
	result := &searchpb.Product{
		Id:            product.Id,
		Name:          product.Name,
		Description:   product.Description,
		Brand:         legacyBrand,
		Category:      legacyCategory,
		SubCategory:   legacySubCategory,
		Slug:          product.Slug,
		BrandId:       int32(product.Brand),
		CategoryId:    int32(product.Category),
		SubCategoryId: int32(product.SubCategory),
		ImageUrl:      product.ImageUrl,
		Status:        int64(product.Status),
		CreatedAt:     formatTime(product.CreatedAt),
		UpdatedAt:     formatTime(product.UpdatedAt),
		DeletedAt:     formatTime(product.DeletedAt),
	}
	return result
}

func formatTime(value time.Time) string {
	if value.IsZero() {
		return ""
	}
	return value.UTC().Format(time.RFC3339)
}
