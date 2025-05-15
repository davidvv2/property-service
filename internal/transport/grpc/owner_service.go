package grpc

import (
	"context"
	"property-service/api/proto"
	"property-service/internal/properties/app/command"
	"property-service/internal/properties/app/query"
	port "property-service/internal/properties/ports"
)

// MyOwnerService implements proto.OwnerServiceServer.
type MyOwnerService struct {
	proto.UnimplementedOwnerServiceServer
	AppService *port.ServiceImpl
}

func (s *MyOwnerService) CreateOwner(ctx context.Context, req *proto.CreateOwnerRequest) (*proto.CreateOwnerResponse, error) {
	s.AppService.Log.Debug("Creating new owner")
	err := s.AppService.CreateOwner(ctx, command.CreateOwnerCommand{
		OwnerID:   req.Id,
		Name:      req.Name,
		Email:     req.Email,
		Telephone: req.Telephone,
		Server:    "Test",
	})
	if err != nil {
		s.AppService.Log.Error("Failed to create owner", err)
		return nil, err
	}
	s.AppService.Log.Debug("Owner created successfully")
	// Return the response
	return &proto.CreateOwnerResponse{
		Id: req.Id,
	}, nil
}

func (s *MyOwnerService) ReadOwner(ctx context.Context, req *proto.ReadOwnerRequest) (*proto.ReadOwnerResponse, error) {
	s.AppService.Log.Debug("Reading owner with ID:", req.Id)
	owner, err := s.AppService.GetOwner(ctx, query.GetOwnerQuery{
		ID:     req.Id,
		Server: "Test",
	})
	if err != nil {
		s.AppService.Log.Error("Failed to read owner", err)
		return nil, err
	}
	s.AppService.Log.Debug("Owner read successfully:", owner)
	// Return the response
	return &proto.ReadOwnerResponse{
		Id:        req.Id,
		Name:      owner.Name(),
		Email:     owner.Email(),
		Telephone: owner.Telephone(),
	}, nil
}

func (s *MyOwnerService) UpdateOwner(ctx context.Context, req *proto.UpdateOwnerRequest) (*proto.UpdateOwnerResponse, error) {
	s.AppService.Log.Debug("Updating owner with ID:", req.Id)
	err := s.AppService.UpdateOwner(ctx, command.UpdateOwnerCommand{
		OwnerID:   req.Id,
		Name:      req.Name,
		Email:     req.Email,
		Telephone: req.Telephone,
		Server:    "Test",
	})
	if err != nil {
		s.AppService.Log.Error("Failed to update owner", err)
		return nil, err
	}
	s.AppService.Log.Debug("Owner updated successfully")
	// Return the response
	return &proto.UpdateOwnerResponse{
		Id: req.Id,
	}, nil
}

func (s *MyOwnerService) DeleteOwner(ctx context.Context, req *proto.DeleteOwnerRequest) (*proto.DeleteOwnerResponse, error) {
	s.AppService.Log.Debug("Deleting owner with ID:", req.Id)
	err := s.AppService.DeleteOwner(ctx, command.DeleteOwnerCommand{
		OwnerID: req.Id,
		Server:  "Test",
	})
	if err != nil {
		s.AppService.Log.Error("Failed to delete owner", err)
		return nil, err
	}
	s.AppService.Log.Debug("Owner deleted successfully")
	// Return the response
	return &proto.DeleteOwnerResponse{
		Id: req.Id,
	}, nil
}
