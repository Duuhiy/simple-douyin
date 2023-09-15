package service

import (
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/repository/mysql"
	"log"
)

type IUserService interface {
	Register(username, password string) (int64, error)
	Login(username string, password string) (int64, error)
	User(id int64) (*model.User, error)
}

type UserService struct {
	UserRepository mysql.IUserRepository
}

func (u *UserService) User(id int64) (*model.User, error) {
	user, err := u.UserRepository.FindOne(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserService) Login(username string, password string) (int64, error) {
	user, err := u.UserRepository.FindOneByToken(username, password)
	if err != nil {
		return -1, err
	}
	return user.Id, nil
}

func NewUserService(repository mysql.IUserRepository) IUserService {
	return &UserService{UserRepository: repository}
}

func (u *UserService) Register(username, password string) (int64, error) {
	// 2. 插入user表中
	user := &model.User{
		Name:     username,
		Password: password,
	}
	user, err := u.UserRepository.Insert(user)
	if err != nil {
		log.Println("插入数据库出错", err)
		return -1, err
	}
	return user.Id, nil
}
