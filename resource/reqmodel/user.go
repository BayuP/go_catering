package reqmodel

//CreateUserReq ...
type CreateUserReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Address  string `json:"address"`
}
