package database

import (
	"property-service/pkg/errors"

	"github.com/google/uuid"
)

func StringToID(idHex string) (uuid.UUID, error) {
	id, err := uuid.Parse(idHex)

	if err != nil {
		return uuid.Nil, err
	}
	return id, nil
}

func IDToString(id uuid.UUID) (string, error) {
	newID := id.String()
	if newID == "" {
		return "", errors.New("failed to convert ID to string")
	}
	return newID, nil
}

func NewID() uuid.UUID {
	return uuid.New()
}

func NewStringID() string {
	return uuid.New().String()

}
