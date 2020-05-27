package resmodel

//LoginRes ...
type LoginRes struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	IsSeller bool   `'json:"is_seller"`
	Token    string `'json:"token"`
}

//AllProduct ...
type AllProduct struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Price        int    `json:"price"`
	Stock        int    `json:"stock"`
	StoreName    string `json:"store_name"`
	StoreID      string `json:"store_id"`
	DailyProduct bool   `json:"daily_product"`
}
