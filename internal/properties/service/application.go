package service

import (
	"property-service/internal/properties/app"
	"property-service/pkg/configs"
	redis "property-service/pkg/infrastructure/cache"
	"property-service/pkg/infrastructure/log"

	"github.com/go-playground/validator/v10"
)

type Dependencies struct {
	Clients client
	Repo    repositories
	Factory factories
	L       log.Logger
	Cacher  redis.Cacher
	Jwt     jwtManagers
	V       *validator.Validate
	Config  configs.Config
}

func NewApplication(config configs.Config) app.Application {
	dep := BuildDependencies(config)
	return app.Application{
		Commands: dep.createCommands(),
		Queries:  dep.createQueries(),
	}
}

func BuildDependencies(config configs.Config) Dependencies {
	// Creates a new logger for the applications.
	logger := log.NewZapImpl(&config.Backend)
	// Creates a new cacher for the application.
	cacher := redis.NewRedisCacher(&config.Caching, logger)
	// Creates a validator, this is used for validating structs.
	validator := validator.New()
	// return the dependency object.
	factories := createFactories(logger, validator, &config)
	return Dependencies{
		Config:  config,
		L:       logger,
		Cacher:  cacher,
		V:       validator,
		Jwt:     createJWTManagers(logger, cacher, validator),
		Clients: createClients(logger, &config),
		Repo:    createRepositories(logger, &config, factories, validator),
		Factory: factories,
	}
}
