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

var redisUtil RedisUtil

func New() (error error) {
	return dbUtilInit()
}

func dbUtilInit() (error error) {
	redisUtil.DB, error = redis.Dial("tcp", config.Get().Redis.Address)
	if error != nil {
		log.Get().Errorf("open redis connection fail: %s\n", error)
		return error
	}
	return
}

func (redisUtil *RedisUtil) Del() {
	redisUtil.DB.Close()
}

// in use of redis:publish command
func (redisUtil *RedisUtil) Publish(topic, message string) (error error) {
	_, error = redisUtil.DB.Do("publish", topic, message)
	if error != nil {
		log.Get().Errorf("redis publish message fail: %s\n", error)
		return error
	}
	return
}

// in use of redis:subscribe command
func (redisUtil *RedisUtil) Subcribe(topic string) (message string, error error) {
	var result interface{}
	result, error = redisUtil.DB.Do("subcribe", topic)
	if error != nil {
		log.Get().Errorf("redis subcribe message fail: %s\n", error)
	}
	fmt.Println(result)
}