package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
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
		NameReplacer:  strings.NewReplacer("user", "User", "video", "Video", "favorite", "Favorite", "comment", "Comment"),
	}})
	if err != nil {
		log.Println("连接 db 出错了")
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("127.0.0.1:6379"),
		Password:     "", // no password set
		DB:           0,  // use default DB
		PoolSize:     100,
		MinIdleConns: 50,
	})

	initRouter(r, db, rdb)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
