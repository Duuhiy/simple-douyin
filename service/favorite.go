package service

import (
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/repository/mysql"
	"github.com/RaymondCode/simple-demo/repository/redis"
	"strconv"
	"strings"
	"time"
)

type IFavoriteService interface {
	FavoriteAction(username, password string, videoId, actionType int64) error
	FavoriteList(username, password string) ([]model.VideoResp, error)
}

type FavoriteService struct {
	FavoriteRepository mysql.IUserRepository
	FavotiteRedis      redis.IRedis
}

func (f FavoriteService) FavoriteAction(username, password string, videoId, actionType int64) error {
	//TODO implement me
	// 1.视频点赞+1
	// 2.用户点赞+1
	// 3.加入favorite表中
	err := f.FavoriteRepository.FavoriteAction(username, password, videoId, actionType)
	if err != nil {
		return err
	}
	key := "favorite:" + username + ":" + strconv.FormatInt(videoId, 10)
	err = f.FavotiteRedis.HSetFavorite(key, time.Now().Unix())
	return err
}

func (f FavoriteService) FavoriteList(username, password string) ([]model.VideoResp, error) {
	//TODO implement me
	// 1. 根据userId查找favorite表
	userId, _ := f.FavoriteRepository.FindOneByToken(username)
	favoriteList, err := f.FavoriteRepository.FavoriteFindByUser(userId)
	if err != nil {
		return nil, err
	}
	if len(favoriteList) == 0 {
		return nil, nil
	}

	var videoListResp []model.VideoResp
	if len(favoriteList) == 1 {
		video, err := f.FavoriteRepository.VideoFindOneByVideo(favoriteList[0].VideoId)
		if err != nil {
			return nil, err
		}
		author, _ := f.FavoriteRepository.FindOne(video.Author)
		isFollow := f.FavotiteRedis.IsExist(userId, author.Id, "follow:")
		userResp := model.UserResp{
			Id:              author.Id,
			Name:            author.Name,
			FollowCount:     author.FollowCount,
			FollowerCount:   author.FollowerCount,
			IsFollow:        isFollow, // 从redis的follow表中查询
			Avatar:          author.Avatar,
			BackgroundImage: author.BackgroundImage,
			Signature:       author.Signature,
			TotalFavorited:  author.TotalFavorited,
			WorkCount:       author.WorkCount,
			FavoriteCount:   author.FavoriteCount,
		}
		videoResp := model.VideoResp{
			Id:            video.Id,
			Author:        userResp,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			Title:         video.Title,
			CommentCount:  video.CommentCount,
			FavoriteCount: video.FavoriteCount,
			IsFavorite:    true,
		}
		videoListResp = append(videoListResp, videoResp)
	} else {
		var s []string
		for _, favorite := range favoriteList {
			s = append(s, strconv.FormatInt(favorite.VideoId, 10))
		}
		idList := "(" + strings.Join(s, ",") + ")"
		// 2.查出喜欢的视频
		videoList, err := f.FavoriteRepository.VideoFindByIdList(idList)
		if err != nil {
			return nil, err
		}
		for _, video := range videoList {
			author, _ := f.FavoriteRepository.FindOne(video.Author)
			isFollow := f.FavotiteRedis.IsExist(userId, author.Id, "follow:")
			userResp := model.UserResp{
				Id:              author.Id,
				Name:            author.Name,
				FollowCount:     author.FollowCount,
				FollowerCount:   author.FollowerCount,
				IsFollow:        isFollow, // 从redis的follow表中查询
				Avatar:          author.Avatar,
				BackgroundImage: author.BackgroundImage,
				Signature:       author.Signature,
				TotalFavorited:  author.TotalFavorited,
				WorkCount:       author.WorkCount,
				FavoriteCount:   author.FavoriteCount,
			}
			videoResp := model.VideoResp{
				Id:            video.Id,
				Author:        userResp,
				PlayUrl:       video.PlayUrl,
				CoverUrl:      video.CoverUrl,
				Title:         video.Title,
				CommentCount:  video.CommentCount,
				FavoriteCount: video.FavoriteCount,
				IsFavorite:    true,
			}
			videoListResp = append(videoListResp, videoResp)
		}
	}
	return videoListResp, err
}

func NewFavoriteService(repository mysql.IUserRepository, rdb redis.IRedis) IFavoriteService {
	return &FavoriteService{repository, rdb}
}
