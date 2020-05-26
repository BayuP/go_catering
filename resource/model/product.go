package model

//Product ...
type Product struct {
	Base
	ID           string `json:"id"`
	Name         string `json:"name"`
	StoreID      string `json:"store_id"`
	Price        int    `json:"price"`
	Stock        int    `json:"stock"`
	DailyProduct bool   `json:"daily_product"`
}
