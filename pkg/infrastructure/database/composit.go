package database

type FinderRemover[Filter, DomainModel any] interface {
	Finder[Filter, DomainModel]
	Remover[Filter]
}

type FinderInserter[Filter, DomainModel any] interface {
	Finder[Filter, DomainModel]
	Inserter[DomainModel]
}

type FinderHelper[Filter, Options, Query, DateQuery, DomainModel any] interface {
	Finder[Filter, DomainModel]
	QueryHelper[Query, DateQuery, Options, Filter]
}

type FinderInserterUpdater[Filter, Partial, DomainModel any] interface {
	Finder[Filter, DomainModel]
	Inserter[DomainModel]
	Updater[Filter, Partial, DomainModel]
}

type FinderInserterUpdaterRemover[Filter, Partial, DomainModel any] interface {
	Finder[Filter, DomainModel]
	Inserter[DomainModel]
	Updater[Filter, Partial, DomainModel]
	Remover[Filter]
}

type FinderInserterUpdaterRemoverGrouper[Filter, Partial, Pipeline, DomainModel any] interface {
	Finder[Filter, DomainModel]
	Inserter[DomainModel]
	Updater[Filter, Partial, DomainModel]
	Remover[Filter]
	Grouper[Pipeline, DomainModel]
}

type EncrypterFinderInserterUpdater[
	Filter,
	Partial,
	DomainModel,
	EncryptData,
	EncryptedData any,
] interface {
	Finder[Filter, DomainModel]
	Inserter[DomainModel]
	Updater[Filter, Partial, DomainModel]
	EncrypterOperator[EncryptData, EncryptedData, DomainModel]
}

type FinderUpdater[Filter, Partial, DomainModel any] interface {
	Finder[Filter, DomainModel]
	Updater[Filter, Partial, DomainModel]
}
