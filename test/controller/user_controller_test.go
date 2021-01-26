package controller

import (
	"GOLANG-REACT-NATIVE/controller"
	"GOLANG-REACT-NATIVE/entity"
	"GOLANG-REACT-NATIVE/service"
	"encoding/json"
	"net/http/httptest"
	"testing"

	unitTest "github.com/Valiben/gin_unit_test"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/mock"
)

type MockService struct {
	mock.Mock
}

type OrdinaryResponse struct {
	Errno  string `json:"errno"`
	Errmsg string `json:"errmsg"`
}

func TestCreateUser(t *testing.T) {
	testController := controller.ProvideUserAPI(service.UserService{})
	router := gin.Default()
	router.Static("/file", "./")
	router.POST("/Users", testController.CreateUser)
	unitTest.SetRouter(router)

	user := entity.User{
		ID:           "1",
		USERNAME:     "akshay",
		MOBILENUMBER: "735005254",
	}

	param := make(map[string]interface{})
	param["ID"] = user.ID
	param["USERNAME"] = user.USERNAME
	param["MOBILENUMBER"] = user.MOBILENUMBER

	resp := OrdinaryResponse{}

	unitTest.TestHandlerUnMarshalResp("Post", "/Users", "form", user, &resp)
}

func TestCreateIssue(t *testing.T) {
	testController := controller.ProvideUserAPI(service.UserService{})
	router := gin.Default()
	router.POST("/Users", testController.CreateIssues)
	unitTest.SetRouter(router)

	user := entity.User{
		ID:           "1",
		USERNAME:     "akshay",
		MOBILENUMBER: "735005254",
	}

	param := make(map[string]interface{})
	param["ID"] = user.ID
	param["USERNAME"] = user.USERNAME
	param["MOBILENUMBER"] = user.MOBILENUMBER

	resp := OrdinaryResponse{}

	err := unitTest.TestHandlerUnMarshalResp("Post", "/Issues", "form", user, &resp)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	testController.CreateIssues(c)

	var got gin.H
	err = json.Unmarshal(w.Body.Bytes(), &got)
	if err != nil {
		t.Fatal(err)
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
