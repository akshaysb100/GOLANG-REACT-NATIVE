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

func (p *UserService) ShowUsers() []*entity.User {
	return p.UserRepository.ShowUsers()
}

func (p *UserService) ShowIssues() []*entity.Isuues {
	return p.UserRepository.ShowIssues()
}
