package database

import (
	"context"
	"strings"
	"time"

	"property-service/pkg/configs"
	"property-service/pkg/errors"
	"property-service/pkg/infrastructure/log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo"
)

var _ Connector[mongo.Client, mongo.ClientEncryption, mongo.Collection] = (*ConnectorMongoImpl[
	mongo.Client, mongo.ClientEncryption, mongo.Collection,
])(nil) // Verify that the mongodb connector implements the Connector interface.

const (
	mongoConnectionTimeout = 10 * time.Second
	// Vault name is the collection for the key vault.
	vaultName = ".__keyVault"
)

type ConnectorMongoImpl[
	Client mongo.Client, EncryptionClient mongo.ClientEncryption, Collection mongo.Collection,
] struct {
	log log.Logger
	// Client is the MongoDB client instance.
	client *Client
	// encryptionClient creates the encryption client.
	encryptionClient *EncryptionClient
	// Database is the MongoDB database instance.
	Database *mongo.Database
	// Collections is a map of collection names to collection instances.
	Collections map[string]*Collection
	// A list of the collection Names.
	CollectionsNames []string
}

func NewMongoConnector(
	log log.Logger, config configs.DatabaseStruct, databaseName string,
) *ConnectorMongoImpl[mongo.Client, mongo.ClientEncryption, mongo.Collection] {
	c, cancel := context.WithTimeout(context.Background(), mongoConnectionTimeout)
	kms := map[string]map[string]interface{}{
		config.CSFLE.KMSProvider: {
			"email":      config.CSFLE.Email,
			"privateKey": config.CSFLE.PrivateKey,
		}}

	client, err := mongo.Connect(c, setOptions(config.URI, databaseName, kms))
	if err != nil {
		cancel()
		log.Panic("failed to connect to MongoDB %+v", err)
		return nil
	}

	// ping the database.
	err = client.Ping(c, nil)
	if err != nil {
		cancel()
		log.Panic(err.Error())
	}

	encryptionClient, encryptionClientErr := newEncryptClient(client, databaseName, kms)
	if encryptionClientErr != nil {
		cancel()
		log.Panic(err.Error())
	}
	database := client.Database(databaseName)
	log.Info("        Connected to MongoDB " + databaseName)
	collectionNames, collections := getCollections(
		c, database, log,
	)

	defer cancel()
	return &ConnectorMongoImpl[mongo.Client, mongo.ClientEncryption, mongo.Collection]{
		log:              log,
		client:           client,
		Database:         database,
		encryptionClient: encryptionClient,
		Collections:      collections,
		CollectionsNames: collectionNames,
	}
}

func (
	cmi ConnectorMongoImpl[Client, EncryptionClient, Collection],
) Ping() bool {
	client := (mongo.Client)(*cmi.client)
	if err := client.Ping(context.Background(), nil); err != nil {
		return false
	}
	return true
}

// /nolint: unused // this is a bug.
func (
	cmi ConnectorMongoImpl[Client, EncryptionClient, Collection],
) getClient() *Client {
	return cmi.client
}

// /nolint: unused // this is a bug.
func (
	cmi ConnectorMongoImpl[Client, EncryptionClient, Collection],
) getEncryptionClient() *EncryptionClient {
	return cmi.encryptionClient
}

func (
	cmi ConnectorMongoImpl[Client, EncryptionClient, Collection],
) GetDatabaseName() string {
	return cmi.Database.Name()
}

func (
	cmi *ConnectorMongoImpl[Client, EncryptionClient, Collection],
) UpdateCollection(c context.Context) error {
	collectionNames, _ := getCollections(c, cmi.Database, cmi.log)
	cmi.CollectionsNames = collectionNames
	return nil
}

func (
	cmi ConnectorMongoImpl[Client, EncryptionClient, Collection],
) GetCollection(collectionName string) (*Collection, error) {
	if collection, ok := cmi.Collections[collectionName]; ok {
		return collection, nil
	}
	cmi.log.Debug("failed to get collection %+v", collectionName)
	return nil, errors.ErrCollectionNotFound
}

func setOptions(uri, database string, kms map[string]map[string]interface{}) *options.ClientOptions {
	opts := options.
		Client().
		ApplyURI(uri)
	// Mongo OpenTelemetry instrumentation.
	opts.Monitor = otelmongo.NewMonitor()
	opts.AutoEncryptionOptions = options.AutoEncryption().
		SetKmsProviders(kms).
		SetKeyVaultNamespace(database + vaultName).
		SetBypassAutoEncryption(true)

	return opts
}

func newEncryptClient(
	client *mongo.Client, database string, kms map[string]map[string]interface{},
) (*mongo.ClientEncryption, error) {
	return mongo.NewClientEncryption(client,
		//nolint:exhaustruct // other options not needed.
		&options.ClientEncryptionOptions{
			KeyVaultNamespace: database + vaultName,
			KmsProviders:      kms,
		})
}

func getCollections(
	c context.Context, db *mongo.Database, log log.Logger,
) ([]string, map[string]*mongo.Collection) {
	names, err := db.ListCollectionNames(c, bson.D{})
	if err != nil {
		log.Panic("", err.Error())
	}
	collectionsNames := make([]string, len(names))
	collections := make(map[string]*mongo.Collection, len(names))
	log.Info("        MongoDB Collections")
	arrayIndex := 0
	for _, name := range names {
		if strings.Contains(name, "system.buckets") {
			continue
		}
		log.Info("        	collections %v", name)
		collectionsNames[arrayIndex] = name
		collections[name] = db.Collection(name)
		arrayIndex++
	}
	return collectionsNames, collections
}
