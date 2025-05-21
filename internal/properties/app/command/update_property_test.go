//go:build cse
// +build cse

package command_test

import (
	"context"
	"property-service/internal/properties/app/command"
	"property-service/internal/properties/domain/property"
	"property-service/internal/properties/service"
	"property-service/pkg/configs"
	"property-service/pkg/infrastructure/database"
	"property-service/pkg/infrastructure/log"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/suite"
)

// TestUpdatePropertyTestSuite is the test suite for the command package.
type UpdatePropertyTestSuite struct {
	suite.Suite
	ctx        context.Context
	log        log.Logger
	config     configs.Config
	validator  *validator.Validate
	handler    command.UpdatePropertyHandler
	propRepo   property.Repository
	params     command.UpdatePropertyCommand
	ServiceDep service.Dependencies
}

// SetupTest initializes the test suite.
func (s *UpdatePropertyTestSuite) SetupSuite() {
	// Initialize the command handler
	s.handler = command.NewUpdatePropertyHandler(
		s.ServiceDep.Repo.PropertyRepository,
		s.log,
		s.validator,
	)
	s.params = command.UpdatePropertyCommand{
		PropertyID:    database.NewStringID(),
		Address:       "123 Main St",
		Description:   "A beautiful property updated",
		Title:         "Beautiful Property update",
		Category:      "House",
		Available:     func() *bool { b := true; return &b }(),
		AvailableDate: time.Now(),
		SaleType:      2,
		Server:        "Test",
	}
}

func (s *UpdatePropertyTestSuite) SetupTest() {
	// Create a property for testing
	_, err := s.ServiceDep.Repo.PropertyRepository.New(
		s.ctx,
		s.params.Server,
		property.NewPropertyParams{
			PropertyID:    s.params.PropertyID,
			OwnerID:       database.NewStringID(),
			Address:       "123 Main St",
			Description:   "A beautiful property",
			Title:         "Beautiful Property",
			Category:      "House",
			Available:     false,
			AvailableDate: s.params.AvailableDate,
			SaleType:      s.params.SaleType,
		},
	)
	if err != nil {
		s.Fail("Failed to create property for testing", err)
	}
}

// TestUpdatePropertyHandler tests the UpdatePropertyHandler.
func (s *UpdatePropertyTestSuite) TestUpdatePropertyHandler() {
	// Update a new property using the handler
	err := s.handler.Handle(s.ctx, s.params)
	s.NoError(err, "Expected no error when creating a property")

	// Verify that the property was updated successfully
	property, err := s.ServiceDep.Repo.PropertyRepository.Get(s.ctx, s.params.Server, s.params.PropertyID)
	s.NoError(err, "Expected no error when finding the property")
	s.NotNil(property, "Expected property to be found")
	s.Equal(s.params.Title, property.Title(), "Expected property title to match")
	s.Equal(s.params.Description, property.Description(), "Expected property description to match")
}

// func (s *UpdatePropertyTestSuite) TearDownSuite() {
// 	// Clean up the test data
// 	err := s.ServiceDep.Repo.PropertyRepository.Delete(s.ctx, s.params.Server, s.params.PropertyID)
// 	if err != nil {
// 		s.log.Error("Failed to delete property after test", err)
// 	}
// }
