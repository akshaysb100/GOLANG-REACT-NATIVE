package controller

import (
	"GOLANG-REACT-NATIVE/controller"
	"GOLANG-REACT-NATIVE/entity"
	"GOLANG-REACT-NATIVE/service"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/mock"
)

type MockService struct {
	mock.Mock
}

func TestCreateUser(t *testing.T) {
	mockRepo := new(MockService)
	testController := controller.ProvideUserAPI(service.UserService{})
	gin.SetMode(gin.TestMode)
	mockRepo.On("BindJSON").Return(entity.User{})

	w := httptest.NewRecorder()
	v := mockRepo.On("ToUser").Return(entity.Isuues{})

	mockRepo.On("CreateUsers").Return(v)

	r := gin.Default()
	r.POST("/Users", testController.CreateIssues)

	_, err := http.NewRequest(http.MethodPost, "/Users", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	var got gin.H
	err = json.Unmarshal(w.Body.Bytes(), &got)
	fmt.Println(w.Code)
	if w.Code != http.StatusOK {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}

	assert.Equal(t, 200, w.Code)
}

func TestCreateIssue(t *testing.T) {
	mockRepo := new(MockService)
	testController := controller.ProvideUserAPI(service.UserService{})
	gin.SetMode(gin.TestMode)
	mockRepo.On("BindJSON").Return(entity.User{})

	w := httptest.NewRecorder()
	v := mockRepo.On("ToIssues").Return(entity.Isuues{})

	mockRepo.On("CreateIssue").Return(v)

	r := gin.Default()
	r.POST("/issues", testController.CreateIssues)

	_, err := http.NewRequest(http.MethodPost, "/issues", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	var got gin.H
	err = json.Unmarshal(w.Body.Bytes(), &got)
	if w.Code != http.StatusOK {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}
	assert.Equal(t, 200, w.Code)
}

func TestShowUsers(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	testController := controller.ProvideUserAPI(service.UserService{})
	testController.ShowUsers(c)

	var got gin.H
	err := json.Unmarshal(w.Body.Bytes(), &got)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 200, w.Code)
}

func TestShowIssues(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	testController := controller.ProvideUserAPI(service.UserService{})
	testController.ShowIssues(c)

	var got gin.H
	err := json.Unmarshal(w.Body.Bytes(), &got)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 200, w.Code)
}
