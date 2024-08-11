package dao

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	searchpb "search-service/proto/search"
	"strconv"

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
	defer res.Body.Close()

	if res.IsError() {
		return nil, fmt.Errorf("error getting product: %s", res.String())
	}

	var product searchpb.Product
	if err := json.NewDecoder(res.Body).Decode(&product); err != nil {
		return nil, err
	}

	return &product, nil
}

func (es *ProductDataStore) FilterProduct(req *searchpb.ProductSearchRequest) (*searchpb.ProductSearchResponse, error) {
	// Build the Elasticsearch query based on the ProductSearchRequest fields
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"filter": []map[string]interface{}{
					{
						"match": map[string]interface{}{
							"category": req.Category,
						},
					},
					{
						"match": map[string]interface{}{
							"subCategory": req.SubCategory,
						},
					},
				},
				"must": []map[string]interface{}{
					{
						"wildcard": map[string]interface{}{
							"name": "*" + req.Keywords + "*",
						},
					},
				},
			},
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
	defer res.Body.Close()

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
				Source searchpb.Product `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}
	if err := json.NewDecoder(res.Body).Decode(&searchResponse); err != nil {
		return nil, err
	}

	// Construct the ProductSearchResponse
	var response searchpb.ProductSearchResponse
	for _, hit := range searchResponse.Hits.Hits {
		product := hit.Source
		response.Products = append(response.Products, &product)
	}
	response.TotalResults = uint32(searchResponse.Hits.Total.Value)
	response.Keywords = req.Keywords
	response.Category = req.Category
	response.SubCategory = req.SubCategory

	return &response, nil
}

func (es *ProductDataStore) UpsertProduct(productId int, product *searchpb.Product) error {
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
		DocumentID: strconv.Itoa(productId), // Ensure this matches the document ID in Elasticsearch
		Body:       bytes.NewReader(body),
		Refresh:    "true", // Optionally set to "true" to immediately refresh the index
	}

	// Execute the update request
	res, err := req.Do(context.Background(), es.client)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// Check for errors in the response
	if res.IsError() {
		return fmt.Errorf("error upserting product: %s", res.String())
	}

	return nil
}

func (es *ProductDataStore) DeleteProduct(productId int) error {
	req := esapi.DeleteRequest{
		Index:      es.index,
		DocumentID: strconv.Itoa(productId),
		Refresh:    "true",
	}

	res, err := req.Do(context.Background(), es.client)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("error deleting product: %s", res.String())
	}

	return nil
}
