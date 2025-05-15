package errors

import (
	"property-service/pkg/errors/codes"
	"property-service/pkg/errors/translation"
)

/***************************************************************************
* This file contains errors relating to the layer design of the code base. *
****************************************************************************/

// NewHandlerError is used when a error has occurred in a CQRS handler. This could be because a function the
// handler is calling has returned a error or a problem has occurred in the handler.
//
// Please use this function to create a constant error for all errors returned by a CQRS handler.
//
// It takes a error, grpc status code and a string as parameters and returns a new AppError. The error passed to the
// function is the error thrown, the code is a error code for the port and the string passed to the function
// is the translation that is sent back and is translated using I18N.
func NewHandlerError(
	err error,
	translation string,
	code codes.Code,
) AppError {
	return AppError{
		err:         err,         // The error that is being thrown.
		translation: translation, // The translation of that error that is sent back to the user.
		errorType:   Handler,     // The error type.
		code:        code,        // The error code for the request.
	}
}

// NewDomainError is used when a error has occurred in the domain layer. This could be because the entity is in a
// invalid state or a business rule returns a error or the domain logic has thrown a error.
//
// Please use this function to create a constant error for all errors returned by a package in the domain layer.
//
// It takes a error, grpc status code and a string as parameters and returns a new AppError. The error passed to the
// function is the error thrown, the code is a error code for the port and the string passed to the function
// is the translation that is sent back and is translated using I18N.
func NewDomainError(
	err error,
	translation string,
	code codes.Code,
) AppError {
	return AppError{
		err:         err,         // The error that is being thrown.
		translation: translation, // The translation of that error that is sent back to the user.
		errorType:   Domain,      // The error type.
		code:        code,        // The error code for the request.
	}
}

// NewRepositoryError is used when a error has occurred in the repository layer. This could be because a
// deadline is exceeded or the underlying database has returned a error or the repository has thrown a error.
//
// Please use this function  to create a constant error for all errors returned by a package in the repository
// layer.
//
// It takes a error, grpc status code and a string as parameters and returns a new AppError. The error passed to the
// function is the error thrown, the code is a error code for the port and the string passed to the function
// is the translation that is sent back and is translated using I18N.
func NewRepositoryError(
	err error,
	translation string,
	code codes.Code,
) AppError {
	return AppError{
		err:         err,         // The error that is being thrown.
		translation: translation, // The translation of that error that is sent back to the user.
		errorType:   Repository,  // The error type.
		code:        code,        // The error code for the request.
	}
}

// NewInfrastructureError is used when a error has occurred in the infrastructure layer. This could be because a
// deadline is exceeded or the a package in the infrastructure layer has returned a error.
//
// Please use this function  to create a constant error for all errors returned by a package in the infrastructure
// layer.
//
// It takes a error, grpc status code and a string as parameters and returns a new AppError. The error passed to the
// function is the error thrown, the code is a error code for the port and the string passed to the function
// is the translation that is sent back and is translated using I18N.
func NewInfrastructureError(
	err error,
	translation string,
	code codes.Code,
) AppError {
	return AppError{
		err:         err,            // The error that is being thrown.
		translation: translation,    // The translation of that error that is sent back to the user.
		errorType:   Infrastructure, // The error type.
		code:        code,           // The error code for the request.
	}
}

// NewClientError is used when a client of the service is used and returns a error. This could be because a deadline
// is exceeded or the service has returned a error.
//
// Please use this function to create a constant error for all errors Returned by a calling client.
//
// It takes a error, grpc status code and a string as parameters and returns a new AppError. The error passed to the
// function is the error thrown, the code is a error code for the port and the string passed to the function
// is the translation that is sent back and is translated using I18N.
func NewClientError(
	err error,
) AppError {
	return AppError{
		err:         err,                            // The error that is being thrown.
		translation: translation.SomethingWentWrong, // The translation of that error that is sent back to the user.
		errorType:   Client,                         // The error type.
		code:        codes.Internal,                 // The error code for the request.
	}
}

// NewInternalError is used when a error has occurred in the internal packages. This could be because of invalid
// data passed to the functions or a bug exists in the code.
//
// Please use this function to create a constant error for all errors Returned by a internal package.
//
// It takes a error parameters and returns a new AppError. The error passed to the
// function is the error thrown. It will use a translation of SomethingWentWrong and use a internal error code.
func NewInternalError(
	err error,
) AppError {
	return AppError{
		err:         err,                            // The error that is being thrown.
		translation: translation.SomethingWentWrong, // The translation of that error that is sent back to the user.
		errorType:   Internal,                       // The error type.
		code:        codes.Internal,                 // The error code for the request.
	}
}

// NewDatabaseError is used when a error has occurred in the database packages. This could be because of invalid
// data passed to the functions or a bug exists in the code.
//
// Please use this function to create a constant error for all errors returned by the database package.
//
// It takes a error parameters and returns a new AppError. The error passed to the
// function is the error thrown. It will use a translation of SomethingWentWrong and use a internal error code.
func NewDatabaseError(
	err error,
) AppError {
	return AppError{
		err:         err,                            // The error that is being thrown.
		translation: translation.SomethingWentWrong, // The translation of that error that is sent back to the user.
		errorType:   Internal,                       // The error type.
		code:        codes.Internal,                 // The error code for the request.
	}
}
