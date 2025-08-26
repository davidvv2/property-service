# Command Layer

This package contains the command handlers for the Property bounded context. It provides write operations to create, update, and delete domain entities (Property and Owner).

## Handlers

- **create_owner.go**: Handles creation of a new owner.
- **create_property.go**: Handles creation of a new property.
- **update_owner.go**: Handles updates to an existing owner.
- **update_property.go**: Handles updates to an existing property.
- **delete_owner.go**: Handles deletion of an owner.
- **delete_property.go**: Handles deletion of a property.

## Test Suites

Each handler has a corresponding test file to ensure correct behavior:

- `create_owner_test.go`
- `create_property_test.go`
- `update_owner_test.go`
- `update_property_test.go`
- `delete_owner_test.go`
- `delete_property_test.go`
- `x_command_test.go`: Initializes and runs all command tests under the `cse` build tag.

## Usage

To run the command tests:

```bash
cd internal/properties/app/command
go test -tags=cse -v
```

## Adding a New Command Handler

1. Create a new `<your_command>.go` that implements the appropriate `Command` interface.
2. Write a corresponding test `<your_command>_test.go` in this directory.
3. Register the new command in the application wiring in `internal/properties/service/application.go` if appropriate.
4. Run `go test -tags=cse -v` to verify.
