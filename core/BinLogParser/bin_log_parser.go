package BinLogParser

import (
	"MySQLNotifier/util/database/redis"
)

type BinLogParser struct {
	RedisUtil *redis.RedisUtil
}

var Parser BinLogParser

func New() (err error) {
	err = redis.New()
	if err != nil {
		panic("Init BinLogParser fail.\n")
		return
	}
	Parser.RedisUtil = redis.Get()
	return
}