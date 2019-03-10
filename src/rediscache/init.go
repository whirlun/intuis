package rediscache

import (
	"github.com/astaxie/beego"
	"github.com/gomodule/redigo/redis"
	"time"
)

var (
	RedisClient   *redis.Pool
	DatabaseNames = []string{"category", "format", "production_group",
		"medium", "video_encode", "audio_encode", "refer_rule"}
	SeedPageProperties = []string{"id", "title", "subtitle", "IMDBPoint", "freeSetting",
		"referRule", "comments", "datetime", "size", "upload", "download", "completed",
		"author"}
)

func InitRedis(data map[string][]string) {
	RedisClient = &redis.Pool{
		MaxIdle:     beego.AppConfig.DefaultInt("redis::maxidle", 10),
		MaxActive:   beego.AppConfig.DefaultInt("redis::maxactive", 50),
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", beego.AppConfig.String("redis::addr"))
			if err != nil {
				beego.AppConfig.Set("redis::ready", "false")
				return nil, err
			} else {
				c.Do("SELECT", beego.AppConfig.DefaultInt("redis::db", 0))
				for dbname, members := range data {
					for _, m := range members {
						c.Do("SADD", dbname, m)
					}
				}
				return c, nil
			}
		},
	}
}

