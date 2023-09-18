package controller

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type IRelationController interface {
	RelationAction(c *gin.Context)
	FollowList(c *gin.Context)
	FollowerList(c *gin.Context)
	FriendList(c *gin.Context)
}

type RelationController struct {
	relation service.IRelationService
}

type UserListResponse struct {
	Response
	UserList []model.UserResp `json:"user_list"`
}

// RelationAction no practical effect, just check if token is valid
func (r *RelationController) RelationAction(c *gin.Context) {
	toUserIdStr := c.Query("to_user_id")
	actionTypeStr := c.Query("action_type")
	username, _ := c.Get("username")
	password, _ := c.Get("password")
	toUserId, err := strconv.ParseInt(toUserIdStr, 10, 64)
	if err != nil {
		log.Println("toUserId 格式错误")
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "toUserId 格式错误"})
	}
	actionType, err := strconv.ParseInt(actionTypeStr, 10, 64)
	if err != nil {
		log.Println("actionType 格式错误")
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "actionType 格式错误"})
	}
	err = r.relation.RelationAction(toUserId, actionType, username.(string), password.(string))
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "关注成功"})
	}
}

// FollowList all users have same follow list
func (r *RelationController) FollowList(c *gin.Context) {
	username, _ := c.Get("username")
	password, _ := c.Get("password")
	//fmt.Println(username, password)
	userList, err := r.relation.FollowList(username.(string), password.(string), "follow:")
	if err != nil {
		c.JSON(http.StatusOK, UserListResponse{
			Response: Response{
				StatusCode: 1,
			},
		})
	} else {
		c.JSON(http.StatusOK, UserListResponse{
			Response: Response{
				StatusCode: 0,
			},
			UserList: userList,
		})
	}
}

// FollowerList all users have same follower list
func (r *RelationController) FollowerList(c *gin.Context) {
	username, _ := c.Get("username")
	password, _ := c.Get("password")
	fmt.Println("FollowerList")
	userList, err := r.relation.FollowList(username.(string), password.(string), "follower:")
	if err != nil {
		c.JSON(http.StatusOK, UserListResponse{
			Response: Response{
				StatusCode: 1,
			},
		})
	} else {
		c.JSON(http.StatusOK, UserListResponse{
			Response: Response{
				StatusCode: 0,
			},
			UserList: userList,
		})
	}
}

// FriendList all users have same friend list
func (r *RelationController) FriendList(c *gin.Context) {
	username, _ := c.Get("username")
	password, _ := c.Get("password")
	userList, err := r.relation.FollowList(username.(string), password.(string), "friend")
	if err != nil {
		c.JSON(http.StatusOK, UserListResponse{
			Response: Response{
				StatusCode: 1,
			},
		})
	} else {
		c.JSON(http.StatusOK, UserListResponse{
			Response: Response{
				StatusCode: 0,
			},
			UserList: userList,
		})
	}
}

func NewRelationController(relation service.IRelationService) IRelationController {
	return &RelationController{relation}
}
