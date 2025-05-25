package grpc

import (
	"context"

	"property-service/api/proto"
	"property-service/internal/properties/app/command"
	"property-service/internal/properties/app/query"
	port "property-service/internal/properties/ports"
	"property-service/pkg/infrastructure/database"

	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// MyPropertyService implements proto.PropertyServiceServer.
type MyPropertyService struct {
	proto.UnimplementedPropertyServiceServer
	AppService *port.ServiceImpl
}

func (s *MyPropertyService) CreateProperty(ctx context.Context, req *proto.CreatePropertyRequest) (*proto.CreatePropertyResponse, error) {
	s.AppService.Log.Debug("Creating new property")
	newID := database.NewStringID()
	err := s.AppService.CreateProperty(ctx, command.CreatePropertyCommand{
		PropertyID:    newID,
		OwnerID:       req.OwnerID,
		Address:       req.Address,
		Description:   req.Description,
		Title:         req.Title,
		Category:      req.Category[0],
		Available:     req.Available,
		AvailableDate: req.AvailableDate.AsTime(),
		SaleType:      uint8(req.SaleType),
		Server:        "Test",
	})
	if err != nil {
		s.AppService.Log.Error("Failed to create property", err)
		return nil, err
	}
	s.AppService.Log.Debug("Property created successfully")
	// Return the response
	return &proto.CreatePropertyResponse{
		Id: newID,
	}, nil
}

func (s *MyPropertyService) ReadProperty(ctx context.Context, req *proto.ReadPropertyRequest) (*proto.ReadPropertyResponse, error) {
	s.AppService.Log.Debug("Reading property with ID:", req.Id)
	property, err := s.AppService.GetProperty(ctx, query.GetPropertyQuery{
		ID:     req.Id,
		Server: "Test",
	})
	if err != nil {
		s.AppService.Log.Error("Failed to read property", err)
		return nil, err
	}
	s.AppService.Log.Debug("Property read successfully:", property)
	// Return the response
	return &proto.ReadPropertyResponse{
		Id:            req.Id,
		OwnerID:       property.OwnerID,
		Address:       property.Address,
		Description:   property.Description,
		Title:         property.Title,
		AvailableDate: timestamppb.New(property.AvailableDate),
		Available:     wrapperspb.Bool(property.Available),
		SaleType:      uint32(property.SaleType),
		Category:      []string{property.Category},
	}, nil
}

func (s *MyPropertyService) UpdateProperty(ctx context.Context, req *proto.UpdatePropertyRequest) (*proto.UpdatePropertyResponse, error) {
	s.AppService.Log.Debug("Updating property with ID:", req.Id)
	err := s.AppService.UpdateProperty(ctx, command.UpdatePropertyCommand{
		PropertyID:  req.Id,
		Address:     req.Address,
		Description: req.Description,
		Title:       req.Title,
		Category:    req.Category[0],
		Available: func() *bool {
			if req.Available != nil {
				b := req.Available.Value
				return &b
			}
			return nil
		}(),
		AvailableDate: req.AvailableDate.AsTime(),
		SaleType:      uint8(req.SaleType),
		Server:        "Test",
	})
	if err != nil {
		s.AppService.Log.Error("Failed to update property", err)
		return nil, err
	}
	s.AppService.Log.Debug("Property updated successfully")
	// Return the response
	return &proto.UpdatePropertyResponse{
		Id: req.Id,
	}, nil
}

func (s *MyPropertyService) DeleteProperty(ctx context.Context, req *proto.DeletePropertyRequest) (*proto.DeletePropertyResponse, error) {
	s.AppService.Log.Debug("Deleting property with ID:", req.Id)
	err := s.AppService.DeleteProperty(ctx, command.DeletePropertyCommand{
		PropertyID: req.Id,
		Server:     "Test",
	})
	if err != nil {
		s.AppService.Log.Error("Failed to delete property", err)
		return nil, err
	}
	s.AppService.Log.Debug("Property deleted successfully")
	// Return the response
	return &proto.DeletePropertyResponse{
		Id: req.Id,
	}, nil
}

func (s *MyPropertyService) ListPropertyByCategory(ctx context.Context, req *proto.PropertyListByCategoryRequest) (*proto.ListPropertyByCategoryResponse, error) {
	s.AppService.Log.Debug("Listing properties")
	properties, err := s.AppService.ListPropertiesByCategory(ctx, query.ListPropertiesByCategoryQuery{
		Server:          "Test",
		Category:        req.Category,
		Sort:            uint8(req.Sort),
		Limit:           uint16(req.Limit),
		PaginationToken: req.PaginationToken,
		Search:          uint8(req.Search),
	})
	if err != nil {
		s.AppService.Log.Error("Failed to list properties", err)
		return nil, err
	}
	s.AppService.Log.Debug("Properties listed successfully")
	// Convert properties to proto format
	var propertyList []*proto.Property
	for _, property := range properties.Properties {
		s.AppService.Log.Debug("Converting property to proto format: %s", property.ID)
		propertyList = append(propertyList, &proto.Property{
			Id:            property.ID,
			OwnerID:       property.OwnerID,
			Address:       property.Address,
			Description:   property.Description,
			Title:         property.Title,
			AvailableDate: timestamppb.New(property.AvailableDate),
			Available:     wrapperspb.Bool(property.Available),
			SaleType:      uint32(property.SaleType),
			Category:      []string{property.Category},
		})
	}
	return &proto.ListPropertyByCategoryResponse{
		Properties: propertyList,
	}, nil
}
