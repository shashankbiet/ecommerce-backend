package model

import "time"

type Brand struct {
	Id        int16     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}
