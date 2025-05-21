// Package jwt provides json web token signing and verifying from a object.
package jwt

import (
	"context"
	"time"
)

// Manager creates a jwt manager that allows you to sign, verify, blacklist and check blacklist tokens of its creation
// type.
type Manager[T AuthClaims] interface {
	// Sign a jwt token
	Sign(token T) (string, error)

	// Verify that a token is valid and return the claims.
	Verify(jwtString string) (*T, error)

	// Check to see if a token is on the blacklist.
	CheckBlacklist(c context.Context, UUID string) (bool, error)

	// Blacklist a given tokens id.
	BlackList(UUID string, expire time.Duration) error
}
