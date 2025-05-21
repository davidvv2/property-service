package errors

// ErrUnimplemented : A unimplemented feature will throw this error.
var ErrUnimplemented = NewSimple("Unimplemented")

// Permissions: The errors below are related to permissions of the service.
var (
	//ErrPermissionDenied: the user requesting that resource does not have valid permissions to access it.
	ErrPermissionDenied = NewSimple("permission denied")
	// ErrTokenAlgoMissMatch: The token algorithm miss matches.
	ErrTokenAlgoMissMatch = NewSimple("JWT Verification Failed Algo miss match")
)

// Factory: The errors below are related to Factory method's.
var (
	// ErrMinSchemaVersion: The schema version passed to the factory is less than the minimum required.
	ErrMinSchemaVersion = NewSimple("minimum schema version not met")
	// ErrMaxSchemaVersion: The schema version passed to the factory is greater than the maximum.
	ErrMaxSchemaVersion = NewSimple("maximum schema version passed")
	// ErrTimeGreaterThanMax: The time passed to the factory is greater than the maximum time allowed.
	ErrTimeGreaterThanMax = NewSimple("time is grater than max period")
	// ErrTimeLessThanMin: The time passed is less than the minimum time allowed.
	ErrTimeLessThanMin = NewSimple("time is less than the minimum period")
	// ErrInvalidConfigFactory: The config passed to the factory is invalid.
	ErrInvalidConfigFactory = NewSimple("invalid config passed to Factory")
)

/*****************
* Infrastructure *
******************/

// Infrastructure.
var (
	// ErrCollectionNotFound: The database collection was not found.
	ErrCollectionNotFound = NewSimple("collection not found")
	// ErrNoUpdate: Nothing was updated.
	ErrNoUpdate = NewSimple("nothing was updated")
)

/*****************
*  Cryptography *
*******************/

// Cryptography related errors, please use these for the crypto package.
var (
	/*******
	* Keys *
	********/

	// ErrInvalidEd25519PubKeySize : The inputted ED25519 public key is invalid, it does not have the correct size.
	ErrInvalidEd25519PubKeySize = NewSimple("invalid public Ed25519 key size")
	// ErrInvalidEd25519PriKeySize : The inputted ED25519 private key is invalid, it does not have the correct size.
	ErrInvalidEd25519PriKeySize = NewSimple("invalid private Ed25519 key size")
	// ErrCastingEd25519PubKey : A error has occurred while casting the public PKCS8 key to ED25519 key object.
	ErrCastingEd25519PubKey = NewSimple("failed to cast key to Ed25519 public key")
	// ErrCastingEd25519PriKey :A error has occurred while casting the private PKCS8 key to ED25519 key object.
	ErrCastingEd25519PriKey = NewSimple("failed to cast key to Ed25519 private key")
	// ErrParsePEMBlock : A error has occurred while parsing the pem encoded certificate from the inputted string.
	ErrParsePEMBlock = NewSimple("failed to parse PEM block")
	// ErrPEMType : The pem type is not as expected, it should be either a public or private key .
	ErrPEMType = NewSimple("unexpected PEM block type")
	// ErrParsePubKey : The public key is invalid format.
	ErrParsePubKey = NewSimple("failed to parse public key")
	// ErrParsePriKey : The private key is invalid format.
	ErrParsePriKey = NewSimple("failed to parse private key")
)
