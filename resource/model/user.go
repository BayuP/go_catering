package model

//User models
type User struct {
	Base
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Address  string `json:"address"`
	IsSeller bool   `json:"is_seller"`
}
