package utils

import (
	"fmt"
	"github.com/chujieyang/redis-clean/conf"
	"github.com/gomodule/redigo/redis"
	"time"
)

var redisPool8 *redis.Pool
var redisPool50 *redis.Pool
var db8Client redis.Conn
var db50Client redis.Conn

func init() {
	initRedisPool()
}

func initRedisPool() {
	redisPool8 = &redis.Pool{
		MaxIdle:     256,
		MaxActive:   3,  // 线程池大小
		IdleTimeout: time.Duration(120),
		Wait: true,
		Dial: func() (redis.Conn, error) {
			return redis.Dial(
				"tcp",
				conf.RedisHost,
				redis.DialReadTimeout(time.Duration(1000)*time.Millisecond),
				redis.DialWriteTimeout(time.Duration(1000)*time.Millisecond),
				redis.DialConnectTimeout(time.Duration(1000)*time.Millisecond),
				redis.DialDatabase(8),
				redis.DialPassword(conf.RedisAuth),
			)
		},
	}
	redisPool50 = &redis.Pool{
		MaxIdle:     256,
		MaxActive:   3,  // 线程池大小
		IdleTimeout: time.Duration(120),
		Wait: true,
		Dial: func() (redis.Conn, error) {
			return redis.Dial(
				"tcp",
				conf.RedisHost,
				redis.DialReadTimeout(time.Duration(1000)*time.Millisecond),
				redis.DialWriteTimeout(time.Duration(1000)*time.Millisecond),
				redis.DialConnectTimeout(time.Duration(1000)*time.Millisecond),
				redis.DialDatabase(50),
				redis.DialPassword(conf.RedisAuth),
			)
		},
	}
	db8Client = redisPool8.Get()
	db50Client = redisPool50.Get()
}

func redisCmdExec(db int, cmd string, args ...interface{}) (interface{}, error) {
	//con := redisPool8.Get()
	//if db == 50 {
	//	con = redisPool50.Get()
	//}
	//if err := con.Err(); err != nil {
	//	return nil, err
	//}
	//defer func() {
	//	if err := con.Close(); err != nil {
	//		fmt.Println(err)
	//	}
	//}()
	//return con.Do(cmd, args...)
	if db == 8 {
		return db8Client.Do(cmd, args...)
	} else {
		return db50Client.Do(cmd, args...)
	}
}

func RemoveRedisKeys(db int, keys []interface{}, count *int) (err error) {
	_, err = redisCmdExec(db, "del", keys...)
	if err != nil {
		fmt.Println(err)
		return
	}
	*count += len(keys)
	return
}