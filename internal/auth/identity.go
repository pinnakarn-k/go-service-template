package auth

import (
	"github.com/gofiber/fiber/v2"
)

const identityKey = "identity"

type Identity struct {
	AccountNo string
	CustCode  string
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
