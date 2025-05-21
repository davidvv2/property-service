# README for Protocol Buffers Definitions

This directory contains the Protocol Buffers definitions for the gRPC service used in this microservice application. The main file is `service.proto`, which defines the service and its methods.

## Generating gRPC Code

To generate the gRPC code from the Protocol Buffers definitions, you need to have the Protocol Buffers compiler (`protoc`) and the Go plugin for `protoc` installed. Follow these steps:

1. Install Protocol Buffers Compiler:
   - Follow the instructions on the [Protocol Buffers GitHub page](https://github.com/protocolbuffers/protobuf) to install `protoc`.

2. Install the Go plugin for `protoc`:
   ```
   go get google.golang.org/protobuf/cmd/protoc-gen-go
   go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
   ```

3. Generate the Go code:
   Navigate to the `api/proto` directory and run the following command:
   ```
   protoc --go_out=. --go-grpc_out=. service.proto
   ```

This will generate the necessary Go files for the gRPC service, which can then be used in your application. 

## Directory Structure

- `service.proto`: Contains the service definition and message types.
- This README.md: Documentation for the Protocol Buffers definitions and code generation instructions.