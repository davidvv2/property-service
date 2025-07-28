//go:build cse
// +build cse

package query_test

import (
	"context"
	"property-service/internal/properties/app/query"
	"property-service/internal/properties/domain/owner"
	"property-service/internal/properties/service"
	"property-service/pkg/configs"
	"property-service/pkg/infrastructure/database"
	"property-service/pkg/infrastructure/log"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/suite"
)

// TestCommandTestSuite is the test suite for the command package.
type GetOwnerTestSuite struct {
	suite.Suite
	ctx        context.Context
	log        log.Logger
	config     configs.Config
	validator  *validator.Validate
	handler    query.GetOwnerHandler
	propRepo   owner.Repository
	params     query.GetOwnerQuery
	ServiceDep service.Dependencies
	newParams  owner.NewOwnerParams
}

func (s *GetOwnerTestSuite) SetupSuite() {
	// Initialize the command handler
	s.handler = query.NewGetOwnerHandler(
		s.ServiceDep.Repo.OwnerRepository,
		s.log,
		s.validator,
	)
	s.params = query.GetOwnerQuery{
		ID: database.NewStringID(),
	}

	s.newParams = owner.NewOwnerParams{
		ID:        s.params.ID,
		Name:      "John Doe",
		Email:     "john@test.com",
		Telephone: "1234567890",
	}
	// Create an owner for testing
	_, err := s.ServiceDep.Repo.OwnerRepository.New(
		s.ctx,
		s.newParams,
	)
	if err != nil {
		s.Fail("Failed to create owner for testing", err)
	}
}

func (s *GetOwnerTestSuite) TestGetOwnerHandler() {
	// Create a new owner using the handler
	owner, err := s.handler.Handle(s.ctx, s.params)
	s.NoError(err, "Expected no error when getting an owner")
	s.NotNil(owner, "Expected owner to be not nil")
	s.Equal(s.newParams.ID, owner.ID(), "Expected owner ID to match")
	s.Equal(s.newParams.Name, owner.Name(), "Expected owner name to match")
	s.Equal(s.newParams.Email, owner.Email(), "Expected owner email to match")
	s.Equal(s.newParams.Telephone, owner.Telephone(), "Expected owner telephone to match")
}

func (s *GetOwnerTestSuite) TearDownSuite() {
	// Clean up the test data
	if err := s.ServiceDep.Repo.OwnerRepository.Delete(s.ctx, s.params.ID); err != nil {
		s.log.Error("Failed to delete test owner: %v", err)
	}
}
