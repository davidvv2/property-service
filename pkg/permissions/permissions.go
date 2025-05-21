package permissions

import (
	"property-service/pkg/errors"
	"property-service/pkg/errors/codes"
	"property-service/pkg/infrastructure/log"
	"property-service/pkg/permissions/scopes"
)

var _ Validator = (*ValidatorJWTImpl)(nil)

type ValidatorJWTImpl struct {
	log log.Logger
}

func NewJWTValidator(
	log log.Logger,
) *ValidatorJWTImpl {
	return &ValidatorJWTImpl{
		log: log,
	}
}

// Can implements Validator.
func (vji *ValidatorJWTImpl) Can(
	actor Actor,
	operation string,
	service string,
	validator func() error,
) error {
	err := validator()

	switch {
	case err != nil:
		vji.log.Debug("Error While Validating using validator %+v", err)
		err = errors.NewAuthenticationError(err)
	case !check(actor.GetScope(), service, operation):
		vji.log.Debug("permission invalid ")
		err = errors.NewHandlerError(
			errors.ErrPermissionDenied,
			codes.PermissionDenied,
		)
	}

	if err != nil {
		return err
	}

	return nil
}

//nolint:cyclop // Its a simple function
func check(sc scopes.Scopes, service string, permission string) bool {

	return true
}
