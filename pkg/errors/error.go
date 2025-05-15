package errors

import "errors"

// New will return a new error.
func New(err string) error {
	return errors.New(err)
}

// Join will join n errors together using error wrapping.
func Join(errs ...error) error {
	return errors.Join(errs...)
}

// Compare: Checks to see if two errors are equal.
func Compare(err error, target error) bool {
	return errors.Is(err, target)
}

// As: Checks for the error in the error tree and returns true if found.
// If found it will also set the target to that error in the error tree.
func AsAppError(err error, target *AppError) bool {
	return errors.As(err, target)
}
