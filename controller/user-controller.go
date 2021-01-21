package controller

import (
	"GOLANG-REACT-NATIVE/entity"
	"GOLANG-REACT-NATIVE/service"
	"GOLANG-REACT-NATIVE/usrmapper"
	"fmt"
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

	// createdUser := p.ProductService.CreateUser(usrmapper.ToUser(userDTO))
	// usrmapper.ToUser(createdUser)
	c.JSON(http.StatusOK, gin.H{"Users": p.UserService.CreateUser(usrmapper.ToUser(toUser))})
}

func (p *UserAPI) CreateIssues(c *gin.Context) {
	var toIssues entity.Isuues
	fmt.Println(toIssues)
	err := c.BindJSON(&toIssues)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	// createdUser := p.ProductService.CreateUser(usrmapper.ToUser(userDTO))
	// usrmapper.ToUser(createdUser)
	c.JSON(http.StatusOK, gin.H{"issues": p.UserService.CreateIssue(usrmapper.ToIssues(toIssues))})
}

func (p *UserAPI) ShowUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"showUsers": p.UserService.ShowUsers()})
}

func (p *UserAPI) ShowIssues(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"showIssues": p.UserService.ShowIssues()})
}

// type UserController interface {
// 	Signup(ctx *gin.Context) entity.User
// 	FindAll() []entity.User
// }

// type controller struct {
// 	service service.UserServive
// }

// func New(service service.UserServive) UserController {
// 	return &controller{
// 		service: service,
// 	}
// }

// func (c *controller) FindAll() []entity.User {

// 	return c.service.FindAll()
// }

// func (c *controller) Signup(ctx *gin.Context) entity.User {
// 	var user entity.User
// 	ctx.BindJSON(&user)
// 	c.service.Signup(user)
// 	return user
// }
