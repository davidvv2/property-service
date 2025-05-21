package database

import (
	"context"

	"property-service/pkg/configs"
	"property-service/pkg/infrastructure/log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	// Deterministic encryption  algorithm.
	deterministic = "AEAD_AES_256_CBC_HMAC_SHA_512-Deterministic"
	// Random encryption  algorithm.
	random = "AEAD_AES_256_CBC_HMAC_SHA_512-Random"
)

type EncrypterMongoImpl[
	EncryptData any,
	EncryptedData primitive.Binary,
] struct {
	connector Connector[mongo.Client, mongo.ClientEncryption, mongo.Collection]
	log       log.Logger
	// Encryption Client for mongodb
	encrypter        *mongo.ClientEncryption
	masterKey        map[string]interface{}
	kmsProvider      string
	collectionSuffix string
}

func NewMongoEncrypter(
	log log.Logger,
	connector Connector[mongo.Client, mongo.ClientEncryption, mongo.Collection],
	config configs.DatabaseStruct,
	collectionSuffix string,
) *EncrypterMongoImpl[any, primitive.Binary] {
	return &EncrypterMongoImpl[any, primitive.Binary]{
		log:       log,
		encrypter: connector.getEncryptionClient(),
		masterKey: map[string]interface{}{
			"projectId": config.CSFLE.ProjectID,
			"location":  config.CSFLE.Location,
			"keyRing":   config.CSFLE.KeyRing,
			"keyName":   config.CSFLE.KeyName,
		},
		kmsProvider:      config.CSFLE.KMSProvider,
		connector:        connector,
		collectionSuffix: collectionSuffix,
	}
}

// CreateDEK:  Creates a new Data Encryption Key in the database. Please use this when creating a new user.
func (emi *EncrypterMongoImpl[EncryptData, EncryptedData]) CreateDEK(
	c context.Context, altKey string,
) error {
	// Create a new DEK and return the error.
	_, err := emi.encrypter.CreateDataKey(
		c, emi.kmsProvider,
		options.DataKey().
			SetMasterKey(emi.masterKey).
			SetKeyAltNames([]string{altKey}),
	)
	return err
}

// DeleteDEK : will delete the key with the altKey provided.
func (emi *EncrypterMongoImpl[EncryptData, EncryptedData]) DeleteDEK(
	c context.Context, altKey string,
) error {
	// Find the key by AltKey name and return a error if one does not exist.
	result := emi.encrypter.GetKeyByAltName(c, altKey)
	var resultKey key
	if err := result.Decode(&resultKey); err != nil {
		return err
	}
	// Delete the key if a key is found.
	_, deleteKeyErr := emi.encrypter.DeleteKey(c, resultKey.ID)
	return deleteKeyErr
}

// EncryptDeterministically: will encrypt the value provided deterministically meaning the same each time.
func (emi *EncrypterMongoImpl[EncryptData, EncryptedData]) Deterministically(
	c context.Context, data EncryptData, keyAltNames string,
) (EncryptedData, error) {
	return emi.encrypt(c, data, keyAltNames, deterministic)
}

// EncryptRandomly : Will encrypt the data randomly each time, this means you can not search for this data.
func (emi *EncrypterMongoImpl[EncryptData, EncryptedData]) Randomly(
	c context.Context, data EncryptData, keyAltNames string,
) (EncryptedData, error) {
	return emi.encrypt(c, data, keyAltNames, random)
}

func (emi *EncrypterMongoImpl[EncryptData, EncryptedData]) encrypt(
	c context.Context, data EncryptData, keyAltNames string, encryptAlgorithm string,
) (EncryptedData, error) {
	// Marshal the document into a bson value and data pair. Note this is uses reflection obviously.
	nameRawValueType, nameRawValueData, err := bson.MarshalValue(data)
	if err != nil {
		return EncryptedData{}, err
	}

	// Encrypt the data and return the randomly encrypted data and a error if one exists.
	encryptedData, err := emi.encrypter.Encrypt(
		c,
		bson.RawValue{Type: nameRawValueType, Value: nameRawValueData},
		options.Encrypt().
			SetAlgorithm(encryptAlgorithm).
			SetKeyAltName(keyAltNames),
	)
	if err != nil {
		emi.log.Debug("errWhileEncrypting %+v on %+v with %+v and %+v", err, data, nameRawValueData, nameRawValueType)
		return EncryptedData{}, err
	}
	return (EncryptedData)(encryptedData), err
}

type key struct {
	MasterKey    masterKey          `bson:"masterKey" validate:"required"`
	ID           primitive.Binary   `bson:"_id" validate:"required"`
	KeyMaterial  primitive.Binary   `bson:"keyMaterial" validate:"required"`
	KeyAltNames  []string           `bson:"keyAltNames" validate:"required"`
	CreationDate primitive.DateTime `bson:"creationDate" validate:"required"`
	UpdateDate   primitive.DateTime `bson:"updateDate" validate:"required"`
	Status       int64              `bson:"status" validate:"required"`
}

type masterKey struct {
	Provider  string `bson:"provider" validate:"required"`
	ProjectID string `bson:"projectId" validate:"required"`
	Location  string `bson:"location" validate:"required"`
	KeyRing   string `bson:"keyRing" validate:"required"`
	KeyName   string `bson:"keyName" validate:"required"`
}
