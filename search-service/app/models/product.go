package model

import "time"

//easyjson:json
type Product struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Brand       string    `json:"brand"`
	Category    string    `json:"category"`
	SubCategory string    `json:"subCategory"`
	ImageId     string    `json:"imageId"`
	Weight      float32   `json:"weight"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
