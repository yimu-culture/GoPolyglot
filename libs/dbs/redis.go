package dbs

import (
	"GoPolyglot/libs/configs"
	"fmt"
	"github.com/gin-gonic/gin"
	redis "github.com/go-redis/redis/v7"
	"time"
)

var GRedis map[string]*redis.Client

func init() {
	GRedis = make(map[string]*redis.Client)
}

func InitRedis() (err error) {
	for _, v := range configs.GConfig.Databases.Redis {
		opt := &redis.Options{
			Addr:         fmt.Sprintf("%s:%d", v.Address, v.Port),
			DB:           v.Db,
			Password:     v.Password,
			PoolSize:     v.PoolSize,
			MinIdleConns: v.MinIdleConns,
			DialTimeout:  time.Duration(v.DialTimeout) * time.Second,
		}

		if v.Username != "" {
			opt.Username = v.Username
		}

		client := redis.NewClient(opt)

		if _, err = client.Ping().Result(); err != nil {
			return err
		}

		GRedis[v.Asname] = client
	}

	return
}

func SetExInt(c *gin.Context, cacheName, key string, value int64, ts int64) (err error) {
	_, err = GRedis[cacheName].WithContext(c).Do("Set", key, value, "EX", ts).Result()
	return
}

func GetInt(c *gin.Context, cacheName, key string) (val int64, err error) {
	val, err = GRedis[cacheName].WithContext(c).Do("GET", key).Int64()
	return
}

func SetExStr(c *gin.Context, cacheName, key, value string, ts int64) (err error) {
	_, err = GRedis[cacheName].WithContext(c).Do("Set", key, value, "EX", ts).Result()
	return
}

func GetStr(c *gin.Context, cacheName, key string) (val string) {
	vall := GRedis[cacheName].WithContext(c).Do("GET", key).Val()
	if vall != nil {
		val = vall.(string)
	}
	return
}

func GetTtl(c *gin.Context, cacheName, key string) (ttl int64) {
	ttl = GRedis[cacheName].WithContext(c).TTL(key).Val().Milliseconds()
	return
}

func Del(c *gin.Context, cacheName, key string) (err error) {
	_, err = GRedis[cacheName].WithContext(c).Do("DEL", key).Result()
	return
}
