package command

import (
	"context"

	"property-service/internal/properties/domain/owner"
	"property-service/pkg/decorator"
	"property-service/pkg/errors"
	"property-service/pkg/errors/codes"
	"property-service/pkg/infrastructure/log"

	"github.com/go-playground/validator/v10"
)

// CreateOwnerCommand : This is the create owner request in a struct format.
type CreateOwnerCommand struct {
	OwnerID   string `validate:"required"`
	Name      string `validate:"required"`
	Email     string `validate:"required"`
	Telephone string `validate:"required"`
	Server    string `validate:"required"`
}

// CreateOwnerHandler is a CQRS endpoint that handles a command to create an owner.
// It implements the CommandHandler interface for the CreateOwnerCommand.
// The handler creates a new owner in the database.
type CreateOwnerHandler decorator.CommandHandler[CreateOwnerCommand]

type CreateOwnerHandlerImpl struct {
	repository owner.Repository
	validator  *validator.Validate
	log        log.Logger
}

// NewCreateOwnerHandler creates a new instance of CreateOwnerHandler,
// applying necessary decorators for logging and validation.
func NewCreateOwnerHandler(
	repository owner.Repository,
	logger log.Logger,
	validator *validator.Validate,
) CreateOwnerHandler {
	if repository == nil {
		logger.Panic("nil repository")
	}
	return decorator.ApplyCommandDecorators(
		CreateOwnerHandlerImpl{
			repository: repository,
			validator:  validator,
			log:        logger,
		},
		logger,
		validator,
	)
}

// Validate the create owner command.
func (cph CreateOwnerHandlerImpl) Handle(
	c context.Context, cmd CreateOwnerCommand,
) error {
	if _, registerErr := cph.repository.New(
		c,
		cmd.Server,
		owner.NewOwnerParams{
			ID:        cmd.OwnerID,
			Name:      cmd.Name,
			Email:     cmd.Email,
			Telephone: cmd.Telephone,
		},
	); registerErr != nil {
		return errors.NewHandlerError(
			registerErr,
			codes.Internal,
		)
	}
	return nil
}
