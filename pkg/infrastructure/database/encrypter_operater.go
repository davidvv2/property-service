package database

import "context"

type EncrypterOperator[
	EncryptData,
	EncryptedData,
	DomainModel any,
] interface {
	Encrypter[EncryptData, EncryptedData]
	InsertDocument(c context.Context, key string, data DomainModel) (string, error)
}
