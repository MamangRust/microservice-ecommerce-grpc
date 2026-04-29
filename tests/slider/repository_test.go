package slider_test

import (
	"context"
	"testing"

	"github.com/MamangRust/microservice-ecommerce-grpc-slider/repository"
	db "github.com/MamangRust/microservice-ecommerce-pkg/database/schema"
	"github.com/MamangRust/microservice-ecommerce-shared/domain/requests"
	"github.com/MamangRust/microservice-ecommerce-test"
	"github.com/stretchr/testify/suite"
)

type SliderRepositoryTestSuite struct {
	tests.BaseTestSuite
	repo *repository.Repositories
}

func (s *SliderRepositoryTestSuite) SetupSuite() {
	s.BaseTestSuite.SetupSuite()
	s.RunServiceMigrations("slider")
	queries := db.New(s.DBPool())
	s.repo = repository.NewRepositories(queries)
}

func (s *SliderRepositoryTestSuite) TestSliderLifecycle() {
	ctx := context.Background()

	req := &requests.CreateSliderRequest{
		Nama:     "Spring Collection",
		FilePath: "http://example.com/spring.jpg",
	}

	created, err := s.repo.SliderCommand.Create(ctx, req)
	s.NoError(err)
	s.NotNil(created)
	s.Equal(req.Nama, created.Name)

	found, err := s.repo.SliderQuery.FindByID(ctx, int(created.SliderID))
	s.NoError(err)
	s.Equal(created.Image, found.Image)
}

func TestSliderRepositorySuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(SliderRepositoryTestSuite))
}
