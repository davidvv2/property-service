package errors

type Type struct {
	t string
}

// //nolint: gochecknoglobals // constant errors types  across the application.
var (
	Handler        = Type{"Handler"}        // The Handler type is used for errors in the CQRS handler.
	Domain         = Type{"Domain"}         // The domain type is used for errors in the Domain layer.
	Repository     = Type{"Repository"}     // The Repository type is used for errors in the Repository layer.
	Infrastructure = Type{"Infrastructure"} // The infrastructure type is used for errors in the infrastructure layer.
	Client         = Type{"Client"}         // The Client type is used for errors returned by clients.
	Internal       = Type{"Internal"}       // The internal type is used for internal package errors.

	Authentication  = Type{"Authentication"}  // The authentication type is used for authentication related issues.
	InvalidArgument = Type{"InvalidArgument"} // The Invalid argument type is used for invalid arguments.
)
