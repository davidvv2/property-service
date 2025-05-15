package database

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var _ FinderInserter[bson.M, any] = (*FinderInserterMongoImpl[bson.M, any])(nil)

type FinderInserterMongoImpl[Filter, DomainModel any] struct {
	Finder[Filter, DomainModel]
	Inserter[DomainModel]
}

func NewMongoFinderInserter[DomainModel any](
	finder Finder[bson.M, DomainModel],
	inserter Inserter[DomainModel],
) FinderInserter[bson.M, DomainModel] {
	return FinderInserterMongoImpl[bson.M, DomainModel]{
		finder,
		inserter,
	}
}

type FinderInserterUpdaterMongoImpl[Filter, Partial, DomainModel any] struct {
	Finder[Filter, DomainModel]
	Inserter[DomainModel]
	Updater[Filter, Partial, DomainModel]
}

func NewMongoFinderInserterUpdater[DatabaseModel, DomainModel any](
	finder Finder[bson.M, DomainModel],
	inserter Inserter[DomainModel],
	updater Updater[bson.M, bson.M, DomainModel],
) FinderInserterUpdaterMongoImpl[bson.M, bson.M, DomainModel] {
	return FinderInserterUpdaterMongoImpl[bson.M, bson.M, DomainModel]{
		Finder:   finder,
		Inserter: inserter,
		Updater:  updater,
	}
}

var _ EncrypterFinderInserterUpdater[
	bson.M,
	bson.M,
	any,
	primitive.Binary,
	any,
] = (*EncrypterFinderInserterUpdaterMongoImpl[
	bson.M,
	bson.M,
	any,
	primitive.Binary,
	any,
])(nil)

type EncrypterFinderInserterUpdaterMongoImpl[
	Filter,
	Partial,
	DomainModel any,
	EncryptData,
	EncryptedData any,
] struct {
	FinderInserterUpdaterMongoImpl[Filter, Partial, DomainModel]
	EncrypterOperator[EncryptData, EncryptedData, DomainModel]
	Encrypter Encrypter[any, primitive.Binary]
}

func NewMongoEncrypterFinderInserterUpdater[
	DomainModel,
	EncryptData,
	EncryptedData any,
](
	finder Finder[primitive.M, DomainModel],
	inserter Inserter[DomainModel],
	updater Updater[bson.M, bson.M, DomainModel],
	encrypterOperator EncrypterOperator[EncryptData, EncryptedData, DomainModel],
	encrypter Encrypter[any, primitive.Binary],
) EncrypterFinderInserterUpdaterMongoImpl[
	bson.M,
	bson.M,
	DomainModel,
	EncryptData,
	EncryptedData,
] {
	return EncrypterFinderInserterUpdaterMongoImpl[
		bson.M,
		bson.M,
		DomainModel,
		EncryptData,
		EncryptedData,
	]{
		FinderInserterUpdaterMongoImpl: NewMongoFinderInserterUpdater[EncryptedData, DomainModel](finder, inserter, updater),
		EncrypterOperator:              encrypterOperator,
		Encrypter:                      encrypter,
	}
}

type FinderUpdaterMongoImpl[DomainModel any] struct {
	Finder[bson.M, DomainModel]
	Updater[bson.M, bson.M, DomainModel]
}

func NewMongoFinderUpdater[DomainModel any](
	finder Finder[bson.M, DomainModel],
	updater Updater[bson.M, bson.M, DomainModel],
) FinderUpdaterMongoImpl[DomainModel] {
	return FinderUpdaterMongoImpl[DomainModel]{
		Finder:  finder,
		Updater: updater,
	}
}

type FinderInserterUpdaterRemoverMongoImpl[DomainModel any] struct {
	Finder[bson.M, DomainModel]
	Inserter[DomainModel]
	Updater[bson.M, bson.M, DomainModel]
	Remover[bson.M]
}

func NewMongoFinderInserterUpdaterRemover[DomainModel any](
	finder Finder[bson.M, DomainModel],
	inserter Inserter[DomainModel],
	updater Updater[bson.M, bson.M, DomainModel],
	remover Remover[bson.M],
) FinderInserterUpdaterRemoverMongoImpl[DomainModel] {
	return FinderInserterUpdaterRemoverMongoImpl[DomainModel]{
		Finder:   finder,
		Updater:  updater,
		Inserter: inserter,
		Remover:  remover,
	}
}
