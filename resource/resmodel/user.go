package resmodel

//LoginRes ...
type LoginRes struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	IsSeller bool   `'json:"is_seller"`
	Token    string `'json:"token"`
}
