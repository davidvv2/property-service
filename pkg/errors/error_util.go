package errors

import "errors"

func NewSimple(err string) error {
	return errors.New(err)
}

// Join joins multiple errors using error wrapping.
func Join(errs ...error) error {
	return errors.Join(errs...)
}

// Compare checks if two errors are equal.
func Compare(err error, target error) bool {
	return errors.Is(err, target)
}

// AsAppError checks for an AppError in the error tree and assigns it to target if found.
func AsAppError(err error, target *AppError) bool {
	return errors.As(err, target)
}
