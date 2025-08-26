# JWT Package

This package provides support for issuing and validating JSON Web Tokens (JWT) using Ed25519 signing.

## Package Structure

- `manager.go` &mdash; Defines the `JWTManager` interface for signing and verifying tokens.
- `manager_ed25519_impl.go` &mdash; Implements `JWTManager` using Ed25519 keys for strong, modern cryptographic signatures.
- `claims_authentication.go` &mdash; Defines custom JWT claims (`AuthClaims`) and helper functions for token generation and validation.
- `README.md` &mdash; This documentation file.

## Getting Started

### Installation

```bash
go get property-service/pkg/jwt
```

### Configuration

Provide your Ed25519 private and public keys as byte slices when creating the manager.

```go
import (
  "property-service/pkg/jwt"
)

// Load or generate your Ed25519 key pair
privateKey := loadEd25519PrivateKey()
publicKey := privateKey.Public().(ed25519.PublicKey)

// Create a new JWT manager
manager := jwt.NewEd25519Manager(privateKey, publicKey)
```

### Generating a Token

```go
claims := jwt.AuthClaims{
  Subject:   "user-id-123",
  Issuer:    "your-service",
  ExpiresAt: time.Now().Add(time.Hour),
}
tokenString, err := manager.SignToken(claims)
if err != nil {
  // handle signing error
}
```

### Validating a Token

```go
parsedClaims, err := manager.VerifyToken(tokenString)
if err != nil {
  // invalid or expired token
}
// use parsedClaims.Subject, etc.
```

## Extending

To support a different signature algorithm, implement the `JWTManager` interface in a new file and register it in your application code.

## Testing

Run unit tests for the JWT package:

```bash
cd pkg/jwt
go test -v
```
