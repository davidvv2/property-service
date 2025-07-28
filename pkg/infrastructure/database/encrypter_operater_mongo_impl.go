package database

import (
	"context"
	"reflect"

	"property-service/pkg/errors"
	"property-service/pkg/helper/factory"
	"property-service/pkg/helper/structure"
	"property-service/pkg/infrastructure/log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var _ EncrypterOperator[
	any, primitive.Binary, any,
] = (*EncrypterOperatorMongoImpl[
	any, primitive.Binary, any, any,
])(nil)

const (
	Deterministic = "deterministic"
	Random        = "random"
	BSON          = "bson"
	ENCRYPT       = "encrypt"
)

type EncrypterOperatorMongoImpl[
	EncryptData any,
	EncryptedData primitive.Binary,
	DomainModel, DatabaseModel any] struct {
	connector Connector[mongo.Client, mongo.ClientEncryption, mongo.Collection]
	Encrypter[EncryptData, EncryptedData]
	inserter    Inserter[map[string]interface{}]
	converter   structure.Converter[DomainModel]
	log         log.Logger
	encryptTags map[string][]string
	collection  string
}

func NewMongoEncrypterOperator[DomainModel, DatabaseModel any](
	connector Connector[mongo.Client, mongo.ClientEncryption, mongo.Collection],
	encrypter Encrypter[any, primitive.Binary],
	converter structure.Converter[DomainModel],
	log log.Logger,
	collection string,
) *EncrypterOperatorMongoImpl[any, primitive.Binary, DomainModel, DatabaseModel] {
	// Create a new inserter.
	inserter := NewMongoInserter[map[string]interface{}]( //
		log,
		collection,
		newFakeFactory(),
		connector,
	)

	tag, err := converter.
		GetOldManipulator().
		GetAllTagWithValidator(ENCRYPT, func(tag []string, value reflect.Kind) ([]string, error) {
			log.Info("Tags %+v", tag)
			if len(tag) > 1 {
				log.Panic("the encrypt tag must only have one value but has two %+v", tag)
			}
			if value == reflect.Bool && tag[0] == Deterministic {
				log.Panic("bool can not be of type deterministic encryption %+v", tag)
			}
			return tag, nil
		})
	if err != nil {
		log.Panic("error has occurred %+v", err)
	}
	var strModel DomainModel
	log.Info("created new mongodb encrypter operator for %T with encrypted fields %+v ", strModel, tag)
	return &EncrypterOperatorMongoImpl[any, primitive.Binary, DomainModel, DatabaseModel]{
		connector:   connector,
		log:         log,
		encryptTags: tag,
		Encrypter:   encrypter,
		converter:   converter,
		collection:  collection,
		inserter:    inserter,
	}
}

func (eomi EncrypterOperatorMongoImpl[EncryptData, EncryptedData, DomainModel, DatabaseModel],
) InsertDocument(c context.Context, key string, data DomainModel) (string, error) {
	// Create a new DEK for the alt key provided.
	errCreatingDek := eomi.Encrypter.CreateDEK(c, key)

	// Flatten the struct passed.
	flat, conversionErr := eomi.converter.Convert(&data, BSON)
	var err error
	switch {
	case conversionErr != nil:
		err = errors.NewDatabaseError(conversionErr)
	case errCreatingDek != nil:
		err = errors.NewDatabaseError(errCreatingDek)
	}
	if err != nil {
		return "", err
	}

	bsonEncrypt, err := eomi.encryptFields(c, flat, key)
	if err != nil {
		return "", err
	}

	res, err := eomi.inserter.InsertOne(c, bsonEncrypt)
	if err != nil {
		return "", errors.NewDatabaseError(err)
	}
	// Return the resulting ID.
	return res, nil
}

func (eomi EncrypterOperatorMongoImpl[EncryptData, EncryptedData, DomainModel, DatabaseModel],
) encryptFields(
	c context.Context, document map[string]interface{}, key string,
) (map[string]interface{}, error) {
	for k, v := range eomi.encryptTags {
		var err error
		if v[0] == Deterministic {
			document[k], err = eomi.Deterministically(c, (document[k]).(EncryptData), key)
		} else {
			document[k], err = eomi.Randomly(c, (document[k]).(EncryptData), key)
		}
		if err != nil {
			eomi.log.Debug("error while encrypting %+v", err)
			return nil, err
		}
	}
	return document, nil
}

var _ factory.Factory[map[string]interface{}, map[string]interface{}] = (*fakeFactory)(nil)

type fakeFactory struct {
}

// ToDatabase implements factory.Factory.
func (*fakeFactory) ToDatabase(i map[string]interface{}) (*map[string]interface{}, error) {
	return &i, nil
}

func (*fakeFactory) ToDomain(i map[string]interface{}) (*map[string]interface{}, error) {
	return &i, nil
}

func newFakeFactory() *fakeFactory {
	return &fakeFactory{}
}
