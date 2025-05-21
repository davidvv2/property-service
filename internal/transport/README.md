# Transport Layer

The transport directory provides the networking interface for the Property Service. It contains a single grpc directory that handles and translates both incoming HTTP and gRPC requests into queries and commands executed via the ports package.

## Directory Structure

```
internal/transport
└── grpc
    // gRPC server implementations, interceptors, service registrations,
    // and logic for translating HTTP requests into internal queries and commands.
```

## Integration

- The grpc layer converts external requests into the internal models used by the ports package.
- Configuration and dependency injection ensure that all request translations and service invocations are properly wired.

## Getting Started

1. Configure server settings (ports, middleware, etc.) via the environment configuration.
2. Ensure that the grpc components integrate with the service initialization provided by the `internal/properties` module.
3. Run the service via the main command applications in `cmd/server` or `cmd/gateway`.