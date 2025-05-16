package database

import "go.mongodb.org/mongo-driver/bson"

var _ PaginationHelper = (*PaginationHelperMongoImpl)(nil)

type PaginationHelperMongoImpl struct{}

// TextPaginationHelper is an implementation of PaginationHelper for text searches.
func (t *PaginationHelperMongoImpl) TextPaginationHelper(
	index string,
	path string,
	value string,
	sort bson.D, // changed to bson.D for ordered sort criteria
	search uint8, // 0: no token, 1: searchAfter, 2: searchBefore
	paginationToken string,
) (bson.D, error) {
	// Build the inner $search stage as a bson.D to preserve order.
	searchStage := bson.D{
		{Key: "index", Value: index},
		{Key: "text", Value: bson.D{
			{Key: "query", Value: value},
			{Key: "path", Value: path},
		}},
		{Key: "sort", Value: sort},
	}

	if paginationToken != "" {
		switch search {
		case 1:
			searchStage = append(searchStage, bson.E{Key: "searchAfter", Value: []interface{}{paginationToken}})
		case 2:
			searchStage = append(searchStage, bson.E{Key: "searchBefore", Value: []interface{}{paginationToken}})
		}
	}

	return bson.D{{Key: "$search", Value: searchStage}}, nil
}

// EqualsPaginationHelper is an implementation of PaginationHelper using the equals operator.
func (t *PaginationHelperMongoImpl) EqualsPaginationHelper(
	index string,
	path string,
	value string,
	sort bson.D,
	search uint8, // 0: no token, 1: searchAfter, 2: searchBefore
	paginationToken string,
) (bson.D, error) {
	// Build the inner equals search stage as a bson.D.
	searchStage := bson.D{
		{Key: "index", Value: index},
		{Key: "equals", Value: bson.D{
			{Key: "value", Value: value},
			{Key: "path", Value: path},
		}},
		{Key: "sort", Value: sort},
	}

	if paginationToken != "" {
		switch search {
		case 1:
			searchStage = append(searchStage, bson.E{Key: "searchAfter", Value: []interface{}{paginationToken}})
		case 2:
			searchStage = append(searchStage, bson.E{Key: "searchBefore", Value: []interface{}{paginationToken}})
		}
	}

	return bson.D{{Key: "$search", Value: searchStage}}, nil
}
