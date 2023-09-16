package redis

import "github.com/go-redis/redis"

type IRedis interface {
	Exist(key string) bool
	HSetFavorite(key string, value int64) error
}

type Redis struct {
	rdb *redis.Client
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
