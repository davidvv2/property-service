# Decorator

This `decorator` library provides a set of "middleware" for CQRS that is implemented by using a decorator pattern.

The decorator pattern allows behaviors to be added to an individual object, either statically or dynamically, without affecting the behaviour of other objects from the same class.

We use this functionality to provide middleware like actions to every command or query operation. By using this kind of middleware it is independent rather than tied to a specific technology like gRPC or REST requests.

## Purpose

The purpose of the `decorator` library is to do common functionality to every command or query call. It provides a cleaner, more efficient and maintainable method to deploy common functionality to command and query operations.

## Usage

To use the `decorator` library, copy the code into your project under the internal/decorator folder and then use it as a import. The decorator library has the following functions:

### Decorators

---

- `Command`: The command decorator applies the common handler and is the base for the command decorator. It defines the Command Handler interface that is used across all commands. The decorator uses generics to pass the command input.
- `Logging`: This decorator will log all errors and successful commands or queries.
- `Validator`: This decorator will validate the command or query input struct.
- `Query`: The query decorator applies the common handler and is the base for the query decorator. It defines the Query Handler interface that is used across all queries. The decorator uses generics to pass the query input and the output.
