package redis

import (
	"context"
	"fmt"
	"github.com/RaymondCode/simple-demo/model"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
)

type IRedis interface {
	Exist(key string) bool
	HSetFavorite(key string, value int64) error
	FollowAdd(id int64, id2 int64) error
	FollowRemove(id int64, id2 int64) error
	FollowList(id int64, relationType string) ([]string, error)
	IsExist(id int64, id2 int64, keyPrefix string) bool
	FriendList(id int64) ([]string, error)
	ZAddMsm(key string, value string, score int64) error
	ZRangeByScore(key string, score string) ([]redis.Z, error)
	FavoriteFindByUser(id int64) ([]int64, error)
	FavoriteZAdd(id int64, list []model.Favorite) error
	VideoZrange(idList []int64) ([]model.Video, error)
}

type Redis struct {
	rdb *redis.Client
}

func (r Redis) VideoZrange(idList []int64) ([]model.Video, error) {
	var videoList []model.Video
	for _, id := range idList {
		key := fmt.Sprintf("video:%d", id)
		var video model.Video
		if err := r.rdb.HGetAll(context.Background(), key).Scan(&video); err != nil {
			return nil, err
		}
		videoList = append(videoList, video)
	}
	return videoList, nil
}

func (r Redis) FavoriteZAdd(id int64, list []model.Favorite) error {
	key := fmt.Sprintf("user:%d", id)
	var zList []redis.Z
	err := r.rdb.ZAdd(context.Background(), key, zList...).Err()
	return err
}

func (r Redis) FavoriteFindByUser(id int64) ([]int64, error) {
	var fav []int64
	key := fmt.Sprintf("user:%d", id)
	if err := r.rdb.ZRevRange(context.Background(), key, 0, -1).Err(); err != nil || len(fav) == 0 {
		return nil, err
	}
	return fav, nil
}

func (r Redis) ZRangeByScore(key string, score string) ([]redis.Z, error) {
	//TODO implement me
	max := strconv.Itoa(int(time.Now().Unix()))
	msgs, err := r.rdb.ZRangeByScoreWithScores(context.Background(), key, &redis.ZRangeBy{
		Min: score,
		Max: max,
	}).Result()
	return msgs, err
}

func (r Redis) ZAddMsm(key string, value string, score int64) error {
	//TODO implement me
	err := r.rdb.ZAdd(context.Background(), key, redis.Z{float64(score), value}).Err()
	return err
}

func (r Redis) FriendList(id int64) ([]string, error) {
	//TODO implement me
	key := "follow:" + strconv.FormatInt(id, 10)
	followIdList, err := r.rdb.SMembers(context.Background(), key).Result()
	if err != nil {
		return nil, err
	}
	var idList []string
	for _, i := range followIdList {
		// 可以查别人的follow，也可以查自己的follower
		fkey := "follow:" + i
		value := strconv.FormatInt(id, 10)
		isFriend := r.rdb.SIsMember(context.Background(), fkey, value).Val()
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
	isFollow := r.rdb.SIsMember(context.Background(), key, value).Val()
	return isFollow
}

func (r Redis) FollowList(id int64, relationType string) ([]string, error) {
	//TODO implement me
	key := relationType + strconv.FormatInt(id, 10)
	fmt.Println("fmt.Println redis", key)
	followIdList, err := r.rdb.SMembers(context.Background(), key).Result()
	//fmt.Println(followIdList)
	return followIdList, err
}

func (r Redis) FollowRemove(id int64, id2 int64) error {
	//TODO implement me
	key := "follow:" + strconv.FormatInt(id, 10)
	value := strconv.FormatInt(id2, 10)
	pipe := r.rdb.Pipeline()
	pipe.SRem(context.Background(), key, value)
	key = "follower:" + strconv.FormatInt(id2, 10)
	value = strconv.FormatInt(id, 10)
	pipe.SRem(context.Background(), key, value)
	_, err := pipe.Exec(context.Background())
	return err
}

func (r Redis) FollowAdd(id int64, id2 int64) error {
	//TODO implement me
	key := "follow:" + strconv.FormatInt(id, 10)
	value := strconv.FormatInt(id2, 10)
	_, err := r.rdb.Pipelined(context.Background(), func(pipe redis.Pipeliner) error {
		pipe.SAdd(context.Background(), key, value)
		key = "follower:" + strconv.FormatInt(id2, 10)
		value = strconv.FormatInt(id, 10)
		pipe.SAdd(context.Background(), key, value)
		return nil
	})
	return err
}

func (r Redis) HSetFavorite(key string, value int64) error {
	//TODO implement me
	return r.rdb.SAdd(context.Background(), key, value).Err()
}

func (r Redis) Exist(key string) bool {
	//TODO implement me
	return r.rdb.Exists(context.Background(), key).Val() >= 1
}

func NewRedis(rdb *redis.Client) IRedis {
	return Redis{rdb}
}
