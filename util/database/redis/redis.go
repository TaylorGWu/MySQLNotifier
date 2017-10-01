package redis

import (
	"github.com/garyburd/redigo/redis"
	"MySQLNotifier/config"
	"MySQLNotifier/util/log"
	"fmt"
)

type RedisUtil struct {
	DB redis.Conn
}

var dbUtil RedisUtil

func New() (err error) {
	return dbUtilInit()
}

func Get() (*RedisUtil) {
	return &dbUtil
}

func dbUtilInit() (err error) {
	dbUtil.DB, err = redis.Dial("tcp", config.Get().Redis.Address)
	if err != nil {
		log.Get().Errorf("open redis connection fail: %s\n", err)
		return err
	}
	return
}

func (redisUtil *RedisUtil) Close() {
	redisUtil.DB.Close()
}

// in use of redis:publish command
func (redisUtil *RedisUtil) Publish(topic, message string) (err error) {
	_, err = redisUtil.DB.Do("publish", topic, message)
	if err != nil {
		log.Get().Errorf("redis publish message fail: %s\n", err)
		return err
	}
	return
}

// in use of redis:subscribe command
func (redisUtil *RedisUtil) Subcribe(topic string) (message string, err error) {
	var result interface{}
	result, err = redisUtil.DB.Do("subcribe", topic)
	if err != nil {
		log.Get().Errorf("redis subcribe message fail: %s\n", err)
	}
	fmt.Println(result)
	return "", nil
}