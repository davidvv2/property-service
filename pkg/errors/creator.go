package errors

import (
	"property-service/pkg/errors/codes"
	"property-service/pkg/infrastructure/log"
)

type Creator interface {
	NewHandlerError(
		err error,
		translation string,
		code codes.Code,
	) AppError

	NewDomainError(
		err error,
		translation string,
		code codes.Code,
	) AppError

	NewRepositoryError(
		err error,
		translation string,
		code codes.Code,
	) AppError

	NewInfrastructureError(
		err error,
		translation string,
		code codes.Code,
	) AppError

	NewClientError(
		err error,
		translation string,
		code codes.Code,
	) AppError

	NewInternalError(
		err error,
	) AppError

	NewDatabaseError(
		err error,
	) AppError

	NewInvalidArgumentError(
		err error,
		translation string,
	) AppError

	NewAuthenticationError(
		err error,
		translation string,
	) AppError
}

type CreatorImpl struct {
	LayerErrorImpl
}

func NewCreator(l log.Logger) CreatorImpl {
	return CreatorImpl{
		newLayerCreator(l),
	}
}
