package model

import jwt "github.com/dgrijalva/jwt-go"

// Token ...
type Token struct {
	//UserID uint
	ID       string
	Username string
	IsSeller bool
	//Email  string
	*jwt.StandardClaims
}
