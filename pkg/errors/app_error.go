package errors

import (
	"errors"
	"fmt"

	"property-service/pkg/errors/codes"
)

// Type represents the error layer.
type Type string

// Predefined layers.
const (
	Handler        Type = "Handler"        // The Handler type is used for errors in the CQRS handler.
	Domain         Type = "Domain"         // The domain type is used for errors in the Domain layer.
	Repository     Type = "Repository"     // The Repository type is used for errors in the Repository layer.
	Infrastructure Type = "Infrastructure" // The infrastructure type is used for errors in the infrastructure layer.
	Client         Type = "Client"         // The Client type is used for errors returned by clients.
	Internal       Type = "Internal"       // The internal type is used for internal package errors.
	Database       Type = "Database"       // The database type is used for errors in the database layer.

	Authentication  Type = "Authentication"  // The authentication type is used for authentication related issues.
	InvalidArgument Type = "InvalidArgument" // The Invalid argument type is used for invalid arguments.
)

// AppError encapsulates error metadata using an underlying error,
// an error layer, and an error code.
type AppError struct {
	err       error      // The underlying error.
	errorType Type       // The error layer.
	code      codes.Code // The error code.
}

// Error implements the error interface.
func (a AppError) Error() string {
	return fmt.Sprintf("type:%s, code:%s, message:%s", a.errorType, a.code, a.err.Error())
}

// Unwrap returns the underlying error.
func (a AppError) Unwrap() error {
	return a.err
}

// Code returns the error code.
func (a AppError) Code() codes.Code {
	return a.code
}

// Type returns the error layer/type.
func (a AppError) Type() Type {
	return a.errorType
}

// New creates a new AppError with the provided message.
// This default constructor uses Client as the layer and Internal as the error code.
func New(msg string) AppError {
	return AppError{
		err:       errors.New(msg),
		errorType: Client,
		code:      codes.Internal,
	}
}

// NewCustom creates a new AppError from an existing error, with an explicit layer and error code.
func NewCustom(err error, errorType Type, code codes.Code) AppError {
	return AppError{
		err:       err,
		errorType: errorType,
		code:      code,
	}
}

func NewDomainError(err error, code codes.Code) AppError {
	return NewCustom(err, Domain, code)
}

func NewHandlerError(err error, code codes.Code) AppError {
	return NewCustom(err, Handler, code)
}

func NewRepositoryError(err error, code codes.Code) AppError {
	return NewCustom(err, Repository, code)
}

func NewInfrastructureError(err error, code codes.Code) AppError {
	return NewCustom(err, Infrastructure, code)
}

func NewClientError(err error, code codes.Code) AppError {
	return NewCustom(err, Client, code)
}

func NewInternalError(err error) AppError {
	return NewCustom(err, Internal, codes.Internal)
}

func NewDatabaseError(err error) AppError {
	return NewCustom(err, Database, codes.Internal)
}
func NewAuthenticationError(err error) AppError {
	return NewCustom(err, Authentication, codes.Unauthenticated)
}
func NewInvalidArgumentError(err error) AppError {
	return NewCustom(err, InvalidArgument, codes.InvalidArgument)
}
