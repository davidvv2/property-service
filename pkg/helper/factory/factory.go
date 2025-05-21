package factory

type Factory[DomainModel, DatabaseModel any] interface {
	ToDomain(DatabaseModel) (*DomainModel, error)
	ToDatabase(DomainModel) (*DatabaseModel, error)
}
