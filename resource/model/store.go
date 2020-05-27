package model

//Store models
type Store struct {
	Base
	ID           string `json:"id"`
	UserID       string `json:"user_id"`
	StoreName    string `json:"store_name"`
	StoreAddress string `json:"store_address"`
	SellingArea  string `json:"selling_area"`
}
