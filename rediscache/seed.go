package rediscache

import (
	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
	"github.com/pkg/errors"
	"strconv"
	"hdchina/modeldef"
)

func GetSeedProperties(ch chan map[string][]string) {
	rc := RedisClient.Get()
	defer rc.Close()
	_, err := rc.Do("PING")
	if err != nil {
		beego.Error("unable to connect to redis")
		beego.AppConfig.Set("redis.enable", "false")
	}
	if redisenabled, _ := strconv.ParseBool(beego.AppConfig.String("redis.enable")); redisenabled {
		result := make(map[string][]string)
		for _, dbname := range DatabaseNames {
			value, _ := redis.Values(rc.Do("SMEMBERS", "dbname"))
			properties := make([]string, 0, 1)
			for _, v := range value {
				properties = append(properties, v.(string))
			}
			result[dbname] = properties
		}
		ch <- result
	} else {
		result := make(map[string][]string)
		ch <- result
	}
}



func ValidateProperties(ch chan error, params []string) {
	rc := RedisClient.Get()
	defer rc.Close()
	_, err := rc.Do("PING")
	if err != nil {
		beego.Error("unable to connect to redis")
		beego.AppConfig.Set("redis::enable", "false")
	}
	if redisenabled, _ := strconv.ParseBool(beego.AppConfig.String("redis::enable")); redisenabled {
		for index, dbname := range DatabaseNames {
			value, _ := redis.Values(rc.Do("SISMEMBER", dbname, params[index]))
			for _, v := range value {
				if v.(int) == 0 {
					ch <- errors.New(dbname)
					break
				}
			}
		}
		ch <- nil
	} else {
		ch <- errors.New("noredisconnect")
	}
}

func GetSeedById(ch <-chan []int, resultch chan [][]string) {
	rc := RedisClient.Get()
	defer rc.Close()
	_, err := rc.Do("PING")
	if err != nil {
		beego.Error("unable to connect to redis")
		beego.AppConfig.Set("redis.enable", "false")
	}
	ids := <-ch
	result := make([][]string, len(ids), 0)
	if redisenabled, _ := strconv.ParseBool(beego.AppConfig.String("redis::enable")); redisenabled {
		for _, id := range ids {
			value, _ := redis.Values(rc.Do("HMGET", string(id), SeedPageProperties))
			temp := make([]string, len(SeedPageProperties), 0)
			for _, v := range value {
				if v.(string) != "nil" {
					temp = append(temp, v.(string))
				} else {
					resultch <- nil
					return
				}
			}
			result = append(result ,temp)
		}
		resultch <- result
	}
	resultch <- nil
}

func SaveSeedById(seed modeldef.SeedPage) {
	rc := RedisClient.Get()
	defer rc.Close()
	_, err := rc.Do("PING")
	if err != nil {
		beego.Error("unable to connect to redis")
		beego.AppConfig.Set("redis.enable", "false")
	}
	if redisenabled, _ := strconv.ParseBool(beego.AppConfig.String("redis::enable")); redisenabled {
		rc.Do("HMSET", "title", seed.Title,
			"subtitle", seed.Subtitle, "IMDBPoint", seed.IMDBPoint,
			"freeSetting", seed.FreeSetting, "referRule", seed.ReferRule,
			"comments", seed.Comments, "datetime", seed.Datetime,
			"size", seed.Size, "upload", seed.Upload, "download", seed.Download,
			"completed", seed.Completed, "author", seed.Author)
	}
}