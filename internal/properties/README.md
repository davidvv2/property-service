# Internal Properties Module

This directory contains all the business logic and domain definitions for managing properties and owners. It is organized by Clean Architecture principles with separate subdirectories for adapters, application layer, domain models, ports, and service-specific business logic.

## Directory Structure

```
internal/properties
├── adapters
│   // ...existing repository adapter implementations (e.g., Mongo repositories)...
├── app
│   ├── command
│   │   // Command handlers for property and owner CRUD operations
│   └── query
       // Query handlers for retrieving property and owner data
├── domain
│   ├── property
│   │   // Domain model for properties including interfaces and factory implementations
│   └── owner
       // Domain model for owners including interfaces and factory implementations
├── ports
   // Abstractions for interaction with external layers (e.g., repository interfaces)
└── service
    // Business logic services for processing domain-specific operations
```

## Overview

- **Adapters:**  
  Implements the concrete repository interfaces for persistence (using, for example, MongoDB). This layer converts between domain models and database representations.

- **App:**  
  Contains the application use cases.  
  - **Command:** Implements command handlers that process create, update, and delete operations for both properties and owners.  
  - **Query:** Implements query handlers that encapsulate the logic to retrieve and present domain data.

- **Domain:**  
  Defines the core domain models and business rules.  
  - **Property:** Contains the models, interfaces, and factory methods for property entities.  
  - **Owner:** Contains the models, interfaces, and factory methods for owner entities.

- **Ports:**  
  Exposes the interfaces for the infrastructure layer to interact with the domain. These abstractions allow for flexibility in adapting different data sources.

- **Service:**  
  Provides business logic services that orchestrate multiple use cases and help coordinate between commands, queries, and data adapters.

## Getting Started

To leverage the properties module:
1. Use the factories provided in the domain layer to create domain entities.
2. Use the command and query handlers in the app layer to perform operations.
3. Ensure that repository implementations in the adapters layer are properly configured and injected via the ports.
4. Use the services subdirectory for coordinating complex business workflows across multiple domain entities.

## Testing

The module includes comprehensive tests within the `app/command` and `app/query` subdirectories. Make sure to run integration tests using the provided build tags (e.g., `cse`) to verify correct behavior.

## Contributing

Contributions to this module should follow Domain-Driven Design practices and aim to keep a clear separation between business logic, infrastructure, and presentation layers. Please ensure that new functionality is accompanied by corresponding tests.