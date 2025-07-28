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

// TestListPropertiesByCategoryTestSuite is the test suite for the command package.
type ListPropertiesByCategoryTestSuite struct {
	suite.Suite
	ctx        context.Context
	log        log.Logger
	config     configs.Config
	validator  *validator.Validate
	handler    query.ListPropertiesByCategoryHandler
	propRepo   property.Repository
	params     query.ListPropertiesByCategoryQuery
	newParams  property.NewPropertyParams
	ServiceDep service.Dependencies
}

// SetupSuite initializes the test suite.
func (s *ListPropertiesByCategoryTestSuite) SetupSuite() {
	// Initialize the command handler
	s.handler = query.NewListPropertiesByCategoryHandler(
		s.ServiceDep.Repo.PropertyRepository,
		s.log,
		s.validator,
	)
	s.params = query.ListPropertiesByCategoryQuery{
		Category: "House",
		Sort:     1,
		Search:   1,
		Limit:    2,
	}
}

// TestCreatePropertyHandler tests the CreatePropertyHandler.
func (s *ListPropertiesByCategoryTestSuite) TestCreatePropertyHandler() {
	// gets property using the handler
	result, err := s.handler.Handle(s.ctx, s.params)
	s.NoError(err, "Expected no error when finding the property")
	s.log.Info("Result: ", result)
}

func (s *ListPropertiesByCategoryTestSuite) TearDownSuite() {
	// Clean up the test data
}
