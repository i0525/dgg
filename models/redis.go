package models

import (
	"time"
	"github.com/gomodule/redigo/redis"
	"os"
	"strconv"
	"fmt"
	"encoding/json"
	"errors"
)


type Redis struct {
	pool     *redis.Pool
}

var Dredis *Redis


func InitRedis() {
	Dredis = new(Redis)
	RedisMax,_ :=strconv.Atoi(os.Getenv("RedisMax"))
	RedisMaxActive,_:=strconv.Atoi(os.Getenv("RedisMaxActive"))
	Dredis.pool = &redis.Pool{
		MaxIdle:     RedisMax,
		MaxActive:   RedisMaxActive,
		IdleTimeout: time.Duration(120),
		Dial: func() (redis.Conn, error) {
			RedisDB,_ :=strconv.Atoi(os.Getenv("RedisDB"))

			return redis.Dial(
				"tcp",
				os.Getenv("RedisAddress"),//"127.0.0.1:6379"
				redis.DialReadTimeout(time.Duration(1000)*time.Millisecond),
				redis.DialWriteTimeout(time.Duration(1000)*time.Millisecond),
				redis.DialConnectTimeout(time.Duration(1000)*time.Millisecond),
				redis.DialDatabase(RedisDB),
				//redis.DialUsername(os.Getenv("RedisUsername") ),
				//redis.DialPassword(os.Getenv("RedisPassword")),
			)
		},
	}
}


func Set(k, v string) {
	c := Dredis.pool.Get()
	defer c.Close()
	_, err := c.Do("SET", k, v)
	if err != nil {
		fmt.Println("set error", err.Error())
	}
}

func GetStringValue(k string) string {
	c := Dredis.pool.Get()
	defer c.Close()
	username, err := redis.String(c.Do("GET", k))
	if err != nil {
		fmt.Println("Get Error: ", err.Error())
		return ""
	}
	return username
}


func SetKeyExpire(k string, ex int) {
	c := Dredis.pool.Get()
	defer c.Close()
	_, err := c.Do("EXPIRE", k, ex)
	if err != nil {
		fmt.Println("set error", err.Error())
	}
}

func CheckKey(k string) bool {
	c := Dredis.pool.Get()
	defer c.Close()
	exist, err := redis.Bool(c.Do("EXISTS", k))
	if err != nil {
		fmt.Println(err)
		return false
	} else {
		return exist
	}
}

func DelKey(k string) error {
	c := Dredis.pool.Get()
	defer c.Close()
	_, err := c.Do("DEL", k)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func SetJson(k string, data interface{}) error {
	c := Dredis.pool.Get()
	defer c.Close()
	value, _ := json.Marshal(data)
	n, _ := c.Do("SETNX", k, value)
	if n != int64(1) {
		return errors.New("set failed")
	}
	return nil
}

func getJsonByte(key string) ([]byte, error) {
	c := Dredis.pool.Get()
	jsonGet, err := redis.Bytes(c.Do("GET", key))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return jsonGet, nil
}