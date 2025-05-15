package database

import "go.mongodb.org/mongo-driver/bson/primitive"

// NewID creates a new ID from a string representation of a MongoDB ObjectId.
func StringToID(idHex string) (primitive.ObjectID, error) {
	id, err := primitive.ObjectIDFromHex(idHex)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return id, nil
}

// IDToString :: Transforms a mongodb id primitive.objectID to a string.
func IDToString(id primitive.ObjectID) (string, error) {
	return id.Hex(), nil
}

// NewID :: Creates a new database ID.
func NewID() primitive.ObjectID {
	return primitive.NewObjectID()
}

// NewStringID :: Creates a new database string id.
func NewStringID() string {
	return primitive.NewObjectID().Hex()
}

// StringToBatchID Will return a list of lists of mongodb ids in the batch size defined constant.
func StringToBatchID(stringID []string) ([][]primitive.ObjectID, error) {
	lengthIDs := len(stringID)
	// Create the batches of ids.
	var databaseIDs = make([][]primitive.ObjectID, (lengthIDs/BatchSize)+1)
	for i, dbIndex := 0, -1; i < lengthIDs; i++ {
		// Allocate the array for each batch.
		if i%BatchSize == 0 {
			dbIndex++
			switch {
			case lengthIDs < BatchSize:
				databaseIDs[dbIndex] = make([]primitive.ObjectID, lengthIDs)
			default:
				databaseIDs[dbIndex] = make([]primitive.ObjectID, BatchSize)
			}
		}
		// Convert the id to a objectID.
		oid, err := primitive.ObjectIDFromHex(stringID[i])
		if err != nil {
			return nil, err
		}
		databaseIDs[dbIndex][i-(dbIndex*BatchSize)] = oid
	}
	return databaseIDs, nil
}
