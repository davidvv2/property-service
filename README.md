# Property Service

This is a microservice for handling property listings.

## Overview

The project is written in Go and is organized into several key directories:

- **api/proto/** – Contains `.proto` files defining the gRPC and HTTP APIs.
- **cmd/server/** – Contains the main application for the gRPC server. See [cmd/server/main.go](cmd/server/main.go).
- **cmd/gateway/** – Contains the main application for the HTTP gateway. See [cmd/gateway/main.go](cmd/gateway/main.go).
- **internal/** – Houses the core business logic and service implementations.
- **pkg/** – Contains shared utilities and helper packages.

## Environment Configuration

Application configuration is managed via environment variables defined in [dev.env](dev.env). This file contains settings for:

- Backend configuration
- Database (MongoDB) access
- Google Cloud and caching (Redis) configuration
- Emailing and JWT configuration

## Build & Deployment

A [Makefile](Makefile) is provided to simplify common tasks:

- **Build Images:**  
  - Run `make build_grpc` to build the gRPC server Docker image.
  - Run `make build_gateway` to build the HTTP gateway Docker image.
  
- **Push Images:**  
  - Run `make push_grpc` and `make push_gateway` to push the images to your container registry.

- **Deploy to Kubernetes:**  
  - Run `make deploy` to apply deployments, services, and configmaps.
  - Other targets such as `restart`, `delete`, and `start_server` aid in managing the deployments.

- **Generate Protobuf Artifacts:**  
  - Run `make gen_proto` to generate the Go code from the Protobuf definitions.

## Development Tips

- Use the configurations in [dev.env](dev.env) for local development.
- Follow the settings in [`.vscode/settings.json`](.vscode/settings.json) for testing and linting.
- Refer to the [go.mod](go.mod) file for dependency management and version information.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.