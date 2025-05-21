package decorator

import (
	"context"

	"property-service/pkg/errors"
	"property-service/pkg/infrastructure/log"

	"github.com/go-playground/validator/v10"
)

type validationDecorator[C any] struct {
	base      CommandHandler[C]
	validator *validator.Validate
	logger    log.Logger
}

func (d validationDecorator[C]) Handle(c context.Context, cmd C,
) error {
	if err := d.validator.Struct(cmd); err != nil {
		d.logger.Error("Invalid request: %+v", err.Error())
		return errors.NewInvalidArgumentError(err)
	}
	return d.base.Handle(c, cmd)
}

type queryValidationDecorator[Q any, R any] struct {
	base      QueryHandler[Q, R]
	validator *validator.Validate
	logger    log.Logger
}

func (d queryValidationDecorator[C, R]) Handle(c context.Context, cmd C,
) (R, error) {
	// Validate the request.
	if reqValidationErr := d.validator.Struct(cmd); reqValidationErr != nil {
		d.logger.Error("Invalid request: %+v", reqValidationErr.Error())
		return *new(R), errors.NewInvalidArgumentError(
			reqValidationErr,
		)
	}
	// Carry out the endpoint.
	res, err := d.base.Handle(c, cmd)
	// Validate the response being sent back.
	if resValidationErr := d.validator.Struct(res); resValidationErr != nil {
		d.logger.Error("Invalid response: %+v", resValidationErr.Error())
		return *new(R), errors.NewInternalError(err)
	}
	// Return the  response
	return res, err
}
