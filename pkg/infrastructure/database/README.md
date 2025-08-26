# Database Infrastructure Package

This package provides generic database abstractions and MongoDB implementations for common data operations within this project. It centralizes connectors, sessions, CRUD helpers, pagination, and utility functions to simplify interaction with the database.

## Package Structure

- **connector.go** / **connector_mongo_impl.go**
  - Define and implement a `Connector` for establishing MongoDB client connections.
- **session.go** / **session_mongo_impl.go**
  - Expose database session interfaces and Mongo-specific behavior.
- **creator.go** / **creator_mongo_impl.go**
  - Helpers to insert new documents into collections.
- **finder.go** / **finder_mongo_impl.go**
  - Query helpers for retrieving single documents.
- **iterator.go** / **iterator_mongo_impl.go**
  - Cursor-based iteration over query results.
- **inserter.go** / **inserter_mongo_impl.go**
  - Bulk insert operations.
- **updater.go** / **updater_mongo_impl.go**
  - Update helpers for modifying documents by filter or ID.
- **remover.go** / **remover_mongo_impl.go**
  - Delete helpers for removing documents.
- **grouper.go** / **grouper_monog_impl.go**
  - Aggregation and grouping support.
- **composit.go** / **composit_mongo_impl.go**
  - Composite operations combining multiple steps.
- **encrypter.go** / **encrypter_mongo_impl.go**, **encrypter_operater.go**, etc.
  - Client-side encryption utilities for sensitive fields.
- **pagination_helper.go** / **pagination_helper_impl.go**
  - Cursor‐based pagination support for MongoDB Atlas Search or simple filters.
- **query_model.go**
  - Helper to convert generic queries into BSON models.
- **type_conversions.go**
  - Converters between Go types (UUID, time, etc.) and BSON-friendly formats.

## Getting Started

1. **Initialize Connector**

```go
import (
  "context"
  "property-service/pkg/infrastructure/database"
)

cfg := yourConfig() // load Mongo URI, options
conn, err := database.NewMongoConnector(ctx, cfg.Mongo)
if err != nil {
  // handle error
}
```

2. **Create a Session**

```go
session := conn.Session()
defer session.Close()
```

3. **Use CRUD Helpers**

```go
// Insert a document
err = session.Creator("collectionName").Insert(ctx, document)

// Find one
err = session.Finder("collectionName").FindOne(ctx, filter, &result)

// Update by ID
err = session.Updater("collectionName").UpdateOneByID(ctx, id, update)

// Remove documents
err = session.Remover("collectionName").Delete(ctx, filter)
```

4. **Advanced Operations**

- Use `Iterator` for streaming through large result sets.
- Use `Grouper` for aggregation pipelines.
- Use `PaginationHelper` for cursor-based paging in Atlas Search.
- Use `Encrypter`, `Composite`, and type converters as needed.

## Extending and Testing

- To support a new database provider, implement the same interfaces in a new subpackage.
- All MongoDB helpers have corresponding tests in their directories. Run:

```bash
cd pkg/infrastructure/database
go test -v
```
