package model

import "time"

//easyjson:json
type SubCategory struct {
	Id        int16     `json:"id"`
	Name      string    `json:"name"`
	Category  string    `json:"category"`
	CreatedAt time.Time `json:"createdAt"`
}
