package service

import (
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/repository/mysql"
)

type IVideoService interface {
	Feed() ([]model.VideoResp, error)
	Publish(title, author, password string, data []byte) error
	PublishList(userId int64) ([]model.VideoResp, error)
}

type VideoService struct {
	videoRepository mysql.IUserRepository
}

func (v VideoService) Feed() ([]model.VideoResp, error) {
	//TODO implement me
	videoList, err := v.videoRepository.VideoFindAll()
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
			IsFavorite:    true,
		}
		videoListResp = append(videoListResp, videoResp)
	}
	return videoListResp, err
}

func (v VideoService) PublishList(userId int64) ([]model.VideoResp, error) {
	//TODO implement me
	videoList, err := v.videoRepository.FindAllByAuthor(userId)
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
			IsFavorite:    true,
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

func NewVideoService(video mysql.IUserRepository) IVideoService {
	return &VideoService{video}
}