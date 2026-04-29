package user_test

import (
	"github.com/MamangRust/microservice-ecommerce-shared/domain/requests"
	"github.com/MamangRust/microservice-ecommerce-grpc-user/repository"
	db "github.com/MamangRust/microservice-ecommerce-pkg/database/schema"
	"github.com/MamangRust/microservice-ecommerce-shared/pb"
	"github.com/MamangRust/microservice-ecommerce-test"
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
)

type UserRepositoryTestSuite struct {
	tests.BaseTestSuite
	repo   *repository.Repositories
	userID int
}

func (s *UserRepositoryTestSuite) SetupSuite() {
	s.BaseTestSuite.SetupSuite()

	queries := db.New(s.DBPool())
	s.RunServiceMigrations("role")
	s.RunServiceMigrations("user")
	s.SetupRoleService()
	roleClient := pb.NewRoleQueryServiceClient(s.Conns["role"])
	s.repo = repository.NewRepositories(queries, roleClient)
}

func (s *UserRepositoryTestSuite) TearDownSuite() {
	s.BaseTestSuite.TearDownSuite()
}

func (s *UserRepositoryTestSuite) Test1_CreateUser() {
	req := &requests.CreateUserRequest{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
		Password:  "password123",
	}

	user, err := s.repo.UserCommand.Create(context.Background(), req)
	s.NoError(err)
	s.NotNil(user)
	s.Equal(req.FirstName, user.Firstname)
	s.Equal(req.Email, user.Email)
	s.userID = int(user.UserID)
}

func (s *UserRepositoryTestSuite) Test2_FindById() {
	s.Require().NotZero(s.userID)
	ctx := context.Background()

	found, err := s.repo.UserQuery.FindByID(ctx, s.userID)
	s.NoError(err)
	s.NotNil(found)
	s.Equal(s.userID, int(found.UserID))
}

func TestUserRepositorySuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(UserRepositoryTestSuite))
}
