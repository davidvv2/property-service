// Package decorator: this package is used to add functionality to command and query operators
package decorator

import (
	"context"

	"property-service/pkg/infrastructure/log"
)

type commandLoggingDecorator[C any] struct {
	base   CommandHandler[C]
	logger log.Logger
}

func (d commandLoggingDecorator[C]) Handle(ctx context.Context, cmd C) error {
	handlerType := generateActionName(cmd)
	d.logger.Debug("Executing command")

	err := d.base.Handle(ctx, cmd)

	if err == nil {
		d.logger.Info("command %+v executed successfully %+v", handlerType, cmd)
	} else {
		d.logger.Error("Failed to execute command  %+v with  %+v. Error:%+v",
			handlerType, cmd, err)
	}

	return err
}

type queryLoggingDecorator[C any, R any] struct {
	base   QueryHandler[C, R]
	logger log.Logger
}

func (d queryLoggingDecorator[C, R]) Handle(ctx context.Context, cmd C) (R, error) {
	handlerType := generateActionName(cmd)
	d.logger.Debug("Executing query")
	b, err := d.base.Handle(ctx, cmd)

	if err == nil {
		d.logger.Info("Query %+v executed successfully %+v", handlerType, cmd)
	} else {
		d.logger.Error("Failed to execute query  %+v with  %+v. Error:%+v",
			handlerType, cmd, err)
	}

	return b, err
}
