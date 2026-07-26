package auth

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	CustCode  string `json:"custCode"`
	AccountID string `json:"accountId"`

	jwt.RegisteredClaims
}
