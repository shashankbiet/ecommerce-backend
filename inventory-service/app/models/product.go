package model

import "time"

//easyjson:json
type Product struct {
	Id          int64         `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Slug        string        `json:"slug"`
	Brand       int16         `json:"brand"`
	Category    int16         `json:"category"`
	SubCategory int16         `json:"sub_category"`
	ImageUrl    string        `json:"image_url"`
	Status      ProductStatus `json:"status"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
	DeletedAt   time.Time     `json:"deleted_at"`
}
