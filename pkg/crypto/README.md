# Crypto Package

This package provides cryptographic functionalities essential to the property-service project. In particular, the signing package focuses on generating, loading, and verifying Ed25519 key pairs.

## Signing Package

- Generates Ed25519 key pairs securely.
- Loads PEM-encoded public and private keys and validates them.
- Signs messages and verifies signatures.
- Used by the JWT module to securely sign and verify authentication tokens. For example, the JWT manager loads keys with `signing.MustLoad` when initializing.

## Usage Example

```go
// Example usage:
keys := signing.MustLoad(ed25519PublicKeyPEM, ed25519PrivateKeyPEM)
signedMessage := keys.SignMessage([]byte("message to sign"))
isValid := keys.VerifySignature([]byte("message to sign"), signedMessage)
```

## Error Handling

Any errors during key generation, loading, or conversion are wrapped in custom error types for consistent error management.