package models

import "time"

//Product product
type Product struct {
	ID      int       `json:"id"`
	Name    string    `json:"name"`
	Sku     string    `json:"sku"`
	Qty     int       `json:"qty"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}
