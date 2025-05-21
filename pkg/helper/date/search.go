package date

import (
	"time"

	"property-service/pkg/errors"
)

// DefaultSearchTime : creates a default search start time and end time for a lookup.
func DefaultSearchTime(start string, end string) (time.Time, time.Time, error) {
	startDate, err := time.Parse(time.RFC3339, start)
	if err != nil {
		return time.Now(), time.Now().AddDate(0, -1, 0),
			errors.NewInternalError(err)
	}
	endDate, err := time.Parse(time.RFC3339, end)
	if err != nil {
		return time.Now(), time.Now().AddDate(0, -1, 0),
			errors.NewInternalError(err)
	}
	return startDate, endDate, nil
}
