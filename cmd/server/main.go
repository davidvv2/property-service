package main

import (
	"net"

	"property-service/api/proto"
	port "property-service/internal/properties/ports"

	transport "property-service/internal/transport/grpc"
	"property-service/pkg/configs"
	interceptor "property-service/pkg/infrastructure/grpc"
	"property-service/pkg/infrastructure/log"

	_ "google.golang.org/genproto/googleapis/rpc/errdetails"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	// Load environment variables.
	if err := godotenv.Load("dev.env"); err != nil {
		panic("Error loading .env file: " + err.Error())
	}
	// Load configs and logger.
	cfg := configs.New() // initialize your config
	logger := log.NewZapImpl(&cfg.Backend)

	// Initialize the port service.
	portService := port.NewService(cfg, logger)

	// Initialize your gRPC services with the same port service.
	ownerService := &transport.MyOwnerService{
		AppService: portService,
	}
	// Initialize your gRPC service, injecting the port service.
	propService := &transport.MyPropertyService{
		AppService: portService,
	}

	// Start gRPC server.
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		logger.Fatal("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.UnaryErrorInterceptor), // Register your interceptor here.
	)
	proto.RegisterOwnerServiceServer(grpcServer, ownerService)
	proto.RegisterPropertyServiceServer(grpcServer, propService)
	logger.Info("Server started on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		logger.Fatal("failed to serve: %v", err)
	}
}
