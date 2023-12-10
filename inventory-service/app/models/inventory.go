package model

import "time"

//easyjson:json
type Inventory struct {
	ProductId int
	SKU       int
	Price     int
	SalePrice int
	CreatedAt time.Time
	UpdatedAt time.Time
}
