//go:build cse
// +build cse

package command_test

import (
	"context"
	"property-service/internal/properties/app/command"
	"property-service/internal/properties/domain/owner"
	"property-service/internal/properties/service"
	"property-service/pkg/configs"
	"property-service/pkg/infrastructure/database"
	"property-service/pkg/infrastructure/log"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/suite"
)

// TestCommandTestSuite is the test suite for the command package.
type NewOwnerTestSuite struct {
	suite.Suite
	ctx        context.Context
	log        log.Logger
	config     configs.Config
	validator  *validator.Validate
	handler    command.CreateOwnerHandler
	propRepo   owner.Repository
	params     command.CreateOwnerCommand
	ServiceDep service.Dependencies
}

// SetupTest initializes the test suite.
func (s *NewOwnerTestSuite) SetupTest() {
	// Load env from file.
	s.config = configs.New()
	s.log = log.NewZapImpl(&s.config.Backend)
	s.validator = validator.New()
	s.ctx = context.Background()

	// Initialize the command handler
	s.handler = command.NewCreateOwnerHandler(
		s.ServiceDep.Repo.OwnerRepository,
		s.log,
		s.validator,
	)
	s.params = command.CreateOwnerCommand{
		OwnerID:   database.NewStringID(),
		Name:      "John Doe",
		Email:     "test@emails.com",
		Telephone: "1234567890",
		Server:    "Test",
	}
}

// TestCreateOwnerHandler tests the CreateOwnerHandler.
func (s *NewOwnerTestSuite) TestCreateOwnerHandler() {
	// Create a new owner using the handler
	err := s.handler.Handle(s.ctx, s.params)
	s.NoError(err, "Expected no error when creating a owner")

	// Verify that the owner was created successfully
	owner, err := s.ServiceDep.Repo.OwnerRepository.Get(s.ctx, s.params.Server, s.params.OwnerID)
	s.NoError(err, "Expected no error when finding the owner")
	s.NotNil(owner, "Expected owner to be found")
	s.Equal(s.params.Name, owner.Name(), "Expected owner name to match")
	s.Equal(s.params.Email, owner.Email(), "Expected owner email to match")
}

func (s *NewOwnerTestSuite) TearDownSuite() {
	// Clean up the test data
	if err := s.ServiceDep.Repo.OwnerRepository.Delete(s.ctx, s.params.Server, s.params.OwnerID); err != nil {
		s.log.Error("Failed to delete test owner: %v", err)
	}
}
