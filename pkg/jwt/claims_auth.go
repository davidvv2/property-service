package jwt

import "property-service/pkg/permissions/scopes"

type Permission interface {
	IsRefreshToken() bool
	IsPasswordResetToken() bool
	IsEmailVerified() bool
	GetScope() scopes.Scopes
}
