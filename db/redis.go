/**
 * @Author: lixiang
 * @Date: 2025-03-18 20:57
 * @Description: redis.go
 */

package db

import (
	"AstraLend-backend/config"
	"AstraLend-backend/log"
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

var RedisConn *redis.Pool

func InitRedis() *redis.Pool {
	log.Logger.Info("Init Redis")
	redisConfig := config.Config.Redis

	RedisConn := &redis.Pool{
		MaxIdle:     redisConfig.MaxIdle,
		MaxActive:   redisConfig.MaxActive,
		Wait:        true,
		IdleTimeout: time.Duration(redisConfig.IdleTimeout) * time.Second,
		Dial: func() (redis.Conn, error) {
			//建立链接
			c, err := redis.Dial("tcp", fmt.Sprintf("%s:%s", redisConfig.Address, redisConfig.Port))
			if err != nil {
				return nil, err
			}
			//验证密码
			_, err = c.Do("auth", redisConfig.Password)
			if err != nil {
				panic("redis auth failed" + err.Error())
			}
			_, err = c.Do("select", redisConfig.Db)
			if err != nil {
				panic("redis select failed" + err.Error())
			}
			return c, nil
		},
	}
	err := RedisConn.Get().Err()
	if err != nil {
		panic("redis connect failed" + err.Error())
	}
	return RedisConn
}
func RedisSet(key string, data interface{}, aliveSecond int) error {
	conn := RedisConn.Get()
	defer func() {
		_ = conn.Close()
		log.Logger.Info("连接已关闭")
	}()
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	if aliveSecond > 0 {
		_, err = conn.Do("SETEX", key, aliveSecond, jsonData)
	} else {
		_, err = conn.Do("SET", key, jsonData)
	}
	if err != nil {
		return err
	}
	return nil
}
