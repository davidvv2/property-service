// Package app is the application layer of the property microservice.
// It encapsulates the business logic for processing property and owner commands and queries.
// The design follows Clean Architecture principles using DDD Lite and CQRS patterns.
// The Application struct aggregates the available command and query handlers.
package app

import (
	"property-service/internal/properties/app/command"
	"property-service/internal/properties/app/query"
)

// Application is the entry point for the property microservice application.
// It groups the command and query handlers into separate categories.
type Application struct {
	Commands Commands
	Queries  Queries
}

// Commands holds the command handlers for processing property and owner actions.
type Commands struct {
	CreateProperty command.CreatePropertyHandler
	DeleteProperty command.DeletePropertyHandler
	UpdateProperty command.UpdatePropertyHandler
	CreateOwner    command.CreateOwnerHandler
	DeleteOwner    command.DeleteOwnerHandler
	UpdateOwner    command.UpdateOwnerHandler
}

// Queries holds the query handlers for retrieving property and owner information.
type Queries struct {
	GetProperty              query.GetPropertyHandler
	GetOwner                 query.GetOwnerHandler
	ListPropertiesByCategory query.ListPropertiesByCategoryHandler
}
