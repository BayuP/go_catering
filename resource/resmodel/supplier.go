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

// DelivProduct ...
type DelivProduct struct {
	ID              string `json:"id"`
	ProductID       string `json:"product_id"`
	CustomerID      string `json:"customer_id"`
	CustomerAddress string `json:"cust_address"`
	DailyProduct    bool   `json:"daily_product"`
}
