package main

import (
	"encoding/json"
	"fmt"
	"github.com/RaymondCode/simple-demo/config"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	gmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"strings"
)

func main() {
	//go service.RunMessageServer()
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)
	defer logger.Sync()
	// viper读取nacos配置
	v := viper.New()
	v.SetConfigFile("config/config.yaml")
	if err := v.ReadInConfig(); err != nil {
		zap.L().Panic(err.Error())
	}
	var nacoscfg config.NacosConfig
	if err := v.Unmarshal(&nacoscfg); err != nil {
		zap.L().Panic(err.Error())
	}
	//fmt.Println("nacos", nacoscfg)

	// 读取nacos中的配置
	sc := []constant.ServerConfig{
		{
			IpAddr: nacoscfg.Host,
			Port:   nacoscfg.Port,
		},
	}

	cc := constant.ClientConfig{
		NamespaceId:         nacoscfg.Namespace, // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "tmp/nacos/log",
		CacheDir:            "tmp/nacos/cache",
		LogLevel:            "debug",
	}
	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": sc,
		"clientConfig":  cc,
	})
	if err != nil {
		zap.L().Panic(err.Error())
	}
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: nacoscfg.DataId,
		Group:  nacoscfg.Group,
	})
	if err != nil {
		zap.L().Panic(err.Error())
	}
	var cfg config.DouyinConfig
	if err = json.Unmarshal([]byte(content), &cfg); err != nil {
		zap.S().Info(err)
	}
	fmt.Println(cfg)

	r := gin.Default()

	dbcfg := cfg.Db
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbcfg.User, dbcfg.Password, dbcfg.Host, dbcfg.Port, dbcfg.Dbname)
	fmt.Println(dsn)
	//dsn := "root:123456@tcp(127.0.0.1:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(gmysql.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
		NameReplacer:  strings.NewReplacer("user", "User", "video", "Video", "favorite", "Favorite", "comment", "Comment", "message", "Message"),
	}})
	if err != nil {
		zap.L().Panic("连接 db 出错了")
	}

	rdbcfg := cfg.Rdb
	//fmt.Println(rdbcfg)
	rdb := redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", rdbcfg.Host, rdbcfg.Port),
		Password:     rdbcfg.Password, // no password set
		DB:           rdbcfg.DB,       // use default DB
		PoolSize:     rdbcfg.PoolSize,
		MinIdleConns: 50,
	})

	initRouter(r, db, rdb)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
