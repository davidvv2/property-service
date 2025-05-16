//go:build cse
// +build cse

package query_test

import (
	"context"
	"property-service/internal/properties/app/query"
	"property-service/internal/properties/domain/property"
	"property-service/internal/properties/service"
	"property-service/pkg/configs"
	"property-service/pkg/infrastructure/log"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/suite"
)

// TestListPropertiesByCategoriesTestSuite is the test suite for the command package.
type ListPropertiesByCategoriesTestSuite struct {
	suite.Suite
	ctx        context.Context
	log        log.Logger
	config     configs.Config
	validator  *validator.Validate
	handler    query.ListPropertiesByCategoriesHandler
	propRepo   property.Repository
	params     query.ListPropertiesByCategoriesQuery
	newParams  property.NewPropertyParams
	ServiceDep service.Dependencies
}

// SetupSuite initializes the test suite.
func (s *ListPropertiesByCategoriesTestSuite) SetupSuite() {
	// Initialize the command handler
	s.handler = query.NewListPropertiesByCategoriesHandler(
		s.ServiceDep.Repo.PropertyRepository,
		s.log,
		s.validator,
	)
	s.params = query.ListPropertiesByCategoriesQuery{
		Category: "House",
		Sort:     1,
		Search:   1,
		Limit:    2,
		Server:   "Test",
	}
}

// TestCreatePropertyHandler tests the CreatePropertyHandler.
func (s *ListPropertiesByCategoriesTestSuite) TestCreatePropertyHandler() {
	// gets property using the handler
	result, err := s.handler.Handle(s.ctx, s.params)
	s.NoError(err, "Expected no error when finding the property")
	s.log.Info("Result: ", result)
}

func (s *ListPropertiesByCategoriesTestSuite) TearDownSuite() {
	// Clean up the test data
}
