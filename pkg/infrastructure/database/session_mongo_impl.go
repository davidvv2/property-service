package database

import (
	"context"

	"property-service/pkg/errors"
	"property-service/pkg/errors/codes"

	"go.mongodb.org/mongo-driver/mongo"
)

var _ Session[SessionReceiver] = (*SessionMongoImpl[SessionReceiver])(nil)

// SessionReceiver : Function.
type SessionReceiver func(c mongo.SessionContext) (interface{}, error)

type SessionMongoImpl[Session SessionReceiver] struct {
	connector Connector[
		mongo.Client,
		mongo.ClientEncryption,
		mongo.Collection,
	]
}

func NewMongoSession(
	connector Connector[mongo.Client, mongo.ClientEncryption, mongo.Collection],
) *SessionMongoImpl[SessionReceiver] {
	return &SessionMongoImpl[SessionReceiver]{
		connector: connector,
	}
}

func (
	gmi *SessionMongoImpl[Session],
) Execute(c context.Context, call Session) error {
	session, err := gmi.connector.getClient().StartSession()
	if err != nil {
		return errors.NewInfrastructureError(
			err,
			codes.Internal,
		)
	}
	defer session.EndSession(c)
	_, e := session.WithTransaction(c, call)
	return e
}
