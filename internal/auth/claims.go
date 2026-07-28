package auth

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	AccountNo string `json:"accountNo"`
	CustCode  string `json:"custCode"`

	jwt.RegisteredClaims
}
