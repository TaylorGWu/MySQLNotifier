package redis

import (
	"testing"
	"MySQLNotifier/config"
)

func TestRedis(t *testing.T) {
	config.ConfigNotifier()
	err := New()
	defer Get().Close()
	if err != nil {
		t.Errorf("init redis fail:%s\n", err)
	}

	err = Get().Publish("topic1", "test")
	if err != nil {
		t.Errorf("redis publish fail:%s\n", err)
	}
}
