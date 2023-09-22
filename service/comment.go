package service

import (
	"errors"
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/repository/mysql"
	"github.com/RaymondCode/simple-demo/repository/redis"
	"log"
	"strconv"
)

type ICommentService interface {
	CommentAction(args *CommentActionArgs, username, password string) error
	CommentList(videoId int64) ([]model.CommentResp, error)
}

type CommentService struct {
	CommentRepository mysql.IUserRepository
	CommentRedis      redis.IRedis
}

func (c CommentService) CommentList(videoId int64) ([]model.CommentResp, error) {
	//TODO implement me
	commentList, err := c.CommentRepository.CommentFindByVideo(videoId)
	if err != nil {
		return nil, err
	}
	var commentListResp []model.CommentResp
	for _, comment := range commentList {
		user, _ := c.CommentRepository.FindOne(comment.UserId)
		commentResp := model.CommentResp{
			Id:         comment.Id,
			User:       *user,
			Content:    comment.Contents,
			CreateDate: comment.CreateAt.String(),
		}
		commentListResp = append(commentListResp, commentResp)
	}
	return commentListResp, nil
}

type CommentActionArgs struct {
	Token       string `json:"token" form:"token"`
	VideoId     string `json:"video_id" form:"video_id"`
	ActionType  string `json:"action_type" form:"action_type"`
	CommentId   string `json:"comment_id" form:"comment_id"`
	CommentText string `json:"comment_text" form:"comment_text"`
}

func (c CommentService) CommentAction(args *CommentActionArgs, username, password string) error {
	//TODO implement me
	videoId, err := strconv.ParseInt(args.VideoId, 10, 64)
	if err != nil {
		log.Println("videoId 格式错误")
		return err
	}
	actionType, err := strconv.ParseInt(args.ActionType, 10, 64)
	if err != nil {
		log.Println("actionType 格式错误")
		return err
	}
	switch actionType {
	case 1:
		// 发布评论
		userId, _ := c.CommentRepository.FindOneByToken(username, password)
		commentInstance := model.Comment{
			UserId:   userId,
			VideoId:  videoId,
			Contents: args.CommentText,
		}
		err = c.CommentRepository.CommentAdd(&commentInstance)
		log.Println("CommentAction service", err)
		return err
	case 2:
		// 删除评论
		commentId, _ := strconv.ParseInt(args.CommentId, 10, 64)
		err = c.CommentRepository.CommentRemove(commentId, videoId)
		return err
	default:
		return errors.New("actionType 格式错误")
	}
}

func NewCommentService(repository mysql.IUserRepository, rdb redis.IRedis) ICommentService {
	return &CommentService{repository, rdb}
}
