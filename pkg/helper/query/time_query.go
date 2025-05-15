package query

import (
	"property-service/pkg/infrastructure/database"
)

type Time database.TimeQuery[string]

// Query : meta struct passed to get user list.
type Query database.Query[string]
