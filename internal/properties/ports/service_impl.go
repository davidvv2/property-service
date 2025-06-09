package port

import (
	"context"
	"property-service/internal/properties/app"
	"property-service/internal/properties/app/command"
	"property-service/internal/properties/app/query"
	"property-service/internal/properties/domain/owner"
	"property-service/internal/properties/domain/property"
	"property-service/internal/properties/service"
	"property-service/pkg/configs"
	"property-service/pkg/infrastructure/log"
)

// ServiceImpl: holds the Dependencies.
type ServiceImpl struct {
	App app.Application
	Log log.Logger
}

func NewService(
	configs configs.Config,
	log log.Logger,
) *ServiceImpl {
	return &ServiceImpl{
		App: service.NewApplication(
			configs,
		),
		Log: log,
	}
}

// Property CRUD operations
func (s *ServiceImpl) CreateProperty(
	ctx context.Context,
	params command.CreatePropertyCommand,
) error {
	return s.App.Commands.CreateProperty.Handle(ctx, params)
}

func (s *ServiceImpl) UpdateProperty(
	ctx context.Context,
	params command.UpdatePropertyCommand,
) error {
	return s.App.Commands.UpdateProperty.Handle(ctx, params)
}

func (s *ServiceImpl) DeleteProperty(
	ctx context.Context,
	params command.DeletePropertyCommand,
) error {
	return s.App.Commands.DeleteProperty.Handle(ctx, params)
}

func (s *ServiceImpl) GetProperty(
	ctx context.Context,
	params query.GetPropertyQuery,
) (*property.Property, error) {
	return s.App.Queries.GetProperty.Handle(ctx, params)
}

func (s *ServiceImpl) ListPropertiesByCategory(
	ctx context.Context,
	params query.ListPropertiesByCategoryQuery,
) (*query.ListPropertiesByCategoryResult, error) {
	return s.App.Queries.ListPropertiesByCategory.Handle(ctx, params)
}

func (s *ServiceImpl) ListPropertiesByOwner(
	ctx context.Context,
	params query.ListPropertiesByOwnerQuery,
) (*query.ListPropertiesByOwnerResult, error) {
	return s.App.Queries.ListPropertiesByOwner.Handle(ctx, params)
}

// Owner CRUD operations
func (s *ServiceImpl) CreateOwner(
	ctx context.Context,
	params command.CreateOwnerCommand,
) error {
	return s.App.Commands.CreateOwner.Handle(ctx, params)
}

func (s *ServiceImpl) UpdateOwner(
	ctx context.Context,
	params command.UpdateOwnerCommand,
) error {
	return s.App.Commands.UpdateOwner.Handle(ctx, params)
}

func (s *ServiceImpl) DeleteOwner(
	ctx context.Context,
	params command.DeleteOwnerCommand,
) error {
	return s.App.Commands.DeleteOwner.Handle(ctx, params)
}

func (s *ServiceImpl) GetOwner(
	ctx context.Context,
	params query.GetOwnerQuery,
) (*owner.Owner, error) {
	return s.App.Queries.GetOwner.Handle(ctx, params)
}
