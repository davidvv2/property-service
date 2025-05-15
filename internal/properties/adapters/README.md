# Adapters for Properties Module

This directory contains concrete implementations of repository adapters that interface with the database (e.g., MongoDB) to persist and retrieve property and owner data.

## Overview
- **Property Repository:**  
  Implements the property.Repository interface using MongoDB.  
- **Owner Repository:**  
  Implements the owner.Repository interface using MongoDB.  
- Additional query helper functions are provided to support complex database operations.

## Directory Structure

```
internal/properties/adapters
├── property_repository_mongo_impl.go  // MongoDB implementation for property repository
├── owner_repository_mongo_impl.go     // MongoDB implementation for owner repository
```

## Customization
- Modify repository implementations as business requirements change.
- Update query helpers as needed while ensuring adherence to domain-driven interfaces.

## Contributing
- Ensure that changes align with the interfaces defined in the ports package.
- Update tests accordingly.