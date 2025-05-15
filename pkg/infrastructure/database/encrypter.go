package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var _ Encrypter[any, primitive.Binary] = (*EncrypterMongoImpl[any, primitive.Binary])(nil)

type Encrypter[Data, EncryptedData any] interface {
	CreateDEK(c context.Context, altKey string) error
	DeleteDEK(c context.Context, altKey string) error
	Deterministically(c context.Context, data Data, keyAltNames string) (EncryptedData, error)
	Randomly(c context.Context, data Data, keyAltNames string) (EncryptedData, error)
}
