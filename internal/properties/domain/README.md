# Domain Models

This directory contains the core domain models and business rules for managing properties and owners. The domain layer encapsulates the entities, their validation rules, and factory methods to create or map between domain models and persistence models.

## Directory Structure

```
internal/properties/domain
├── property
│   ├── factory.go           // Factory interface and configuration for properties
│   ├── factory_impl.go      // Concrete factory implementation for properties
│   ├── model.go             // Domain model for a property, with accessor methods
│   └── repository.go        // Repository interface for properties
└── owner
    ├── factory.go           // Factory interface and configuration for owners
    ├── factory_impl.go      // Concrete factory implementation for owners
    ├── model.go             // Domain model for an owner, with accessor methods
    └── repository.go        // Repository interface for owners
```

## Overview

- **Domain Models:**  
  Entities such as Property and Owner, including their business attributes and validation rules.

- **Factories:**  
  Each domain entity has an associated factory (and implementation) that is responsible for creating new instances and mapping between persistence and domain representations.

- **Repositories:**  
  The domain layer defines repository interfaces to abstract data persistence. Implementations for these interfaces are provided in the adapters layer.

## Customization

- Update the factory configurations in `factory.go` or `factory_impl.go` to adjust how entities are instantiated.
- Enhance business rules by modifying the models or adding new domain behavior.
- Ensure consistency between domain validations and external requirements.

## Contributing

When modifying this folder:
- Follow Domain-Driven Design guidelines.
- Update tests in the application layer if domain behavior changes.
- Use meaningful error messages and validate configuration in factories.