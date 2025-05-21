# Header

The `header` package holds grpc header related functions.

## Purpose

The purpose of the `header` package holds grpc header related functions.

### Usage

- `Get(context context.Context) (AuthHeaders, error)`: This will return the auth headers.
- `GetWithoutAuth(context context.Context) (Headers, error)`: Get the headers without authentication token.
