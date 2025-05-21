package permissions

import (
	"property-service/pkg/permissions/scopes"
)

type Actor interface {
	GetScope() scopes.Scopes
}

type Validator interface {
	Can(
		actor Actor,
		operation string,
		service string,
		validator func() error,
	) error
}
