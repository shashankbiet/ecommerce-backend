package model

import "time"

//easyjson:json
type Inventory struct {
	ProductId     int64     `json:"product_id"`
	SKU           int32     `json:"sku"`
	MRP           int32     `json:"mrp"`
	PurchasePrice int32     `json:"purchase_price"`
	SalePrice     int32     `json:"sale_price"`
	Quantity      int32     `json:"quantity"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
