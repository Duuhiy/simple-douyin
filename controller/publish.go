package controller

import (
	"bytes"
	"fmt"
	"github.com/RaymondCode/simple-demo/model"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"strconv"
)

type VideoListResponse struct {
	Response
	VideoList []model.VideoResp `json:"video_list"`
}

// Publish check token then save upload file to public directory
func (u *VideoController) Publish(c *gin.Context) {
	title := c.PostForm("title")
	data, err := c.FormFile("data")
	if err != nil {
		log.Println("读取上传的文件出错", err)
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	file, err := data.Open()
	if err != nil {
		log.Println("文件data.Open()出错", err)
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
	}
	buf := new(bytes.Buffer)
	_, _ = io.Copy(buf, file)
	username, _ := c.Get("username")
	password, _ := c.Get("password")
	finalName := fmt.Sprintf("%s_%s", username, title)

	err = u.videoService.Publish(finalName, username.(string), password.(string), buf.Bytes())
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}

// PublishList all users have same publish video list
func (u *VideoController) PublishList(c *gin.Context) {
	userIdStr := c.Query("user_id")
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: Response{
				StatusCode: 1,
			},
		})
	}
	videolist, err := u.videoService.PublishList(userId)
	if err != nil {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: Response{
				StatusCode: 1,
			},
			VideoList: videolist,
		})
	}
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: videolist,
	})
}
