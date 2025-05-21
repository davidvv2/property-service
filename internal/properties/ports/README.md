# Properties Ports Package

This directory defines the abstractions ("ports") for the properties module. These interfaces enable loose coupling between the domain/application layers and the underlying infrastructure (e.g., persistence, external services).

## Overview

- **Purpose:**  
  Provides interfaces that define the contracts for repository operations and other external interactions. This allows the application layer to remain independent of specific implementations.

- **Key Interfaces:**  
  - Repository interfaces for properties and owners  
  - Any other abstractions required for interaction with external layers

## Usage

- Implement these interfaces in the adapters layer.  
- Inject the ports into application services and command/query handlers to facilitate dependency inversion and easier testing.

## Directory Structure

```
internal/properties
└── ports
    └── ...existing code...
```

## Customization

- Update or extend the interfaces in this package as new requirements arise.  
- Ensure consistent adherence to domain boundaries when modifying these abstractions.

## Contributing

When contributing changes:
- Maintain clear and concise interfaces.
- Update accompanying tests as needed.