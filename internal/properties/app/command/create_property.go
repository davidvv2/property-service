package command

import (
	"context"
	"time"

	"property-service/internal/properties/domain/property"
	"property-service/pkg/address"
	"property-service/pkg/decorator"
	"property-service/pkg/errors"
	"property-service/pkg/errors/codes"
	"property-service/pkg/infrastructure/log"

	"github.com/go-playground/validator/v10"
)

// CreatePropertyCommand : This is the create property request in a struct format.
type CreatePropertyCommand struct {
	PropertyID    string          `validate:"required"`
	OwnerID       string          `validate:"required"`
	Category      string          `validate:"required"`
	Description   string          `validate:"required"`
	Title         string          `validate:"required"`
	Available     bool            `validate:"required"`
	AvailableDate time.Time       `validate:"required"`
	Address       address.Address `validate:"required"`
	SaleType      uint8           `validate:"required"`
}

// CreatePropertyHandler is a CQRS endpoint that handles a command to create a property.
// It implements the CommandHandler interface for the CreatePropertyCommand.
// The handler creates a new property in the database.
type CreatePropertyHandler decorator.CommandHandler[CreatePropertyCommand]

type CreatePropertyHandlerImpl struct {
	repository property.Repository
	validator  *validator.Validate
	log        log.Logger
}

// NewCreatePropertyHandler creates a new instance of CreatePropertyHandler,
// applying necessary decorators for logging and validation.
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

// Validate the create property command.
func (cph CreatePropertyHandlerImpl) Handle(
	c context.Context, cmd CreatePropertyCommand,
) error {
	if _, registerErr := cph.repository.New(
		c,
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
