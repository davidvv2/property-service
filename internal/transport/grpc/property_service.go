package grpc

import (
	"context"

	"property-service/api/proto"
	"property-service/internal/properties/app/command"
	"property-service/internal/properties/app/query"
	port "property-service/internal/properties/ports"
	"property-service/pkg/address"

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
	err := s.AppService.CreateProperty(ctx, command.CreatePropertyCommand{
		PropertyID: req.Id,
		OwnerID:    req.OwnerID,
		Address: address.Address{
			FirstLine:  req.Address.FirstLine,
			Street:     req.Address.Street,
			City:       req.Address.City,
			County:     req.Address.County,
			Country:    req.Address.Country,
			PostalCode: req.Address.Postcode,
			GeoJSON: &address.GeoJSONCoordinates{
				Type:        "Point",
				Coordinates: [2]float64{float64(*req.Address.Latitude), float64(*req.Address.Longitude)},
			},
		},
		Description:   req.Description,
		Title:         req.Title,
		Category:      req.Category,
		Available:     req.Available,
		AvailableDate: req.AvailableDate.AsTime(),
		SaleType:      uint8(req.SaleType),
	})
	if err != nil {
		s.AppService.Log.Error("Failed to create property", err)
		return nil, err
	}
	s.AppService.Log.Debug("Property created successfully")
	// Return the response
	return &proto.CreatePropertyResponse{
		Id: req.Id,
	}, nil
}

func (s *MyPropertyService) ReadProperty(ctx context.Context, req *proto.ReadPropertyRequest) (*proto.Property, error) {
	s.AppService.Log.Debug("Reading property with ID:%s", req.Id)
	property, err := s.AppService.GetProperty(ctx, query.GetPropertyQuery{
		ID: req.Id,
	})
	if err != nil {
		s.AppService.Log.Error("Failed to read property", err)
		return nil, err
	}
	s.AppService.Log.Debug("Property read successfully:", property)
	var latitude *float32
	var longitude *float32
	if property.Address.GeoJSON != nil {
		coords := property.Address.GeoJSON.Coordinates
		if len(coords) == 2 {
			// note: GeoJSON.Coordinates is [lng, lat]
			long := float32(coords[0])
			lat := float32(coords[1])
			latitude = &lat
			longitude = &long
		}
	}

	return &proto.Property{
		Id:      property.ID,
		OwnerID: property.OwnerID,
		Address: &proto.Address{
			FirstLine: property.Address.FirstLine,
			Street:    property.Address.Street,
			City:      property.Address.City,
			County:    property.Address.County,
			Country:   property.Address.Country,
			Postcode:  property.Address.PostalCode,
			Latitude:  latitude,
			Longitude: longitude,
		},
		Description:   property.Description,
		Title:         property.Title,
		AvailableDate: timestamppb.New(property.AvailableDate),
		Available:     wrapperspb.Bool(property.Available),
		SaleType:      uint32(property.SaleType),
		Category:      property.Category,
	}, nil
}

func (s *MyPropertyService) UpdateProperty(ctx context.Context, req *proto.UpdatePropertyRequest) (*proto.UpdatePropertyResponse, error) {
	s.AppService.Log.Debug("Updating property with ID:", req.Id)

	err := s.AppService.UpdateProperty(ctx, command.UpdatePropertyCommand{
		PropertyID: req.Id,
		Address: address.Address{
			FirstLine:  req.Address.FirstLine,
			Street:     req.Address.Street,
			City:       req.Address.City,
			County:     req.Address.County,
			Country:    req.Address.Country,
			PostalCode: req.Address.Postcode,
			GeoJSON: &address.GeoJSONCoordinates{
				Type:        "Point",
				Coordinates: [2]float64{float64(*req.Address.Latitude), float64(*req.Address.Longitude)},
			},
		},
		Description:   req.Description,
		Title:         req.Title,
		Category:      req.Category[0],
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

func (s *MyPropertyService) ListPropertyByCategory(ctx context.Context, req *proto.PropertyListByCategoryRequest) (*proto.ListPropertyResponse, error) {
	s.AppService.Log.Debug("Listing properties")
	properties, err := s.AppService.ListPropertiesByCategory(ctx, query.ListPropertiesByCategoryQuery{
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
		var latitude *float32
		var longitude *float32

		if property.Address.GeoJSON != nil {
			lat := float32(property.Address.GeoJSON.Coordinates[0])
			lng := float32(property.Address.GeoJSON.Coordinates[1])
			latitude = &lat
			longitude = &lng
		}
		s.AppService.Log.Debug("Converting property to proto format: %s", property.ID)
		propertyList = append(propertyList, &proto.Property{
			Id:      property.ID,
			OwnerID: property.OwnerID,
			Address: &proto.Address{
				FirstLine: property.Address.FirstLine,
				Street:    property.Address.Street,
				City:      property.Address.City,
				County:    property.Address.County,
				Country:   property.Address.Country,
				Postcode:  property.Address.PostalCode,
				Latitude:  latitude,
				Longitude: longitude,
			},
			Description:     property.Description,
			Title:           property.Title,
			AvailableDate:   timestamppb.New(property.AvailableDate),
			Available:       wrapperspb.Bool(property.Available),
			SaleType:        uint32(property.SaleType),
			Category:        property.Category,
			PaginationToken: property.PaginationToken,
		})
	}
	return &proto.ListPropertyResponse{
		Properties: propertyList,
	}, nil
}

func (s *MyPropertyService) ListPropertyByOwner(ctx context.Context, req *proto.PropertyListByOwnerRequest) (*proto.ListPropertyResponse, error) {
	s.AppService.Log.Debug("Listing properties by owner")
	properties, err := s.AppService.ListPropertiesByOwner(ctx, query.ListPropertiesByOwnerQuery{
		Server:          "Test",
		Owner:           req.OwnerID,
		Sort:            uint8(req.Sort),
		Limit:           uint16(req.Limit),
		PaginationToken: req.PaginationToken,
		Search:          uint8(req.Search),
	})
	if err != nil {
		s.AppService.Log.Error("Failed to list properties by owner", err)
		return nil, err
	}
	s.AppService.Log.Debug("Properties by owner listed successfully")
	var propertyList []*proto.Property
	for _, property := range properties.Properties {
		var latitude *float32
		var longitude *float32

		if property.Address.GeoJSON != nil {
			lat := float32(property.Address.GeoJSON.Coordinates[0])
			lng := float32(property.Address.GeoJSON.Coordinates[1])
			latitude = &lat
			longitude = &lng
		}
		s.AppService.Log.Debug("Converting property to proto format: %s", property.ID)
		propertyList = append(propertyList, &proto.Property{
			Id:      property.ID,
			OwnerID: property.OwnerID,
			Address: &proto.Address{
				FirstLine: property.Address.FirstLine,
				Street:    property.Address.Street,
				City:      property.Address.City,
				County:    property.Address.County,
				Country:   property.Address.Country,
				Postcode:  property.Address.PostalCode,
				Latitude:  latitude,
				Longitude: longitude,
			},
			Description:     property.Description,
			Title:           property.Title,
			AvailableDate:   timestamppb.New(property.AvailableDate),
			Available:       wrapperspb.Bool(property.Available),
			SaleType:        uint32(property.SaleType),
			Category:        property.Category,
			PaginationToken: property.PaginationToken,
		})
	}
	return &proto.ListPropertyResponse{
		Properties: propertyList,
	}, nil
}
