// Package errors implements a custom error, that is used to send the appropriate error
// response independent of the port. It provides a consistent way to create errors
// across the service that provides the necessary metadata to send back to the calling service.
// This allows for clean separation and creation of error responses to the client.
package errors

import (
	"property-service/pkg/errors/codes"
)

// AppError helps encapsulate relevant error information in a structured manner.
// It is helpful as it creating consistent error handling and translations.
type AppError struct {
	err         error      // The error that is being thrown.
	translation string     // The translation of that error that is sent back to the user.
	errorType   Type       // The error type.
	code        codes.Code // The error code for the request.
}

// Error will return the error string.
func (a AppError) Error() string {
	return a.err.Error()
}

// Error will return the error string.
func (a AppError) GetError() error {
	return a.err
}

// Translation will return a translation for the error that then can be looked up and returned to the user.
func (a AppError) Translation() string {
	return a.translation
}

// ErrorType will return the type of the error.
func (a AppError) ErrorType() Type {
	return a.errorType
}

// Status will return the status code to return back to the user.
func (a AppError) Status() codes.Code {
	return a.code
}
