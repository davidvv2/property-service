//go:build cse
// +build cse

package query_test

import (
	"context"
	"property-service/internal/properties/app/query"
	"property-service/internal/properties/domain/property"
	"property-service/internal/properties/service"
	"property-service/pkg/configs"
	"property-service/pkg/infrastructure/database"
	"property-service/pkg/infrastructure/log"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/suite"
)

// TestGetPropertyTestSuite is the test suite for the command package.
type GetPropertyTestSuite struct {
	suite.Suite
	ctx        context.Context
	log        log.Logger
	config     configs.Config
	validator  *validator.Validate
	handler    query.GetPropertyHandler
	propRepo   property.Repository
	params     query.GetPropertyQuery
	newParams  property.NewPropertyParams
	ServiceDep service.Dependencies
}

// SetupSuite initializes the test suite.
func (s *GetPropertyTestSuite) SetupSuite() {
	// Initialize the command handler
	s.handler = query.NewGetPropertyHandler(
		s.ServiceDep.Repo.PropertyRepository,
		s.log,
		s.validator,
	)
	s.params = query.GetPropertyQuery{
		ID:     database.NewStringID(),
		Server: "Test",
	}
	s.newParams = property.NewPropertyParams{
		PropertyID:    s.params.ID,
		OwnerID:       database.NewStringID(),
		Address:       "123 Main St",
		Description:   "A beautiful property",
		Title:         "Beautiful Property",
		Category:      "House",
		Available:     true,
		AvailableDate: time.Now(),
		SaleType:      1,
	}

	if _, err := s.ServiceDep.Repo.PropertyRepository.New(
		s.ctx,
		s.params.Server,
		s.newParams,
	); err != nil {
		s.Fail("Failed to create property for testing", err)
	}
}

// TestCreatePropertyHandler tests the CreatePropertyHandler.
func (s *GetPropertyTestSuite) TestCreatePropertyHandler() {
	// gets property using the handler
	property, err := s.handler.Handle(s.ctx, s.params)
	s.NoError(err, "Expected no error when finding the property")
	s.NotNil(property, "Expected property to be found")
	s.Equal(s.newParams.Title, property.Title, "Expected property title to match")
	s.Equal(s.newParams.Description, property.Description, "Expected property description to match")
}

func (s *GetPropertyTestSuite) TearDownSuite() {
	// Clean up the test data
	err := s.ServiceDep.Repo.PropertyRepository.Delete(s.ctx, s.params.Server, s.params.ID)
	if err != nil {
		s.log.Error("Failed to delete property after test", err)
	}
}
