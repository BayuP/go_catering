package reqmodel

//ReqCreateStore ...
type ReqCreateStore struct {
	StoreName    string `json:"store_name"`
	StoreAddress string `json:"store_address"`
	SellingArea  string `json:"selling_area"`
}

//ReqCreateProduct ...
type ReqCreateProduct struct {
	Name         string `json:"name"`
	StoreID      string `json:"store_id"`
	Price        int    `json:"price"`
	Stock        int    `json:"stock"`
	DailyProduct bool   `json:"daily_product"`
}

//UpdateDeliverStatus ...
type UpdateDeliverStatus struct {
	ID            string `json:"id"`
	DeliverStatus int    `json:"deliver_status"`
	DailyProduct  bool   `json:"daily_product"`
}
