package service

import (
	"GOLANG-REACT-NATIVE/entity"
	"GOLANG-REACT-NATIVE/repositories"
)

type UserServive interface {
	Signup(entity.User) entity.User
	FindAll() []entity.User
}

type UserService struct {
	UserRepository repositories.UserRepository
}

func ProvideUserService(p repositories.UserRepository) UserService {
	return UserService{UserRepository: p}
}

func (p *UserService) CreateUser(user entity.User) entity.User {
	return p.UserRepository.CreateUser(user)
}

func (p *UserService) CreateIssue(issues entity.Isuues) entity.Isuues {
	return p.UserRepository.CreateIssue(issues)
}

func (p *UserService) ShowUsers() int {
	return p.UserRepository.ShowUsers()
}

func (p *UserService) ShowIssues() int {
	return p.UserRepository.ShowIssues()
}

// func (p *UserService) ShowUsers() entity.Count {
// 	return p.UserRepository.ShowUsers()
// }

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
