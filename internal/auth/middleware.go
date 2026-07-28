package auth

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Middleware(cookieName string, jwtSecret string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString := c.Cookies(cookieName)
		if tokenString == "" {
			return ErrMissingAuthCookie
		}

		claims := new(Claims)

		token, err := jwt.ParseWithClaims(
			tokenString,
			claims,
			func(token *jwt.Token) (any, error) {
				if token.Method != jwt.SigningMethodHS256 {
					return nil, fmt.Errorf(
						"unexpected signing method: %s",
						token.Method.Alg(),
					)
				}

				return []byte(jwtSecret), nil
			},
			jwt.WithValidMethods([]string{
				jwt.SigningMethodHS256.Alg(),
			}),
		)
		if err != nil {
			if errors.Is(err, jwt.ErrTokenExpired) {
				return ErrTokenExpired
			}

			return ErrInvalidToken
		}

		if !token.Valid {
			return ErrInvalidToken
		}

		if claims.CustCode == "" {
			return ErrCustCodeNotFound
		}

		if claims.AccountNo == "" {
			return ErrAccountNoNotFound
		}

		setIdentity(c, Identity{
			AccountNo: claims.AccountNo,
			CustCode:  claims.CustCode,
		})

		return c.Next()
	}
}
