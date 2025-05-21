package database

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var _ QueryHelper[
	Query[string],
	TimeQuery[string],
	options.FindOptions,
	bson.M,
] = (*QueryHelperMongoImpl[
	Query[string],
	TimeQuery[string],
	options.FindOptions,
	bson.M,
])(nil)

type QueryHelperMongoImpl[
	query Query[string],
	dateQuery TimeQuery[string],
	Options options.FindOptions,
	Filter bson.M,
] struct {
}

func NewMongoQueryHelper() *QueryHelperMongoImpl[
	Query[string],
	TimeQuery[string],
	options.FindOptions, bson.M,
] {
	return &QueryHelperMongoImpl[
		Query[string], TimeQuery[string],
		options.FindOptions, bson.M,
	]{}
}

func (qhmi *QueryHelperMongoImpl[
	Query, DateQuery, Options, Filter,
]) DateQueryHelper(
	queryData DateQuery,
) (Filter, Filter) {
	primitiveQuery := qhmi.convertTimeQuery(((TimeQuery[string])(queryData)))
	if primitiveQuery.Cursor != primitive.NilObjectID {
		return Filter(bson.M{
			"$gt": primitiveQuery.Cursor,
		}), nil
	}
	return nil, Filter(bson.M{
		"$gt": primitive.NewDateTimeFromTime(primitiveQuery.StartDate),
		"$lt": primitive.NewDateTimeFromTime(primitiveQuery.EndDate),
	})
}

func (qhmi *QueryHelperMongoImpl[
	query, DateQuery, Options, Filter,
]) QueryHelper(
	queryData query,
) Filter {
	var queryModel = qhmi.convertQuery((Query[string])(queryData))
	if queryModel.Cursor != primitive.NilObjectID {
		return Filter(bson.M{
			"$gt": queryModel.Cursor,
		})
	}
	return nil
}

// /nolint: unused // bug, it is used.
func (qhmi *QueryHelperMongoImpl[
	query,
	dateQuery,
	Options,
	Filter],
) optionsHelper(
	limit, order, skip int64, sort string, useSkip bool,
) Options {
	opt := options.Find()
	// set sorting options
	opt.SetLimit(limit)
	opt.SetSort(bson.M{sort: order})
	if !useSkip {
		return (Options)(*opt)
	}
	return (Options)(*opt.SetSkip(skip))
}

func (
	qhmi *QueryHelperMongoImpl[
		query,
		dateQuery,
		Options,
		Filter],
) convertTimeQuery(
	queryData TimeQuery[string],
) TimeQuery[primitive.ObjectID] {
	cursorPrimitive, _ := primitive.ObjectIDFromHex(queryData.Cursor)
	idPrimitive, _ := primitive.ObjectIDFromHex(queryData.ID)
	return TimeQuery[primitive.ObjectID]{
		StartDate: queryData.StartDate,
		EndDate:   queryData.EndDate,
		Server:    queryData.Server,
		Sort:      queryData.Sort,
		Order:     queryData.Order,
		Skip:      queryData.Skip,
		Limit:     queryData.Limit,
		Cursor:    cursorPrimitive,
		ID:        idPrimitive,
	}
}

func (
	qhmi *QueryHelperMongoImpl[
		query,
		dateQuery,
		Options,
		Filter],
) convertQuery(
	queryData Query[string],
) Query[primitive.ObjectID] {
	cursorPrimitive, _ := primitive.ObjectIDFromHex(queryData.Cursor)
	idPrimitive, _ := primitive.ObjectIDFromHex(queryData.ID)
	return Query[primitive.ObjectID]{
		Server: queryData.Server,
		Sort:   queryData.Sort,
		Order:  queryData.Order,
		Skip:   queryData.Skip,
		Limit:  queryData.Limit,
		Cursor: cursorPrimitive,
		ID:     idPrimitive,
	}
}
