package controller

import (
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/RaymondCode/simple-demo/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type IVideoController interface {
	Feed(c *gin.Context)
	Publish(c *gin.Context)
	PublishList(c *gin.Context)
}

type VideoController struct {
	videoService service.IVideoService
}

func NewVideoController(video service.IVideoService) IVideoController {
	return &VideoController{video}
}

type FeedResponse struct {
	Response
	VideoList []model.VideoResp `json:"video_list,omitempty"`
	NextTime  int64             `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func (u *VideoController) Feed(c *gin.Context) {
	token := c.Query("token")
	var username, password string
	if token != "" {
		mc, err := utils.ParseToken(token)
		if err != nil {
			log.Println("鉴权出错", err)
			c.JSON(http.StatusUnauthorized, Response{1, "鉴权出错"})
		}
		username, password = mc.Username, mc.Password
	}
	videlList, err := u.videoService.Feed(username, password)
	//fmt.Println(videlList)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, FeedResponse{
			Response:  Response{StatusCode: 1},
			VideoList: videlList,
			NextTime:  time.Now().Unix(),
		})
	}
	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: videlList,
		NextTime:  time.Now().Unix(),
	})
}
