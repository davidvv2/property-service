package jwt

import (
	"context"
	"crypto/ed25519"
	"time"

	"property-service/pkg/errors"
	"property-service/pkg/infrastructure/cache"
	"property-service/pkg/infrastructure/log"

	"github.com/go-jose/go-jose/v3"
	"github.com/go-jose/go-jose/v3/jwt"
	"github.com/go-playground/validator/v10"
)

const (
	RefreshTime      = 2629744 * time.Second
	SessionTime      = 1200 * time.Second
	blacklistTimeout = 5 * time.Second
	signingAlgo      = "EdDSA"
)

// ManagerED25519Impl : This is a jwt object used to sign and create jwt token.
type ManagerED25519Impl[T AuthClaims] struct {
	v         *validator.Validate
	jwtSigner jose.Signer

	Cache cache.Cacher
	log   log.Logger

	issuer  string
	subject string

	PublicKey  ed25519.PublicKey
	privateKey ed25519.PrivateKey
}

// InitStruct : Used to initialise the jwt object.
type InitStruct struct {
	V          *validator.Validate
	Cache      cache.Cacher
	Log        log.Logger
	Issuer     string
	Subject    string
	PublicKey  ed25519.PublicKey
	PrivateKey ed25519.PrivateKey
}

// NewED25519Manager : initialises the jwt class.
func NewED25519Manager[T AuthClaims](a InitStruct) ManagerED25519Impl[T] {
	// Initialise the signing object.
	jwtSigner, err := jose.NewSigner(jose.SigningKey{
		Algorithm: jose.EdDSA,
		Key:       a.PrivateKey,
	}, nil)
	if err != nil {
		panic(err)
	}

	a.Log.Info("Initialised: %T Manager", *new(T))
	return ManagerED25519Impl[T]{
		v:          a.V,
		jwtSigner:  jwtSigner,
		issuer:     a.Issuer,
		subject:    a.Subject,
		PublicKey:  a.PublicKey,
		privateKey: a.PrivateKey,
		log:        a.Log,
		Cache:      a.Cache,
	}
}

// BlackList :Function to blacklist a jwt token in redis.
func (obj ManagerED25519Impl[T]) BlackList(uuid string, expire time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), blacklistTimeout)
	defer cancel()
	// Blacklist the token.
	return obj.Cache.KeySet(ctx, uuid, []byte("true"), expire)
}

// CheckBlacklist : Checks token has not been blacklisted.
func (obj ManagerED25519Impl[T]) CheckBlacklist(c context.Context, uuid string) (bool, error) {
	// Checks to see if the token has been blacklisted.
	return obj.Cache.KeyExist(c, uuid)
}

// Sign : Signs a jwt token.
func (obj ManagerED25519Impl[T]) Sign(token T) (string, error) {
	validationErr := obj.v.Struct(token)
	if validationErr != nil {
		return "", errors.NewInternalError(validationErr)
	}

	return jwt.Signed(obj.jwtSigner).
		Claims(token).
		//nolint: exhaustruct // std claims are defined here and non-std are passed
		Claims(jwt.Claims{
			Issuer:  obj.issuer,
			Subject: obj.subject,
		}).CompactSerialize()
}

// Verify : verifies a jwt token with a given public key.
func (obj ManagerED25519Impl[T]) Verify(jwtString string) (*T, error) {
	var token *T
	// Phase token as a string map
	parsedJWT, errorPhasing := jwt.ParseSigned(jwtString)
	if errorPhasing != nil {
		return nil, errors.New(errors.JWTPhase + errorPhasing.Error())
	}

	// Get jwt header to check algo
	joseHeader := parsedJWT.Headers

	currentTime := time.Now() // Get the current time
	var out jwt.Claims
	// Phase jwt claims
	tokenError := parsedJWT.Claims(obj.PublicKey, &token, &out)

	// Verify the struct with a validator.
	verifyTokenErr := obj.v.Struct(token)
	var err error
	// Switch statement to check and return any errors
	switch {
	// Check to see if their is a error when phasing the token
	case tokenError != nil:
		err = errors.New(errors.JWTStructPhasing + tokenError.Error())
	// Verifies a token to see if it matches the validator.
	case verifyTokenErr != nil:
		err = verifyTokenErr
	// Checks to see if the token is signed by the correct algo
	case joseHeader[0].Algorithm != signingAlgo:
		err = errors.ErrTokenAlgoMissMatch

	// Checks to see if the token has expired
	case currentTime.After(out.Expiry.Time()):
		err = errors.New(errors.JWTExpired + out.Expiry.Time().String() + " current time: " + time.Now().String())

	// Checks to see if the
	case currentTime.Before(out.NotBefore.Time()):
		err = errors.New(errors.JWTNotIssued + out.NotBefore.Time().String())
	}

	if err != nil {
		return nil, errors.NewInternalError(err)
	}

	// return the token if no errors
	return token, nil
}
