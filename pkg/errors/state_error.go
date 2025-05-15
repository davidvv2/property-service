package errors

import "property-service/pkg/errors/codes"

/************************************************
* This file contains errors relating to states. *
*************************************************/

// NewAuthenticationError is used for unauthorised requests.
//
// If a request is unauthorised to do a action or their exists a issue relating to authentication then please use
// this function to create a constant error.
//
// It takes a error and a string as parameters and returns a new AppError. The error passed to the function
// is the error thrown and the string passed to the function is the translation that is sent back and is translated.
func NewAuthenticationError(
	err error,
	translation string,
) AppError {
	return AppError{
		err:         err,                   // The error that is being thrown.
		translation: translation,           // The translation of that error that is sent back to the user.
		errorType:   Authentication,        // The error type.
		code:        codes.Unauthenticated, // The error code for the request.
	}
}

// NewInvalidArgumentError is used when invalid arguments are passed.
//
// If a request has invalid arguments or their exists a issue relating to Invalid argument then please use this
// function to create a constant error.
//
// It takes a error and a string as parameters and returns a new AppError. The error passed to the function
// is the error thrown and the string passed to the function is the translation that is sent back and is translated.
func NewInvalidArgumentError(
	err error,
	translation string,
) AppError {
	return AppError{
		err:         err,                   // The error that is being thrown.
		translation: translation,           // The translation of that error that is sent back to the user.
		errorType:   InvalidArgument,       // The error type.
		code:        codes.InvalidArgument, // The error code for the request.
	}
}
