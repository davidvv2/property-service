package service

import (
	"os"

	"property-service/pkg/crypto/signing"
	"property-service/pkg/infrastructure/cache"
	"property-service/pkg/infrastructure/log"
	"property-service/pkg/jwt"

	"github.com/go-playground/validator/v10"
)

// jwtManagers holds the necessary jwt creation objects for the application.
type jwtManagers struct {
	authentication jwt.Manager[jwt.AuthClaims]
}

// createJWTManagers : will create and return a the necessary jwt creation objects for the application.
func createJWTManagers(logger log.Logger, cacher cache.Cacher, v *validator.Validate) jwtManagers {
	// Load in the public and private keys.
	keys := signing.MustLoad(
		os.Getenv("ed25519PublicKey"),
		os.Getenv("ed25519PrivateKey"),
	)
	// Create the authentication jwt manager.
	authentication := jwt.NewED25519Manager[jwt.AuthClaims](jwt.InitStruct{
		Issuer:    "PropertyService",
		Subject:   "Login",
		PublicKey: keys.PublicKey, PrivateKey: keys.PrivateKey,
		Cache: cacher, Log: logger,
		V: v,
	})

	// Return a struct of all the jwt objects.
	return jwtManagers{
		authentication: authentication,
	}
}
