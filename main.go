package main

import (
	"github.com/gin-gonic/gin"
	gmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"strings"
)

func main() {
	//go service.RunMessageServer()

	r := gin.Default()

	dsn := "root:123456@tcp(127.0.0.1:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(gmysql.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
		NameReplacer:  strings.NewReplacer("User", "user", "Video", "video"),
	}})
	if err != nil {
		log.Println("连接 db user 出错了")
	}

	initRouter(r, db)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
