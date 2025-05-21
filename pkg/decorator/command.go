package decorator

import (
	"context"
	"fmt"
	"strings"

	"property-service/pkg/infrastructure/log"

	"github.com/go-playground/validator/v10"
)

func ApplyCommandDecorators[H any](
	handler CommandHandler[H], logger log.Logger, validator *validator.Validate,
) CommandHandler[H] {
	return validationDecorator[H]{
		base: commandLoggingDecorator[H]{
			base:   handler,
			logger: logger,
		},
		validator: validator,
		logger:    logger,
	}
}

type CommandHandler[C any] interface {
	Handle(ctx context.Context, cmd C) error
}

func generateActionName(handler any) string {
	return strings.Split(fmt.Sprintf("%T", handler), ".")[1]
}
