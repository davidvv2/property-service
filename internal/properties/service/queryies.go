package service

import (
	"property-service/internal/properties/app"
	"property-service/internal/properties/app/query"
)

func (d Dependencies) createQueries() app.Queries {
	return app.Queries{
		GetProperty: query.NewGetPropertyHandler(
			d.Repo.PropertyRepository,
			d.L,
			d.V,
		),
		GetOwner: query.NewGetOwnerHandler(
			d.Repo.OwnerRepository,
			d.L,
			d.V,
		),
		ListPropertiesByCategory: query.NewListPropertiesByCategoryHandler(
			d.Repo.PropertyRepository,
			d.L,
			d.V,
		),
	}
}
