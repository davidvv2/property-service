package translation

const (
	// Unimplemented: is used when the functionality has not been implemented yet.
	Unimplemented = "Unimplemented"

	//SomethingWentWrong: is a catch all term to use when a internal error has occurred.
	SomethingWentWrong = "SomethingWentWrong"

	// PermissionDenied: is used when the user tries to access a resource they do not have permission for.
	PermissionDenied = "PermissionDenied"

	// UnAuthenticated: Please use this whenever the user is UnAuthenticated.
	UnAuthenticated = "UnAuthenticated"

	// InvalidArgument: The arguments sent to the service are invalid.
	InvalidArgument = "InvalidArgument"

	// InvalidJWT : The jwt token passed is invalid.
	InvalidJWT = "InvalidJWT"
)
