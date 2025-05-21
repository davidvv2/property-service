package signing

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"

	"property-service/pkg/errors"
)

// Ed25519KeyPair represents a public/private key pair for Ed25519.
type Ed25519KeyPair struct {
	PublicKey  ed25519.PublicKey
	PrivateKey ed25519.PrivateKey
}

// GenerateEd25519KeyPair generates a new Ed25519 key pair.
func GenerateEd25519KeyPair() (Ed25519KeyPair, error) {
	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return Ed25519KeyPair{
				PublicKey:  nil,
				PrivateKey: nil,
			}, errors.NewInternalError(
				errors.New(errors.FailedEd25519KeyGen + err.Error()),
			)
	}

	return Ed25519KeyPair{
		PublicKey:  publicKey,
		PrivateKey: privateKey,
	}, nil
}

// MustLoad loads an Ed25519 key pair from the given PEM-encoded strings
// and panics if there is an error.
func MustLoad(public string, private string) Ed25519KeyPair {
	key, err := Load(public, private)
	if err != nil {
		panic(errors.FailedLoadKeys + err.Error())
	}
	return key
}

// Load loads an Ed25519 key pair from the given PEM-encoded strings.
func Load(public string, private string) (Ed25519KeyPair, error) {
	priKey, priKeyErr := loadPrivateKey(private)
	pubKey, pubKeyErr := loadPublicKey(public)
	if priKeyErr != nil || pubKeyErr != nil {
		return Ed25519KeyPair{
			PublicKey:  nil,
			PrivateKey: nil,
		}, errors.NewInternalError(errors.Join(priKeyErr, pubKeyErr))
	}
	return Ed25519KeyPair{
		PublicKey:  pubKey,
		PrivateKey: priKey,
	}, nil
}

// loadPrivateKey loads an Ed25519 private key from the given PEM-encoded string.
func loadPrivateKey(pemEncoded string) (ed25519.PrivateKey, error) {
	block, _ := pem.Decode([]byte(pemEncoded))
	switch {
	case block == nil:
		return nil, errors.ErrParsePEMBlock
	case block.Type != "PRIVATE KEY":
		return nil, errors.ErrPEMType
	}

	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, errors.ErrParsePriKey
	}

	privateKey, ok := key.(ed25519.PrivateKey)
	switch {
	case !ok:
		return nil, errors.ErrCastingEd25519PubKey
	case len(privateKey) != ed25519.PrivateKeySize:
		return nil, errors.ErrInvalidEd25519PriKeySize
	}
	return privateKey, nil
}

func loadPublicKey(pemEncodedPub string) (ed25519.PublicKey, error) {
	blockPub, _ := pem.Decode([]byte(pemEncodedPub))
	switch {
	case blockPub == nil:
		return nil, errors.ErrParsePEMBlock
	case blockPub.Type != "PUBLIC KEY":
		return nil, errors.ErrPEMType
	}

	key, err := x509.ParsePKIXPublicKey(blockPub.Bytes)
	if err != nil {
		return nil, errors.ErrParsePubKey
	}

	publicKey, ok := key.(ed25519.PublicKey)
	switch {
	case !ok:
		return nil, errors.ErrCastingEd25519PubKey
	case len(publicKey) != ed25519.PublicKeySize:
		return nil, errors.ErrInvalidEd25519PubKeySize
	}

	return publicKey, nil
}

// SignMessage signs a message using the private key of a Ed25519KeyPair.
func (kp Ed25519KeyPair) SignMessage(message []byte) []byte {
	return ed25519.Sign(kp.PrivateKey, message)
}

// VerifySignature verifies a signature using the public key of a Ed25519KeyPair.
func (kp Ed25519KeyPair) VerifySignature(message, signature []byte) bool {
	return ed25519.Verify(kp.PublicKey, message, signature)
}
