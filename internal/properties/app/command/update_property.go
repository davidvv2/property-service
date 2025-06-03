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

// UpdatePropertyCommand : This is the update property request in a struct format.
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

// UpdatePropertyHandler is a CQRS endpoint that handles a command to update a property.
// It implements the CommandHandler interface for the VerifyDeviceCommand.
// The handler updates the information of a property in the database.
type UpdatePropertyHandler decorator.CommandHandler[UpdatePropertyCommand]

type UpdatePropertyHandlerImpl struct {
	repository property.Repository
	validator  *validator.Validate
	log        log.Logger
}

// NewUpdatePropertyHandler creates a new instance of UpdatePropertyHandler,
// applying necessary decorators for logging and validation.
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

// Handle the update property command.
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
