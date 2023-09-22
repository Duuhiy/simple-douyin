package service

import (
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/repository/mysql"
	"github.com/RaymondCode/simple-demo/utils"
	"log"
)

type IUserService interface {
	Register(username, password string) (int64, error)
	Login(username string, password string) (int64, error)
	User(id int64) (*model.UserResp, error)
}

type UserService struct {
	UserRepository mysql.IUserRepository
}

func (u *UserService) User(id int64) (*model.UserResp, error) {
	user, err := u.UserRepository.FindOne(id)
	if err != nil {
		return nil, err
	}
	userResp := model.UserResp{
		Id:              user.Id,
		Name:            user.Name,
		FollowCount:     user.FollowCount,
		FollowerCount:   user.FollowerCount,
		IsFollow:        false,
		Avatar:          user.Avatar,
		BackgroundImage: user.BackgroundImage,
		Signature:       user.Signature,
		TotalFavorited:  user.TotalFavorited,
		WorkCount:       user.WorkCount,
		FavoriteCount:   user.FavoriteCount,
	}
	return &userResp, nil
}

func (u *UserService) Login(username string, password string) (int64, error) {
	pwd := utils.PwdEncode(password)
	userId, err := u.UserRepository.FindOneByToken(username, pwd)
	if err != nil {
		return -1, err
	}
	//fmt.Println(user)
	return userId, nil
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
