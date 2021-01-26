package controller

import (
	"GOLANG-REACT-NATIVE/entity"
	"GOLANG-REACT-NATIVE/service"
	"GOLANG-REACT-NATIVE/usrmapper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserAPI struct {
	UserService service.UserService
}

func ProvideUserAPI(p service.UserService) UserAPI {
	return UserAPI{UserService: p}
}

func (p *UserAPI) CreateUser(c *gin.Context) {
	var toUser entity.User
	err := c.BindJSON(&toUser)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{"Users": p.UserService.CreateUser(usrmapper.ToUser(toUser))})
}

func (p *UserAPI) CreateIssues(c *gin.Context) {
	var toIssues entity.Isuues
	err := c.BindJSON(&toIssues)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{"issues": p.UserService.CreateIssue(usrmapper.ToIssues(toIssues))})
}

func (p *UserAPI) ShowUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"showUsers": p.UserService.ShowUsers()})
}

func (p *UserAPI) ShowIssues(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"showIssues": p.UserService.ShowIssues()})
}
