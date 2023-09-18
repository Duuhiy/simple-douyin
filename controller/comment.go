package controller

import (
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type ICommentController interface {
	CommentAction(c *gin.Context)
	CommentList(c *gin.Context)
}

type CommentController struct {
	commentService service.ICommentService
}

func NewCommentController(service service.ICommentService) ICommentController {
	return &CommentController{service}
}

type CommentListResponse struct {
	Response
	CommentList []model.CommentResp `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	Response
	Comment model.CommentResp `json:"comment,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func (co *CommentController) CommentAction(c *gin.Context) {
	var args service.CommentActionArgs
	err := c.ShouldBindQuery(&args)
	if err != nil {
		log.Println("args 格式错误", err)
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "args 格式错误"})
	}
	username, _ := c.Get("username")
	password, _ := c.Get("password")
	err = co.commentService.CommentAction(&args, username.(string), password.(string))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "CommentAction 错误"})
	} else {
		c.JSON(http.StatusOK, CommentActionResponse{Response: Response{StatusCode: 0}})
	}
}

// CommentList all videos have same demo comment list
func (co *CommentController) CommentList(c *gin.Context) {
	videoIdStr := c.Query("video_id")
	videoId, err := strconv.ParseInt(videoIdStr, 10, 64)
	if err != nil {
		log.Println("videoId 格式错误", err)
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "videoId 格式错误"})
	}
	commentList, err := co.commentService.CommentList(videoId)
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "CommentList 错误"})
	} else {
		c.JSON(http.StatusOK, CommentListResponse{
			Response:    Response{StatusCode: 0},
			CommentList: commentList,
		})
	}
}
