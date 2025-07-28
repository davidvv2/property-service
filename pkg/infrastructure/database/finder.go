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
		filter Filter,
	) (
		<-chan *DomainModel,
		<-chan error,
		<-chan bool,
	)

	//FindOne: will return one document from the database that matches the filter.
	FindOne(
		c context.Context,
		filter Filter,
	) (*DomainModel, error)

	//FindByID: Will find a item in the database by the id provided.
	FindByID(
		c context.Context,
		id string,
	) (*DomainModel, error)

	// Count:
	Count(
		c context.Context,
		filter Filter,
	) (int64, error)

	// DocumentExists:
	DocumentExists(
		c context.Context,
		id string,
	) (int64, error)

	Query(
		c context.Context,
		filter Filter,
		query Query[string],
	) (
		<-chan *DomainModel,
		<-chan error,
		<-chan bool,
	)

	TimeQuery(
		c context.Context,
		filter Filter,
		timeQuery TimeQuery[string],
	) (
		<-chan *DomainModel,
		int64,
		<-chan error,
		<-chan bool,
	)
}
