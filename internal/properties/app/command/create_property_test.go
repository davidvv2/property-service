//go:build cse
// +build cse

package command_test

import (
	"context"
	"property-service/internal/properties/app/command"
	"property-service/internal/properties/domain/property"
	"property-service/internal/properties/service"
	"property-service/pkg/address"
	"property-service/pkg/configs"
	"property-service/pkg/infrastructure/database"
	"property-service/pkg/infrastructure/log"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/suite"
)

// TestNewPropertyTestSuite is the test suite for the command package.
type NewPropertyTestSuite struct {
	suite.Suite
	ctx        context.Context
	log        log.Logger
	config     configs.Config
	validator  *validator.Validate
	handler    command.CreatePropertyHandler
	propRepo   property.Repository
	params     command.CreatePropertyCommand
	ServiceDep service.Dependencies
}

// SetupTest initializes the test suite.
func (s *NewPropertyTestSuite) SetupTest() {
	// Initialize the command handler
	s.handler = command.NewCreatePropertyHandler(
		s.ServiceDep.Repo.PropertyRepository,
		s.log,
		s.validator,
	)
	s.params = command.CreatePropertyCommand{
		PropertyID: database.NewStringID(),
		OwnerID:    database.NewStringID(),
		Address: address.Address{
			FirstLine:  "42",
			Street:     "Triq ic-Cangar",
			City:       "Victoria",
			County:     "",
			Country:    "Malta",
			PostalCode: "VCT2162",
		},
		Description:   "A beautiful property",
		Title:         "Beautiful Property",
		Category:      "House",
		Available:     true,
		AvailableDate: time.Now(),
		SaleType:      1,
		Server:        "Test",
	}
}

// TestCreatePropertyHandler tests the CreatePropertyHandler.
func (s *NewPropertyTestSuite) TestCreatePropertyHandler() {
	// Create a new property using the handler
	err := s.handler.Handle(s.ctx, s.params)
	s.NoError(err, "Expected no error when creating a property")

	// Verify that the property was created successfully
	property, err := s.ServiceDep.Repo.PropertyRepository.Get(s.ctx, s.params.Server, s.params.PropertyID)
	s.NoError(err, "Expected no error when finding the property")
	s.NotNil(property, "Expected property to be found")
	s.Equal(s.params.Title, property.Title, "Expected property title to match")
	s.Equal(s.params.Description, property.Description, "Expected property description to match")
}

func (s *NewPropertyTestSuite) TearDownSuite() {
	// Clean up the test data
	// err := s.ServiceDep.Repo.PropertyRepository.Delete(s.ctx, s.params.Server, s.params.PropertyID)
	// if err != nil {
	// 	s.log.Error("Failed to delete property after test", err)
	// }
}
