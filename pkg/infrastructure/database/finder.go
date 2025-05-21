package database

import (
	"context"
)

type Finder[
	Filter,
	DomainModel any,
] interface {
	Find(
		c context.Context,
		server string,
		filter Filter,
	) (
		<-chan *DomainModel,
		<-chan error,
		<-chan bool,
	)

	// FindByIDs
	FindByIDs(
		c context.Context,
		server string,
		ids []string,
	) (<-chan *DomainModel, <-chan error, <-chan bool)

	//FindOne: will return one document from the database that matches the filter.
	FindOne(
		c context.Context,
		server string,
		filter Filter,
	) (*DomainModel, error)

	//FindByID: Will find a item in the database by the id provided.
	FindByID(
		c context.Context,
		server,
		id string,
	) (*DomainModel, error)

	// Count:
	Count(
		c context.Context,
		server string,
		filter Filter,
	) (int64, error)

	// DocumentExists:
	DocumentExists(
		c context.Context,
		server string,
		id string,
	) (int64, error)

	Query(
		c context.Context,
		server string,
		filter Filter,
		query Query[string],
	) (
		<-chan *DomainModel,
		<-chan error,
		<-chan bool,
	)

	TimeQuery(
		c context.Context,
		server string,
		filter Filter,
		timeQuery TimeQuery[string],
	) (
		<-chan *DomainModel,
		int64,
		<-chan error,
		<-chan bool,
	)
}
