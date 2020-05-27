package resmodel

// StoreRes ...
type StoreRes struct {
	ID           string `json:"id"`
	StoreName    string `json:"store_name"`
	StoreAddress string `json:"store_address"`
	SellingArea  string `json:"selling_area"`
}

// ProductByStoreRes ...
type ProductByStoreRes struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Price        int    `json:"price"`
	Stock        int    `json:"stock"`
	DailyProduct bool   `json:"daily_product"`
}
