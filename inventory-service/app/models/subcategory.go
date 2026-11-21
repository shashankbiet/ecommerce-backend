package model

import "time"

//easyjson:json
type SubCategory struct {
	Id        int16     `json:"id"`
	Name      string    `json:"name"`
	Category  int16     `json:"category"`
	CreatedAt time.Time `json:"created_at"`
}
