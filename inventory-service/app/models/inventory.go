package model

import "time"

//easyjson:json
type Inventory struct {
	ProductId     int64     `json:"productId"`
	SKU           int32     `json:"sku"`
	PurchasePrice int32     `json:"purchasePrice"`
	SalePrice     int32     `json:"salePrice"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}
