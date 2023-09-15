package main

import (
	"github.com/RaymondCode/simple-demo/controller"
	"github.com/RaymondCode/simple-demo/middleware"
	"github.com/RaymondCode/simple-demo/repository/mysql"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func initRouter(r *gin.Engine, db *gorm.DB) {
	// public directory is used to serve static resources
	r.Static("/static", "./public")

	ur := mysql.NewUUserRepository(db)
	us := service.NewUserService(ur)
	u := controller.NewUserController(us)

	vs := service.NewVideoService(ur)
	v := controller.NewVideoController(vs)
	// basic apis
	apiRouter := r.Group("/douyin")
	apiRouter.GET("/feed/", v.Feed)
	apiRouter.GET("/user/", middleware.Auth, u.UserInfo)
	apiRouter.POST("/user/register/", u.Register)
	apiRouter.POST("/user/login/", u.Login)
	apiRouter.POST("/publish/action/", middleware.AuthPublish, v.Publish)
	apiRouter.GET("/publish/list/", middleware.Auth, v.PublishList)

	// extra apis - I
	apiRouter.POST("/favorite/action/", controller.FavoriteAction)
	apiRouter.GET("/favorite/list/", controller.FavoriteList)
	apiRouter.POST("/comment/action/", controller.CommentAction)
	apiRouter.GET("/comment/list/", controller.CommentList)

	// extra apis - II
	apiRouter.POST("/relation/action/", controller.RelationAction)
	apiRouter.GET("/relation/follow/list/", controller.FollowList)
	apiRouter.GET("/relation/follower/list/", controller.FollowerList)
	apiRouter.GET("/relation/friend/list/", controller.FriendList)
	apiRouter.GET("/message/chat/", controller.MessageChat)
	apiRouter.POST("/message/action/", controller.MessageAction)
}
