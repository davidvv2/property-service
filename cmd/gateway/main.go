package main

import (
	"context"
	"flag"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	_ "google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	proto "property-service/api/proto"
)

func main() {
	// Define command-line flags
	// grpcServerEndpoint := flag.String("grpc-server-endpoint", "property-service.default.svc.cluster.local:50051", "gRPC server endpoint")
	grpcServerEndpoint := flag.String("grpc-server-endpoint", "localhost:50051", "gRPC server endpoint")
	flag.Parse()

	// Create a context that is canceled when the process is terminated.
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Use grpc.Dial to create the connection.
	conn, err := grpc.Dial(*grpcServerEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to dial gRPC server: %v", err)
	}
	defer conn.Close()

	// Create a new ServeMux for the gateway.
	mux := runtime.NewServeMux()

	// after NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	if err := proto.RegisterOwnerServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts); err != nil {
		log.Fatalf("Failed to register HTTP handler: %v", err)
	}
	if err := proto.RegisterPropertyServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts); err != nil {
		log.Fatalf("Failed to register property service HTTP handler: %v", err)
	}
	log.Println("Starting grpc-gateway on :8080")
	if err := http.ListenAndServe("0.0.0.0:8080", mux); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
