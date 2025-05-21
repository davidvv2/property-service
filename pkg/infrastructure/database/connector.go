package database

// Connector is a database connector interface that is used to provide connections to the underling database.
type Connector[Client any, EncryptedClient any, Collection any] interface {
	// GetClient will return the database client.
	getClient() *Client

	// GetDatabaseName will return the name of the database that the object is connected to.
	GetDatabaseName() string

	// getEncryptedClient will return a database client that is used for client level encryption. This normally involves
	// a separate persistent storage mechanism that stores the key ids for csfle that this client interfaces with.
	getEncryptionClient() *EncryptedClient

	// GetCollection will return a collection with the name given, if the collection does not exist it will return a
	// error.
	GetCollection(collectionName string) (*Collection, error)

	Ping() bool
}
