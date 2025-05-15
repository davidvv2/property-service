package errors

const (
	AccountLocked             = "account locked until "
	JWTPhase                  = "failed to parse JWT: "
	JWTExpired                = "token has expired at: "
	JWTCreation               = "error while making token: "
	JWTStructPhasing          = "a error has occurred when phasing token into struct"
	JWTNotIssued              = "token has not been issued yet"
	HashFailedCheckingVersion = "failed to check hash version:"
	FailedRandomBits          = "failed to generate random bits:"
	FailedLoadKeys            = "failed to load key pair: "
	FailedEd25519KeyGen       = "failed to generate Ed25519KeyPairs key:"
)
