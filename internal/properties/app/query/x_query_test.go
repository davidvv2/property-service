//go:build cse
// +build cse

package query_test

import (
	"context"
	"testing"

	"property-service/internal/properties/service"
	"property-service/pkg/configs"
	"property-service/pkg/infrastructure/log"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/suite"
)

func TestQueryTestSuite(t *testing.T) {
	// Load env from file.
	envLoadingError := godotenv.Load("../../../../dev.env")
	if envLoadingError != nil {
		panic("can not load config " + envLoadingError.Error())
	}

	/////nolint: exhaustruct // Magic is done here.
	config := configs.New()
	log := log.NewZapImpl(&config.Backend)
	v := validator.New()
	s := service.BuildDependencies(config)
	// Initialize the test suite
	suite.Run(t, &GetOwnerTestSuite{
		log:        log,
		config:     config,
		validator:  v,
		ctx:        context.Background(),
		ServiceDep: s,
	})
	suite.Run(t, &GetPropertyTestSuite{
		log:        log,
		config:     config,
		validator:  v,
		ctx:        context.Background(),
		ServiceDep: s,
	})
	suite.Run(t, &ListPropertiesByCategoryTestSuite{
		log:        log,
		config:     config,
		validator:  v,
		ctx:        context.Background(),
		ServiceDep: s,
	})
}
