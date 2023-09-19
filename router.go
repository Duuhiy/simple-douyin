package main

import (
	"github.com/RaymondCode/simple-demo/controller"
	"github.com/RaymondCode/simple-demo/middleware"
	"github.com/RaymondCode/simple-demo/repository/mysql"
	"github.com/RaymondCode/simple-demo/repository/redis"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	gredis "github.com/go-redis/redis"
	"gorm.io/gorm"
)

func initRouter(r *gin.Engine, db *gorm.DB, red *gredis.Client) {
	// public directory is used to serve static resources
	r.Static("/static", "./public")

	ur := mysql.NewUUserRepository(db)
	rdb := redis.NewRedis(red)

	us := service.NewUserService(ur)
	u := controller.NewUserController(us)

	vs := service.NewVideoService(ur, rdb)
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
	fs := service.NewFavoriteService(ur, rdb)
	f := controller.NewFavoriteController(fs)
	cs := service.NewCommentService(ur, rdb)
	co := controller.NewCommentController(cs)
	apiRouter.POST("/favorite/action/", middleware.Auth, f.FavoriteAction)
	apiRouter.GET("/favorite/list/", middleware.Auth, f.FavoriteList)
	apiRouter.POST("/comment/action/", middleware.Auth, co.CommentAction)
	apiRouter.GET("/comment/list/", middleware.Auth, co.CommentList)

	// extra apis - II
	rs := service.NewRelationService(ur, rdb)
	rc := controller.NewRelationController(rs)
	apiRouter.POST("/relation/action/", middleware.Auth, rc.RelationAction)
	apiRouter.GET("/relation/follow/list/", middleware.Auth, rc.FollowList)
	apiRouter.GET("/relation/follower/list/", middleware.Auth, rc.FollowerList)
	apiRouter.GET("/relation/friend/list/", middleware.Auth, rc.FriendList)

	ms := service.NewMessageService(ur, rdb)
	m := controller.NewRMessageController(ms)
	apiRouter.GET("/message/chat/", middleware.Auth, m.MessageChat)
	apiRouter.POST("/message/action/", middleware.Auth, m.MessageAction)
}
