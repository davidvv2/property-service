# Application Layer

The application layer encapsulates the business use cases for the properties module. It follows Clean Architecture and CQRS principles by separating command and query responsibilities.

## Directory Structure

```
internal/properties/app
├── command
│   // Command handlers for create, update, and delete operations
└── query
    // Query handlers for retrieving property and owner data
```

## Overview

- **Command Handlers:**  
  Handle state-changing operations (such as creating, updating, and deleting properties and owners) by applying validations and business rules before interacting with the domain and repository layers.

- **Query Handlers:**  
  Retrieve data from the domain layer, transform persistence models into domain responses, and return read-only information.

## Usage

- Instantiate command and query handlers via dependency injection.
- Use the command handlers to perform write operations and the query handlers to fetch data.
- Ensure proper configuration of repositories and factories through the dependency injection mechanism in the higher-level application initialization.

## Testing

- Integration tests for command and query handlers are provided with build tags (e.g., `cse`). Refer to the test files in the respective subdirectories for examples.

## Contributing

- Follow CQRS and Clean Architecture practices when adding new commands or queries.
- Ensure new handlers are covered by tests and that the code remains loosely coupled.