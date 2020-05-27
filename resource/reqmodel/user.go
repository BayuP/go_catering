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
