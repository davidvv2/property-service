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

// TestDeletePropertyTestSuite is the test suite for the command package.
type DeletePropertyTestSuite struct {
	suite.Suite
	ctx        context.Context
	log        log.Logger
	config     configs.Config
	validator  *validator.Validate
	handler    command.DeletePropertyHandler
	propRepo   property.Repository
	params     command.DeletePropertyCommand
	ServiceDep service.Dependencies
}

// SetupTest initializes the test suite.
func (s *DeletePropertyTestSuite) SetupSuite() {
	// Initialize the command handler
	s.handler = command.NewDeletePropertyHandler(
		s.ServiceDep.Repo.PropertyRepository,
		s.log,
		s.validator,
	)
	s.params = command.DeletePropertyCommand{
		PropertyID: database.NewStringID(),
		Server:     "Test",
	}

}

func (s *DeletePropertyTestSuite) SetupTest() {
	prop, err := s.ServiceDep.Repo.PropertyRepository.New(
		s.ctx,
		s.params.Server,
		property.NewPropertyParams{
			PropertyID: s.params.PropertyID,
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
		},
	)
	if err != nil {
		s.Fail("Failed to create property for testing", err)
	}
	s.log.Info("Property created for testing property:\n %+v", prop)
}

// TestCreatePropertyHandler tests the CreatePropertyHandler.
func (s *DeletePropertyTestSuite) TestDeletePropertyValid() {
	// Create a new property using the handler
	err := s.handler.Handle(s.ctx, s.params)
	s.NoError(err, "Expected no error when creating a property")
}
