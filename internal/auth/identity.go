package auth

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

const identityKey = "identity"

var ErrIdentityNotFound = errors.New("identity not found")

type Identity struct {
	CustCode  string
	AccountID string
}

func setIdentity(c *fiber.Ctx, identity Identity) {
	c.Locals(identityKey, identity)
}

func CurrentIdentity(c *fiber.Ctx) (Identity, error) {
	identity, ok := c.Locals(identityKey).(Identity)
	if !ok {
		return Identity{}, ErrIdentityNotFound
	}

	return identity, nil
}
