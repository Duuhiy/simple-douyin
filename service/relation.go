package service

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/repository/mysql"
	"github.com/RaymondCode/simple-demo/repository/redis"
	"log"
	"strconv"
	"strings"
)

type IRelationService interface {
	RelationAction(toUserId, actionType int64, username, password string) error
	FollowList(username string, password, relationType string) ([]model.UserResp, error)
}

type RelationService struct {
	db  mysql.IUserRepository
	rdb redis.IRedis
}

func (r *RelationService) FollowList(username string, password, relationType string) ([]model.UserResp, error) {
	//TODO implement me
	user, _ := r.db.FindOneByToken(username, password)
	//fmt.Println(user)
	var userList []string
	var err error
	if relationType == "friend" {
		userList, err = r.rdb.FriendList(user.Id)
	} else {
		userList, err = r.rdb.FollowList(user.Id, relationType)
	}
	if err != nil {
		log.Println("FollowList service 出错了", err)
		return nil, err
	}
	fmt.Println(userList)
	if len(userList) == 0 {
		return nil, nil
	}
	var userListResp []model.UserResp
	if len(userList) == 1 {
		followedId, _ := strconv.ParseInt(userList[0], 10, 64)
		u, _ := r.db.FindOne(followedId)
		userResp := model.UserResp{
			Id:              u.Id,
			Name:            u.Name,
			FollowCount:     u.FollowCount,
			FollowerCount:   u.FollowerCount,
			IsFollow:        true,
			Avatar:          u.Avatar,
			BackgroundImage: u.BackgroundImage,
			Signature:       u.Signature,
			TotalFavorited:  u.TotalFavorited,
			WorkCount:       u.WorkCount,
			FavoriteCount:   u.FavoriteCount,
		}
		userListResp = append(userListResp, userResp)
	} else {
		idList := "(" + strings.Join(userList, ",") + ")"
		users, _ := r.db.UserFindByIdList(idList)
		for _, u := range users {
			userResp := model.UserResp{
				Id:              u.Id,
				Name:            u.Name,
				FollowCount:     u.FollowCount,
				FollowerCount:   u.FollowerCount,
				IsFollow:        true,
				Avatar:          u.Avatar,
				BackgroundImage: u.BackgroundImage,
				Signature:       u.Signature,
				TotalFavorited:  u.TotalFavorited,
				WorkCount:       u.WorkCount,
				FavoriteCount:   u.FavoriteCount,
			}
			userListResp = append(userListResp, userResp)
		}
	}
	return userListResp, nil
}

func (r *RelationService) RelationAction(toUserId, actionType int64, username, password string) error {
	//TODO implement me
	user, _ := r.db.FindOneByToken(username, password)
	toUser, _ := r.db.FindOne(toUserId)
	switch actionType {
	case 1:
		// 关注
		// 1.用户follow_count++
		// 2.toUserId follower_count++
		// 3.插入relation表中
		follow := model.Relation{
			UserId:   user.Id,
			ToUserId: toUserId,
		}
		err := r.db.RelationAdd(user, toUser, &follow)
		if err != nil {
			return err
		}
		// redis为每个用户维护两个集合，粉丝集合、关注者集合
		err = r.rdb.FollowAdd(user.Id, toUserId)
		if err != nil {
			return err
		}
	case 2:
		err := r.db.RelationRemove(user, toUser)
		if err != nil {
			return err
		}
		err = r.rdb.FollowRemove(user.Id, toUserId)
		if err != nil {
			return err
		}
	default:
	}
	return nil
}

func NewRelationService(db mysql.IUserRepository, rdb redis.IRedis) IRelationService {
	return &RelationService{db, rdb}
}
