# Decorator Package

This package provides decorators that enhance command and query handlers by adding cross-cutting functionalities such as logging and validation. It seamlessly integrates with the command and query packages to ensure that every operation is both validated and logged.

## How It Works

- **Command Decorators:**  
  Wrap command handlers to:
  - Validate incoming request structures.
  - Log execution details.
  
  The `ApplyCommandDecorators` function composes these behaviors and returns a decorated command handler.

- **Query Decorators:**  
  Wrap query handlers to:
  - Validate both the request parameters and the response.
  - Log the process and outcomes of query execution.
  
  The `ApplyQueryDecorators` function applies these decorators, ensuring consistency across all handlers.

## Usage

Integrate decorators during handler initialization. For example:

For commands:
```go
decoratedCommandHandler := decorator.ApplyCommandDecorators(originalCommandHandler, logger, validator)
```

For queries:
```go
decoratedQueryHandler := decorator.ApplyQueryDecorators(originalQueryHandler, logger, validator)
```

## Package Contents

- **command.go:**  
  Defines command handler interfaces and applies the decorators.
  
- **query.go:**  
  Defines query handler interfaces and applies the decorators.
  
- **logging.go:**  
  Implements decorators to log the execution of commands and queries.
  
- **validator.go:**  
  Implements decorators to perform validation using go-playground/validator.

## Benefits

- Centralizes validation and logging, reducing duplication.
- Ensures consistent error handling and debugging.
- Simplifies maintenance by separating cross-cutting concerns.