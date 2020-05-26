package model

import jwt "github.com/dgrijalva/jwt-go"

// Token ...
type Token struct {
	//UserID uint
	ID       string
	Username string
	//Email  string
	*jwt.StandardClaims
}
