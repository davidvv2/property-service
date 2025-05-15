package jwt

import (
	"context"
	"crypto/rsa"
	"time"

	"property-service/pkg/errors"
	"property-service/pkg/infrastructure/cache"
	"property-service/pkg/infrastructure/log"

	"github.com/go-jose/go-jose/v3"
	"github.com/go-jose/go-jose/v3/jwt"
	"github.com/go-playground/validator/v10"
)

// ManagerRS256Impl : This is a jwt object used to sign and create jwt token.
type ManagerRS256Impl[T AuthClaims | VerifyEmail | ResetPassword] struct {
	privateKey rsa.PrivateKey

	jwtSigner jose.Signer

	Cache cache.Cacher

	log log.Logger
	v   *validator.Validate

	issuer  string
	subject string

	PublicKey rsa.PublicKey
}

// InitStruct : Used to initialise the jwt object.
type RS256InitStruct struct {
	PrivateKey rsa.PrivateKey
	Cache      cache.Cacher
	Log        log.Logger
	V          *validator.Validate
	Issuer     string
	Subject    string
	PublicKey  rsa.PublicKey
}

// NewRS256Manager : initialises the jwt class.
func NewRS256Manager[T AuthClaims | VerifyEmail | ResetPassword](a RS256InitStruct) ManagerRS256Impl[T] {
	// Initialise the signing object.
	jwtSigner, err := jose.NewSigner(jose.SigningKey{
		Algorithm: jose.RS256,
		Key:       a.PrivateKey,
	}, nil)
	if err != nil {
		panic(err)
	}

	a.Log.Info("Initialised: %T Manager", *new(T))
	return ManagerRS256Impl[T]{
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
func (obj ManagerRS256Impl[T]) BlackList(uuid string, expire time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), blacklistTimeout)
	defer cancel()
	// Blacklist the token.
	return obj.Cache.KeySet(ctx, uuid, []byte("true"), expire)
}

// CheckBlacklist : Checks token has not been blacklisted.
func (obj ManagerRS256Impl[T]) CheckBlacklist(c context.Context, uuid string) (bool, error) {
	// Checks to see if the token has been blacklisted.
	return obj.Cache.KeyExist(c, uuid)
}

// Sign : Signs a jwt token.
func (obj ManagerRS256Impl[T]) Sign(token T) (string, error) {
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
func (obj ManagerRS256Impl[T]) Verify(jwtString string) (*T, error) {
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
	case joseHeader[0].Algorithm != "RS256":
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
