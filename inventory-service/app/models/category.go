package model

import "time"

//easyjson:json
type Category struct {
	Id        int16     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}
