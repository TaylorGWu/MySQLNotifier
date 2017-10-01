package core

import (
	"MySQLNotifier/util/database/redis"
)

/* query and process binlog inorder to publish update message */
type ModitifyPublisher struct {
	RedisUtil *redis.RedisUtil
}

var Publisher ModitifyPublisher


