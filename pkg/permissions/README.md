# Permissions

The `permissions` library that provides a convenient way to check a users permissions for endpoints.

## Purpose

The purpose of the `permissions` library is to validate jwt tokens and user permissions.

## Usage

To use the `permissions` library, copy the code into your project under the internal/permissions folder and then use it as a import. The permissions package has the following functionality:

### Functions definition

---

- `CheckPermissions(c context.Context, perm Permissions) (*jwt.AuthClaims, error)`: Checks permissions of a jwt and returns a phased and verified jwt claims.

## Sample Code

Below is a sample code demonstrating how to use the `permissions` library:

```go

//example : function to demonstrate log object
func example(c context.Context, logger log.Logger, service string, route string, authHeader headers.AuthHeader, jwt jwt.Manager) error{
    // Check the permissions for the end point called
    token, err := permissions.CheckPermissions(c, permissions.Permissions{
        Log:        logger,
        JWT:        jwt,
        Header:     authHeader,
        Service:    service,
        Permission: route,
     })

    if err != nil {
        return nil, status.Errorf(codes.PermissionDenied, err.Error())
    }
}

```
