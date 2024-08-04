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
				"must": []map[string]interface{}{
					{
						"match": map[string]interface{}{
							"keywords": req.Keywords,
						},
					},
					{
						"match": map[string]interface{}{
							"category": req.Category,
						},
					},
					{
						"match": map[string]interface{}{
							"sub_category": req.SubCategory,
						},
					},
					{
						"term": map[string]interface{}{
							"city_id": req.CityId,
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
		response.Products = append(response.Products, &hit.Source)
	}
	response.TotalResults = uint32(searchResponse.Hits.Total.Value)

	return &response, nil
}

func (es *ProductDataStore) AddProduct(product *searchpb.Product) error {
	body, err := json.Marshal(product)
	if err != nil {
		return err
	}

	req := esapi.IndexRequest{
		Index:      es.index,
		DocumentID: strconv.Itoa(int(product.Id)),
		Body:       bytes.NewReader(body),
		Refresh:    "true",
	}

	res, err := req.Do(context.Background(), es.client)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("error adding product: %s", res.String())
	}

	return nil
}

func (es *ProductDataStore) UpdateProduct(productId int, product *searchpb.Product) error {
	body, err := json.Marshal(product)
	if err != nil {
		return err
	}

	req := esapi.UpdateRequest{
		Index:      es.index,
		DocumentID: strconv.Itoa(productId),
		Body:       bytes.NewReader(body),
		Refresh:    "true",
	}

	res, err := req.Do(context.Background(), es.client)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("error updating product: %s", res.String())
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
