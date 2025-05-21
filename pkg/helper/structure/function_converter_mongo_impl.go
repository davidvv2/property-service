package structure

import "go.mongodb.org/mongo-driver/bson/primitive"

type IDMongoConverter struct {
}

func NewIDConverter() IDMongoConverter {
	return IDMongoConverter{}
}

func (IDMongoConverter) OldToNew(idHex interface{}) (interface{}, error) {
	id, err := primitive.ObjectIDFromHex(idHex.(string))
	if err != nil {
		return primitive.NilObjectID, err
	}
	return id, nil
}

// NewID creates a new ID from a string representation of a MongoDB ObjectId.
func (IDMongoConverter) NewToOld(primitiveID interface{}) (interface{}, error) {
	return primitiveID.(primitive.ObjectID).Hex(), nil
}
