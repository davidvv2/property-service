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

// TestDeleteOwnerTestSuite is the test suite for the command package.
type DeleteOwnerTestSuite struct {
	suite.Suite
	ctx        context.Context
	log        log.Logger
	config     configs.Config
	validator  *validator.Validate
	handler    command.DeleteOwnerHandler
	propRepo   owner.Repository
	params     command.DeleteOwnerCommand
	ServiceDep service.Dependencies
}

// SetupTest initializes the test suite.
func (s *DeleteOwnerTestSuite) SetupSuite() {
	// Initialize the command handler
	s.handler = command.NewDeleteOwnerHandler(
		s.ServiceDep.Repo.OwnerRepository,
		s.log,
		s.validator,
	)
	s.params = command.DeleteOwnerCommand{
		OwnerID: database.NewStringID(),
	}

}

func (s *DeleteOwnerTestSuite) SetupTest() {
	owner, err := s.ServiceDep.Repo.OwnerRepository.New(
		s.ctx,
		owner.NewOwnerParams{
			ID:        s.params.OwnerID,
			Name:      "John Doe",
			Email:     "test@emails.com",
			Telephone: "1234567890",
		},
	)
	if err != nil {
		s.Fail("Failed to create owner for testing", err)
	}
	s.log.Info("Owner created for testing owner:\n %+v", owner)
}

// TestCreateOwnerHandler tests the CreateOwnerHandler.
func (s *DeleteOwnerTestSuite) TestDeleteOwnerValid() {
	// Create a new owner using the handler
	err := s.handler.Handle(s.ctx, s.params)
	s.NoError(err, "Expected no error when creating a owner")
}
