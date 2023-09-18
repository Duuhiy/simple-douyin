package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"strconv"
)

type IRedis interface {
	Exist(key string) bool
	HSetFavorite(key string, value int64) error
	FollowAdd(id int64, id2 int64) error
	FollowRemove(id int64, id2 int64) error
	FollowList(id int64, relationType string) ([]string, error)
	IsExist(id int64, id2 int64, keyPrefix string) bool
	FriendList(id int64) ([]string, error)
}

type Redis struct {
	rdb *redis.Client
}

func (r Redis) FriendList(id int64) ([]string, error) {
	//TODO implement me
	key := "follow:" + strconv.FormatInt(id, 10)
	followIdList, err := r.rdb.SMembers(key).Result()
	if err != nil {
		return nil, err
	}
	var idList []string
	for _, i := range followIdList {
		// 可以查别人的follow，也可以查自己的follower
		fkey := "follow:" + i
		value := strconv.FormatInt(id, 10)
		isFriend := r.rdb.SIsMember(fkey, value).Val()
		if isFriend {
			idList = append(idList, i)
		}
	}
	return idList, nil
}

func (r Redis) IsExist(id int64, id2 int64, keyPrefix string) bool {
	//TODO implement me
	key := keyPrefix + strconv.FormatInt(id, 10)
	//fmt.Println(key)
	value := strconv.FormatInt(id2, 10)
	isFollow := r.rdb.SIsMember(key, value).Val()
	return isFollow
}

func (r Redis) FollowList(id int64, relationType string) ([]string, error) {
	//TODO implement me
	key := relationType + strconv.FormatInt(id, 10)
	fmt.Println("fmt.Println redis", key)
	followIdList, err := r.rdb.SMembers(key).Result()
	//fmt.Println(followIdList)
	return followIdList, err
}

func (r Redis) FollowRemove(id int64, id2 int64) error {
	//TODO implement me
	key := "follow:" + strconv.FormatInt(id, 10)
	value := strconv.FormatInt(id2, 10)
	pipe := r.rdb.Pipeline()
	pipe.SRem(key, value)
	key = "follower:" + strconv.FormatInt(id2, 10)
	value = strconv.FormatInt(id, 10)
	pipe.SRem(key, value)
	_, err := pipe.Exec()
	return err
}

func (r Redis) FollowAdd(id int64, id2 int64) error {
	//TODO implement me
	key := "follow:" + strconv.FormatInt(id, 10)
	value := strconv.FormatInt(id2, 10)
	_, err := r.rdb.Pipelined(func(pipe redis.Pipeliner) error {
		pipe.SAdd(key, value)
		key = "follower:" + strconv.FormatInt(id2, 10)
		value = strconv.FormatInt(id, 10)
		pipe.SAdd(key, value)
		return nil
	})
	return err
}

func (r Redis) HSetFavorite(key string, value int64) error {
	//TODO implement me
	return r.rdb.SAdd(key, value).Err()
}

func (r Redis) Exist(key string) bool {
	//TODO implement me
	return r.rdb.Exists(key).Val() < 1
}

func NewRedis(rdb *redis.Client) IRedis {
	return Redis{rdb}
}
