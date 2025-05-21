// Package app is the application layer of the authentication microservice.
// It contains the business logic for processing user authentication commands and queries.
// It is built following the Clean Architecture principles, implementing DDD Lite and CQRS patterns.
// The Application struct encapsulates the Commands and Queries structs, which define the available
//
//	command and query handlers.
//
// The Commands struct is currently empty, as the current implementation only supports queries.
// The Queries struct contains handlers for querying user login attempt history and login history.
package app

import (
	"property-service/internal/properties/app/command"
	"property-service/internal/properties/app/query"
)

// Application is the entry point for the authentication microservice application.
// It encapsulates the available command and query handlers.
type Application struct {
	Commands Commands
	Queries  Queries
}

// Commands is a struct that holds the application's commands handlers.
type Commands struct {
	CreateProperty command.CreatePropertyHandler
	DeleteProperty command.DeletePropertyHandler
	UpdateProperty command.UpdatePropertyHandler
	CreateOwner    command.CreateOwnerHandler
	DeleteOwner    command.DeleteOwnerHandler
	UpdateOwner    command.UpdateOwnerHandler
}

// Queries is a struct that holds the application's query handlers.
type Queries struct {
	GetProperty              query.GetPropertyHandler
	GetOwner                 query.GetOwnerHandler
	ListPropertiesByCategory query.ListPropertiesByCategoryHandler
}
