package service

import (
	"GOLANG-REACT-NATIVE/entity"
	"GOLANG-REACT-NATIVE/repositories"
)

type UserServive interface {
	Signup(entity.User) entity.User
	FindAll() []entity.User
}

type ProductService struct {
	ProductRepository repositories.ProductRepository
}

func ProvideProductService(p repositories.ProductRepository) ProductService {
	return ProductService{ProductRepository: p}
}

func (p *ProductService) CreateUser(user entity.User) entity.User {
	p.ProductRepository.CreateUser(user)
	return user
}

// type IssueServive interface {
// 	ShowIssues() []entity.Isuues
// }

// type userService struct {
// 	user []entity.User
// }

// type issuesService struct {
// 	issues []entity.Isuues
// }

// func New() UserServive {
// 	return &userService{
// 		user: []entity.User{},
// 	}
// }

// func (service *userService) Signup(user entity.User) entity.User {
// 	service.user = append(service.user, user)
// 	return user
// }

// func (service *userService) FindAll() []entity.User {
// 	return service.user
// }

// func (service *issuesService) ShowIssues() []entity.Isuues {
// 	return service.issues
// }
