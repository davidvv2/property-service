# Properties Service Module

This directory contains the business logic services that orchestrate operations for properties and owners. It bridges the gap between the application (commands/queries) and the underlying persistence (via repositories and factories).

## Overview

- **Repositories:**  
  The module creates concrete repository implementations (e.g., MongoDB-based) that interact with the domain layer. See `repositories.go` for the instantiation of property and owner repositories.

- **Factories:**  
  The module sets up domain factories for properties and owners using configuration and validators. Refer to `factories.go` for how factories are created and injected.

- **Database Operations:**  
  The `database.go` file encapsulates creation of finders, inserters, updaters, and removers for both properties and owners. These components allow the repositories to perform CRUD operations.

- **Commands:**  
  Business logic commands are wired up in `commands.go`, which assembles command handlers for creating, updating, and deleting properties and owners.

- **Application Initialization:**  
  The `application.go` file builds the complete application by integrating dependencies (repositories, factories, logger, JWT, etc.) and exposes both command and query handlers.

## How It Works

1. **Dependency Injection:**  
   The module assembles its dependencies through the `BuildDependencies` function (in `application.go`). This includes setting up repositories, factories, and other supporting services (cache, JWT managers, etc.).

2. **Command & Query Handlers:**  
   Command and query handlers are created using the services in this directory and are then passed to the application layer, enabling CQRS-based operations.

3. **Repository Methods:**  
   The repositories created here are used to interact directly with the MongoDB collections. They are configured via the factories and database connectors.

## Customization

To change configuration or adjust behavior:
- Update the factory configuration in `factories.go`.
- Customize repository behavior in `repositories.go` and `database.go`.
- Adjust command/query wiring in `commands.go` and the application initialization in `application.go`.

## Testing

The module is supported by integration tests (using build tags like `cse`) found in the `app/command` and `app/query` directories. Running these tests helps verify that business workflows are functioning as intended.