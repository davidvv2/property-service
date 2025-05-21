package service

import (
	"property-service/internal/properties/app"
	"property-service/internal/properties/app/command"
)

func (d Dependencies) createCommands() app.Commands {
	return app.Commands{
		// Property commands
		CreateProperty: command.NewCreatePropertyHandler(
			d.Repo.PropertyRepository,
			d.L,
			d.V,
		),
		UpdateProperty: command.NewUpdatePropertyHandler(
			d.Repo.PropertyRepository,
			d.L,
			d.V,
		),
		DeleteProperty: command.NewDeletePropertyHandler(
			d.Repo.PropertyRepository,
			d.L,
			d.V,
		),
		// Owner commands
		CreateOwner: command.NewCreateOwnerHandler(
			d.Repo.OwnerRepository,
			d.L,
			d.V,
		),
		UpdateOwner: command.NewUpdateOwnerHandler(
			d.Repo.OwnerRepository,
			d.L,
			d.V,
		),
		DeleteOwner: command.NewDeleteOwnerHandler(
			d.Repo.OwnerRepository,
			d.L,
			d.V,
		),
	}
}
