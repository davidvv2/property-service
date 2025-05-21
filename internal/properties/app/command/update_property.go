package command

import (
	"context"
	"time"

	"property-service/internal/properties/domain/property"
	"property-service/pkg/decorator"
	"property-service/pkg/errors"
	"property-service/pkg/errors/codes"
	"property-service/pkg/infrastructure/log"

	"github.com/go-playground/validator/v10"
)

// UpdatePropertyCommand : This is the register request in a struct format.
type UpdatePropertyCommand struct {
	PropertyID    string `validate:"required"`
	Available     *bool
	AvailableDate time.Time
	Description   string
	Title         string
	Category      string
	Address       string
	SaleType      uint8
	Server        string `validate:"required"`
}

// Validate the register User command.
func (cph UpdatePropertyHandlerImpl) Handle(
	c context.Context, cmd UpdatePropertyCommand,
) error {
	if registerErr := cph.repository.Update(
		c,
		cmd.Server,
		cmd.PropertyID,
		property.UpdatePropertyParams{
			Available:     cmd.Available,
			AvailableDate: cmd.AvailableDate,
			Description:   cmd.Description,
			Title:         cmd.Title,
			Category:      cmd.Category,
			Address:       cmd.Address,
			SaleType:      cmd.SaleType,
		},
	); registerErr != nil {
		return errors.NewHandlerError(
			registerErr,
			codes.Internal,
		)
	}
	return nil
}

// UpdatePropertyHandler is a CQRS endpoint that handles a command to retrieve a user's login attempt history.
// It implements the CommandHandler interface for the VerifyDeviceCommand.
// The handler retrieves the user's login attempt history from the database and returns it to the caller.
type UpdatePropertyHandler decorator.CommandHandler[UpdatePropertyCommand]

type UpdatePropertyHandlerImpl struct {
	repository property.Repository
	validator  *validator.Validate
	log        log.Logger
}

// NewUpdatePropertyHandler : handles the login attempt query.
func NewUpdatePropertyHandler(
	repository property.Repository,
	logger log.Logger,
	validator *validator.Validate,
) UpdatePropertyHandler {
	if repository == nil {
		logger.Panic("nil repository")
	}
	return decorator.ApplyCommandDecorators(
		UpdatePropertyHandlerImpl{
			repository: repository,
			validator:  validator,
			log:        logger,
		},
		logger,
		validator,
	)
}
