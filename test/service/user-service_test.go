package service

import (
	"GOLANG-REACT-NATIVE/entity"
	"GOLANG-REACT-NATIVE/repositories"
	"GOLANG-REACT-NATIVE/service"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/mock"
)

type MockService struct {
	mock.Mock
}

func (mock *MockService) TestCreateUsers(user entity.User) entity.User {
	args := mock.Called()
	result := args.Get(0)
	return result.(entity.User)

}

func TestCreateUser(t *testing.T) {
	//mockRepo := new(repositories.UserRepository)
	mockRepo := new(MockService)
	user := entity.User{
		ID:           "1",
		USERNAME:     "akshay",
		MOBILENUMBER: "7350055254",
	}
	mockRepo.On("CreateUser").Return(user)

	test := service.ProvideUserService(repositories.UserRepository{})
	testService := service.ProvideUserService(test.UserRepository)
	testService.CreateUser(user)

	assert.Equal(t, "1", user.ID)
	assert.Equal(t, "akshay", user.USERNAME)
	assert.Equal(t, "7350055254", user.MOBILENUMBER)

}

func TestCreateIssue(t *testing.T) {
	//mockRepo := new(repositories.UserRepository)
	mockRepo := new(MockService)
	issue := entity.Isuues{
		ID:    "1",
		ISSUE: "problem in network",
	}
	mockRepo.On("CreateUser").Return(issue)

	test := service.ProvideUserService(repositories.UserRepository{})
	testService := service.ProvideUserService(test.UserRepository)
	testService.CreateIssue(issue)

	assert.Equal(t, "1", issue.ID)
	assert.Equal(t, "problem in network", issue.ISSUE)
}

func TestShowUsers(t *testing.T) {
	mockRepo := new(MockService)
	user := entity.User{
		ID:           "3",
		USERNAME:     "Akshay",
		MOBILENUMBER: "7350055233",
	}

	mockRepo.On("ShowUsers").Return([]entity.User{user})

	test := service.ProvideUserService(repositories.UserRepository{})
	testService := service.ProvideUserService(test.UserRepository)
	result := testService.ShowUsers()

	assert.Equal(t, "3", result[0].ID)
	assert.Equal(t, "Akshay", result[0].USERNAME)
	assert.Equal(t, "7350055233", result[0].MOBILENUMBER)
}

func TestShowIssues(t *testing.T) {
	mockRepo := new(MockService)
	issue := entity.Isuues{
		ID:    "1",
		ISSUE: "Problem in network",
	}

	mockRepo.On("ShowUsers").Return([]entity.Isuues{issue})

	test := service.ProvideUserService(repositories.UserRepository{})
	testService := service.ProvideUserService(test.UserRepository)
	result := testService.ShowIssues()

	assert.Equal(t, "1", result[0].ID)
	assert.Equal(t, "problem in network", result[0].ISSUE)
}
