package database

import "time"

type TimeQuery[ID any] struct {
	ID     ID `validate:"omitempty,len=24"`
	Cursor ID `validate:"omitempty,len=24"`

	StartDate time.Time `validate:"omitempty,datetime=2006-01-02"`
	EndDate   time.Time `validate:"omitempty,datetime=2006-01-02"`

	Server string `validate:"required"`
	Sort   string `validate:"omitempty,oneof=1 -1"`
	Order  int64  `validate:"omitempty"`
	Skip   int64  `validate:"omitempty,min=0"`
	Limit  int64  `validate:"omitempty,min=0,max=10000"`
}

// Query : meta struct passed to get user list.
type Query[ID any] struct {
	ID     ID     `validate:"omitempty,len=24"`
	Server string `validate:"required"`
	Cursor ID     `validate:"omitempty,len=24"`

	Sort  string `validate:"omitempty,oneof=1 -1"`
	Order int64  `validate:"omitempty"`
	Skip  int64  `validate:"omitempty,min=0"`
	Limit int64  `validate:"omitempty,min=0,max=10000"`
}

func MapQuery[Old any, New any](
	mappingFunc func(Old) (New, error),
	query Query[Old],
) (*Query[New], error) {
	// Map IDs
	queryID, err := mappingFunc(query.ID)
	if err != nil {
		return nil, err
	}
	cursorID, err := mappingFunc(query.Cursor)
	if err != nil {
		return nil, err
	}

	return &Query[New]{
		ID:     queryID,
		Server: query.Server,
		Cursor: cursorID,
		Sort:   query.Sort,
		Order:  query.Order,
		Skip:   query.Skip,
		Limit:  query.Limit,
	}, err
}

func MapTimeQuery[Old any, New any](
	mappingFunc func(Old) (New, error),
	query TimeQuery[Old],
) (*TimeQuery[New], error) {
	// Map IDs
	historyID, err := mappingFunc(query.ID)
	if err != nil {
		return nil, err
	}
	cursorID, err := mappingFunc(query.Cursor)
	if err != nil {
		return nil, err
	}

	return &TimeQuery[New]{
		ID:        historyID,
		Server:    query.Server,
		Cursor:    cursorID,
		StartDate: query.StartDate,
		EndDate:   query.EndDate,
		Sort:      query.Sort,
		Order:     query.Order,
		Skip:      query.Skip,
		Limit:     query.Limit,
	}, err
}
