package service

import (
	"property-service/pkg/configs"
	"property-service/pkg/infrastructure/log"
)

type client struct {
}

func createClients(
	l log.Logger,
	config *configs.Config,
) client {
	return client{}
}
