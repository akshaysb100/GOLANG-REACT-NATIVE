package controller

import (
	"GOLANG-REACT-NATIVE/entity"
	"GOLANG-REACT-NATIVE/service"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	Signup(ctx *gin.Context) entity.User
	FindAll() []entity.User
}

type controller struct {
	service service.UserServive
}

func New(service service.UserServive) UserController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.User {
    
	return c.service.FindAll()
}

func (c *controller) Signup(ctx *gin.Context) entity.User {
	var user entity.User
	ctx.BindJSON(&user)
	c.service.Signup(user)
	return user
}
