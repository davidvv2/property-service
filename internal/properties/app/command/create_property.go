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

// CreatePropertyCommand : This is the register request in a struct format.
type CreatePropertyCommand struct {
	PropertyID    string    `validate:"required"`
	OwnerID       string    `validate:"required"`
	Category      string    `validate:"required"`
	Description   string    `validate:"required"`
	Title         string    `validate:"required"`
	Available     bool      `validate:"required"`
	AvailableDate time.Time `validate:"required"`
	Address       string    `validate:"required"`
	SaleType      uint8     `validate:"required"`
	Server        string    `validate:"required"`
}

// Validate the register User command.
func (cph CreatePropertyHandlerImpl) Handle(
	c context.Context, cmd CreatePropertyCommand,
) error {
	if _, registerErr := cph.repository.New(
		c,
		cmd.Server,
		property.NewPropertyParams{
			PropertyID:    cmd.PropertyID,
			OwnerID:       cmd.OwnerID,
			Category:      cmd.Category,
			Description:   cmd.Description,
			Title:         cmd.Title,
			Available:     cmd.Available,
			AvailableDate: cmd.AvailableDate,
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

// CreatePropertyHandler is a CQRS endpoint that handles a command to retrieve a user's login attempt history.
// It implements the CommandHandler interface for the VerifyDeviceCommand.
// The handler retrieves the user's login attempt history from the database and returns it to the caller.
type CreatePropertyHandler decorator.CommandHandler[CreatePropertyCommand]

type CreatePropertyHandlerImpl struct {
	repository property.Repository
	validator  *validator.Validate
	log        log.Logger
}

// NewCreatePropertyHandler : handles the login attempt query.
func NewCreatePropertyHandler(
	repository property.Repository,
	logger log.Logger,
	validator *validator.Validate,
) CreatePropertyHandler {
	if repository == nil {
		logger.Panic("nil repository")
	}
	return decorator.ApplyCommandDecorators(
		CreatePropertyHandlerImpl{
			repository: repository,
			validator:  validator,
			log:        logger,
		},
		logger,
		validator,
	)
}
