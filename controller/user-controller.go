package controller

import (
	"GOLANG-REACT-NATIVE/entity"
	"GOLANG-REACT-NATIVE/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ToUser(userDTO entity.UserDTO) entity.User {
	return entity.User{
		ID:           userDTO.ID,
		USERNAME:     userDTO.USERNAME,
		MOBILENUMBER: userDTO.MOBILENUMBER,
	}
}

func ToUserDTO(userDTO entity.User) entity.UserDTO {
	return entity.UserDTO{
		ID:           userDTO.ID,
		USERNAME:     userDTO.USERNAME,
		MOBILENUMBER: userDTO.MOBILENUMBER,
	}
}

type ProductAPI struct {
	ProductService service.ProductService
}

func ProvideProductAPI(p service.ProductService) ProductAPI {
	return ProductAPI{ProductService: p}
}

func (p *ProductAPI) Create(c *gin.Context) {
	var userDTO entity.UserDTO
	err := c.BindJSON(&userDTO)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	createdProduct := p.ProductService.CreateUser(ToUser(userDTO))

	c.JSON(http.StatusOK, gin.H{"Users": ToUserDTO(createdProduct)})
}

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
