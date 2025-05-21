package database

type Conversion[Old, New any] func(Old) (New, error)

type QueryHelper[Query, DateQuery, Options, Filter any] interface {
	QueryHelper(
		query Query,
	) Filter

	DateQueryHelper(
		query DateQuery,
	) (
		Filter, Filter,
	)

	optionsHelper(
		limit, order, skip int64,
		sort string,
		useSkip bool,
	) Options
}
