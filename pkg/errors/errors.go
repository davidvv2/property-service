package errors

// ErrUnimplemented : A unimplemented feature will throw this error.
var ErrUnimplemented = New("Unimplemented")

// Business logic: The errors below relate to errors thrown as part of the business logic.
var (
	// ErrInvalidHeader: The header of the request sent to the service is invalid.
	ErrInvalidHeader = New("invalided header")
	// ErrConvertingError: The error passed to errors.As is not a AppError type.
	ErrConvertingError = New("error while converting to app error")

	//ErrInvalidLogin:  The users password is incorrect.
	ErrInvalidLogin = New("incorrect password")
	// ErrInvalidMFACode: The user inputted MFA code is invalid.
	ErrInvalidMFACode = New("invalid mfa code")
	// ErrMinimumPasswordLength: The password is less than the minimum length.
	ErrMinimumPasswordLength = New("password is under the minimum password length")
	// ErrMaximumPasswordLength: The password length is beyond the maximum length.
	ErrMaximumPasswordLength = New("password is over the maximum password length")
	// ErrPasswordComplexity: The password entered does not meet the minimum entropy.
	ErrPasswordComplexity = New("password does not meet the minimum password complexity")
	// ErrUserDisabled: The users account is disabled.
	ErrUserDisabled = New("user Disabled")
	// ErrMaxSession: The session count is maxed out.
	ErrMaxSession = New("max session count")
	// ErrMFAAlreadyEnabled: MFA is already enabled on the account.
	ErrMFAAlreadyEnabled = New("mfa already enabled")
	// ErrEmailAlreadyVerified: The users email address is already verified.
	ErrEmailAlreadyVerified = New("email already verified")

	// ErrUpdatingUser: A error has occurred while updating the user.
	ErrUpdatingUser = New("failed to update user")
	//ErrAddingAttempt: A error has occurred when trying to add a attempt to a users attempts.
	ErrAddingAttempt = New("failed to add attempt")
)

// Permissions: The errors below are related to permissions of the service.
var (
	//ErrPermissionDenied: the user requesting that resource does not have valid permissions to access it.
	ErrPermissionDenied = New("permission denied")
	// ErrEmailUnVerified: The users email is not verified.
	ErrEmailUnVerified = New("email is not verified")
	// ErrTokenMissMatch: The refresh and session token users do not match.
	ErrTokenMissMatch = New("refresh and session do not match")
	// ErrTokenAlgoMissMatch: The token algorithm miss matches the one we use.
	ErrTokenAlgoMissMatch = New("JWT Verification Failed Algo miss match")
	// ErrTokenBlacklisted: The jwt token has been blacklisted.
	ErrTokenBlacklisted = New("token is blacklisted")
	// ErrResetToken: The jwt token is a reset token and tried to be used for authentication.
	ErrResetToken = New("token is reset token")
)

// Factory: The errors below are related to Factory method's.
var (
	// ErrMinSchemaVersion: The schema version passed to the factory is less than the minimum required.
	ErrMinSchemaVersion = New("minimum schema version not met")
	// ErrMaxSchemaVersion: The schema version passed to the factory is greater than the maximum.
	ErrMaxSchemaVersion = New("maximum schema version passed")
	// ErrTimeGreaterThanMax: The time passed to the factory is greater than the maximum time allowed.
	ErrTimeGreaterThanMax = New("time is grater than max period")
	// ErrTimeLessThanMin: The time passed is less than the minimum time allowed.
	ErrTimeLessThanMin = New("time is less than the minimum period")
	// ErrInvalidConfigFactory: The config passed to the factory is invalid.
	ErrInvalidConfigFactory = New("invalid config passed to Factory")
)

/*****************
* Infrastructure *
******************/

// Infrastructure.
var (
	// ErrCollectionNotFound: The database collection was not found.
	ErrCollectionNotFound = New("collection not found")
	// ErrNoUpdate: Nothing was updated.
	ErrNoUpdate = New("nothing was updated")
	//ErrResourceAlreadyExists: The resource already exists and therefore can not be recreated.
	ErrResourceAlreadyExists = New("resource already exists")
	// ErrHealthCheck : The health check has failed.
	ErrHealthCheck = New("failed to initialise health test")

	//ErrDomainToPersistenceConversion: A error has occurred when converting from a domain object to a database object.
	ErrDomainToPersistenceConversion = New("a error has occurred when converting from domain to persistence")
	//ErrPersistenceToDomainConversion: A error has occurred when converting from a database object to a domain object.
	ErrPersistenceToDomainConversion = New("a error has occurred when converting from persistence to domain")
	// ErrDeleteCountNothing: The returned delete count was nothing.
	ErrDeleteCountNothing = New("delete count was 0")
)

// Caching errors are used for the cacher object, this could be either the in memory or the redis impl.
var (
	// ErrGettingHashKey : The hash key has not been found or  the value from that hash key could not be retrieved.
	ErrGettingHashKey = New("failed to get hash key")
	// ErrSettingHashKey : The hash value could not be set.
	ErrSettingHashKey = New("failed to set hash key")
	// ErrFieldNotFound: the field in the hash table has not been found.
	ErrFieldNotFound = New("field not found")

	// ErrGettingKey : the key has not been found or  the value from that key could not be retrieved.
	ErrGettingKey = New("failed to get key")
	// ErrSettingKey : The key could not be set, this is either a network issue or a capacity issue in the cacher.
	ErrSettingKey = New("failed to set key")

	// ErrCheckingKey: A error has occurred while trying to check if a key exists.
	ErrCheckingKey = New("failed to check key")
)

// Generic Errors.
var (
	// ErrTypeCasting : is thrown when you can not cast a type.
	ErrTypeCasting = New("can not cast type")
)

/*****************
*  Internal
*******************/

// Cryptography related errors, please use these for the crypto package.
var (
	/*******
	* Keys *
	********/

	// ErrInvalidEd25519PubKeySize : The inputted ED25519 public key is invalid, it does not have the correct size.
	ErrInvalidEd25519PubKeySize = New("invalid public Ed25519 key size")
	// ErrInvalidEd25519PriKeySize : The inputted ED25519 private key is invalid, it does not have the correct size.
	ErrInvalidEd25519PriKeySize = New("invalid private Ed25519 key size")
	// ErrCastingEd25519PubKey : A error has occurred while casting the public PKCS8 key to ED25519 key object.
	ErrCastingEd25519PubKey = New("failed to cast key to Ed25519 public key")
	// ErrCastingEd25519PriKey :A error has occurred while casting the private PKCS8 key to ED25519 key object.
	ErrCastingEd25519PriKey = New("failed to cast key to Ed25519 private key")
	// ErrParsePEMBlock : A error has occurred while parsing the pem encoded certificate from the inputted string.
	ErrParsePEMBlock = New("failed to parse PEM block")
	// ErrPEMType : The pem type is not as expected, it should be either a public or private key .
	ErrPEMType = New("unexpected PEM block type")
	// ErrParsePubKey : The public key is invalid format.
	ErrParsePubKey = New("failed to parse public key")
	// ErrParsePriKey : The private key is invalid format.
	ErrParsePriKey = New("failed to parse private key")

	/*********
	*  Hash  *
	**********/

	// The format of the inputted encoded hash is not correct.
	ErrHashFormat = New("the encoded hash is not in the correct format")
	// The inputted encoded hash uses a incompatible version of argon2.
	ErrHashVersion = New("incompatible version of argon2")
)
