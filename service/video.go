package service

import (
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/repository/mysql"
	"github.com/RaymondCode/simple-demo/repository/redis"
	"log"
	"strconv"
)

type IVideoService interface {
	Feed(username string) ([]model.VideoResp, error)
	Publish(title, author, password string, data []byte) error
	PublishList(username, password string) ([]model.VideoResp, error)
}

type VideoService struct {
	videoRepository mysql.IUserRepository
	videoRedis      redis.IRedis
}

func (v VideoService) Feed(username string) ([]model.VideoResp, error) {
	//TODO implement me
	videoList, err := v.videoRepository.VideoFindAll()
	if err != nil {
		log.Println("v.videoRepository.VideoFindAll()", err)
		return nil, err
	}
	//fmt.Println("VideoService", videoList)
	var videoListResp []model.VideoResp
	for _, video := range videoList {
		author, _ := v.videoRepository.FindOne(video.Author)
		videoResp := model.VideoResp{
			Id:            video.Id,
			Author:        *author,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			Title:         video.Title,
			CommentCount:  video.CommentCount,
			FavoriteCount: video.FavoriteCount,
			IsFavorite:    false,
		}
		if username == "" {
			videoListResp = append(videoListResp, videoResp)
			continue
		}
		// 根据username查找redis中存不存在videoid
		key := "favorite:" + username + ":" + strconv.FormatInt(video.Id, 10)
		if v.videoRedis.Exist(key) {
			log.Println("videoResp.IsFavorite = true")
			videoResp.IsFavorite = true
		}
		videoListResp = append(videoListResp, videoResp)
	}
	return videoListResp, err
}

func (v VideoService) PublishList(username, password string) ([]model.VideoResp, error) {
	//TODO implement me
	user, _ := v.videoRepository.FindOneByToken(username, password)
	videoList, err := v.videoRepository.FindAllByAuthor(user.Id)
	//fmt.Println(user)
	var videoListResp []model.VideoResp
	for _, video := range videoList {
		author, _ := v.videoRepository.FindOne(video.Author)
		videoResp := model.VideoResp{
			Id:            video.Id,
			Author:        *author,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			Title:         video.Title,
			CommentCount:  video.CommentCount,
			FavoriteCount: video.FavoriteCount,
			//IsFavorite:    true,
		}
		videoListResp = append(videoListResp, videoResp)
	}
	return videoListResp, err
}

func (v VideoService) Publish(title, author, password string, data []byte) error {
	//TODO implement me
	// 修改作者work_count
	// 插入video表
	// 上传到oss
	user, _ := v.videoRepository.FindOneByToken(author, password)
	video := model.Video{
		Author: user.Id,
		Title:  title,
	}

	user.WorkCount++
	err := v.videoRepository.Publish(&video, user, data)
	return err
}

func NewVideoService(video mysql.IUserRepository, red redis.IRedis) IVideoService {
	return &VideoService{video, red}
}
