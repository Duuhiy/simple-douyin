package middleware

import (
	"github.com/RaymondCode/simple-demo/controller"
	"github.com/RaymondCode/simple-demo/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Auth(c *gin.Context) {
	token := c.Query("token")
	mc, err := utils.ParseToken(token)
	if err != nil {
		log.Println("鉴权出错", err)
		c.JSON(http.StatusUnauthorized, controller.Response{1, "鉴权出错"})
	}
	c.Set("username", mc.Username)
	c.Set("password", mc.Password)
	c.Next()
}

func AuthPublish(c *gin.Context) {
	token := c.PostForm("token")
	log.Println(token)
	mc, err := utils.ParseToken(token)
	if err != nil {
		log.Println("鉴权出错", err)
		c.JSON(http.StatusUnauthorized, controller.Response{1, "鉴权出错"})
	}
	c.Set("username", mc.Username)
	c.Set("password", mc.Password)
	c.Next()
}
