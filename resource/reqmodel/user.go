package reqmodel

//CreateUserReq ...
type CreateUserReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Address  string `json:"address"`
}

//LoginReq ...
type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	IsSeller bool   `json:"is_seller"`
}

//TransactionReq ..
type TransactionReq struct {
	StoreID     string `json:"store_id"`
	ProductID   string `json:"product_id"`
	Subsription bool   `json:"subs"`
	Price       int    `json:"price"`
	Quantity    int    `json:"quantity"`
}

//UpdateTrxReq ..
type UpdateTrxReq struct {
	TransactionCode string `json:"trx_code"`
	StatusTrx       int    `json:"status_trx"`
}
