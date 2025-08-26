# Query Layer

This package contains the query handlers for the Property bounded context. It provides read-only operations to fetch data and supports pagination for list operations.

## Handlers

- **get_owner.go**: Retrieves a single owner by ID.
- **get_property.go**: Retrieves a single property by ID.
- **list_properties_by_category.go**: Lists properties filtered by category with pagination support.
- **list_properties_by_owner.go**: Lists properties owned by a specific owner with pagination support.

## Test Suites

Each handler has a corresponding test file to ensure correct behavior:

- `get_owner_test.go`
- `get_property_test.go`
- `list_properties_by_category_test.go`
- `ListPropertiesByOwnerTestSuite` in `list_properties_by_owner.go`
- `x_query_test.go`: Initializes and runs all query tests under the `cse` build tag.

## Usage

To run the query tests:

```bash
cd internal/properties/app/query
go test -tags=cse -v
```

## Adding a New Query Handler

1. Create a new `<your_query>.go` that implements the appropriate `Query` interface.
2. Write a corresponding test `<your_query>_test.go` in this directory.
3. Add the handler into the application wiring in `internal/properties/service/application.go` if needed.
4. Run `go test -tags=cse -v` to verify.
