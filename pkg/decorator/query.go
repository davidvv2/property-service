package decorator

import (
	"context"

	"property-service/pkg/infrastructure/log"

	"github.com/go-playground/validator/v10"
)

func ApplyQueryDecorators[H any, R any](
	handler QueryHandler[H, R],
	logger log.Logger,
	validator *validator.Validate,
) QueryHandler[H, R] {
	return queryValidationDecorator[H, R]{
		base: queryLoggingDecorator[H, R]{
			base:   handler,
			logger: logger,
		},
		validator: validator,
		logger:    logger,
	}
}

type QueryHandler[Q any, R any] interface {
	Handle(ctx context.Context, q Q) (R, error)
}
