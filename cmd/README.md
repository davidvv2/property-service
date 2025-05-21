# Command Applications

This directory contains the entry points for the Property Service microservices. These applications define the runtime behavior for the server and gateway components.

## Overview

- **Server**: Implements the gRPC server for property listing operations.
- **Gateway**: Exposes the gRPC services as RESTful HTTP endpoints through a gateway.

## Prerequisites

- Go installed (see [go.mod](../go.mod) for version details).
- Environment variables configured in [dev.env](../dev.env).
- Protobuf artifacts generated via `make gen_proto`.

## Building the Applications

Both applications should be built with the `cse` tag enabled.

Using the Makefile:
```bash
make build_grpc
make build_gateway
```

Alternatively, build directly:
```bash
cd server && go build -tags=cse -o property-service-server main.go
cd gateway && go build -tags=cse -o property-service-gateway main.go
```

## Running the Applications

### Running the gRPC Server
1. Navigate to the `server` directory.
2. Run the server with the `cse` build tag:
   ```bash
   go run -tags=cse main.go
   ```
3. The server listens on the port specified in [dev.env](../dev.env) (default: 8080).

### Running the HTTP Gateway
1. Navigate to the `gateway` directory.
2. Run the gateway with the `cse` build tag:
   ```bash
   go run -tags=cse main.go
   ```
3. The gateway translates HTTP requests into gRPC calls.

## Troubleshooting

- Ensure all environment variables are set in [dev.env](../dev.env).
- Verify that Protobuf artifacts are up-to-date by running `make gen_proto`.
- Review logs for issues related to dependencies such as MongoDB or Redis.

## Additional Resources

- Refer to the [root README](../README.md) for further configuration and deployment details.
- Check the [Makefile](../Makefile) for build and deploy commands.
- Consult the [go.mod](../go.mod) file for dependency versions and compatibility.
