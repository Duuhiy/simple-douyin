package controller

import (
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type IFavoriteController interface {
	FavoriteAction(c *gin.Context)
	FavoriteList(c *gin.Context)
}

type FavoriteController struct {
	favoriteService service.IFavoriteService
}

func NewFavoriteController(service service.IFavoriteService) IFavoriteController {
	return &FavoriteController{service}
}

// FavoriteAction no practical effect, just check if token is valid
func (f *FavoriteController) FavoriteAction(c *gin.Context) {
	videoIdStr := c.Query("video_id")
	actionTypeStr := c.Query("action_type")
	username, _ := c.Get("username")
	password, _ := c.Get("password")
	videoId, err := strconv.ParseInt(videoIdStr, 10, 64)
	if err != nil {
		log.Println("videoId 格式错误", err)
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "videoId 格式错误"})
	}
	actionType, err := strconv.ParseInt(actionTypeStr, 10, 64)
	if err != nil || (actionType != 1 && actionType != 2) {
		log.Println("actionType 格式错误", err)
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "actionType 格式错误"})
	}

	err = f.favoriteService.FavoriteAction(username.(string), password.(string), videoId, actionType)
	if err != nil {
		log.Println("FavoriteAction 出错", err)
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "FavoriteAction 出错"})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	}
}

// FavoriteList all users have same favorite video list
func (f *FavoriteController) FavoriteList(c *gin.Context) {
	//userIdStr := c.Query("user_id")
	//userId, err := strconv.ParseInt(userIdStr, 10, 64)
	//if err != nil {
	//	log.Println("userId 格式错误", err)
	//	c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "userId 格式错误"})
	//}
	usernameAny, _ := c.Get("username")
	passwordAny, _ := c.Get("password")
	username := usernameAny.(string)
	password := passwordAny.(string)
	videoList, err := f.favoriteService.FavoriteList(username, password)
	if err != nil {
		log.Println("f.favoriteService.FavoriteList 格式错误", err)
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "f.favoriteService.FavoriteList 格式错误"})
	}
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: videoList,
	})
}
