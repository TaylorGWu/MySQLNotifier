package mysql

import (
	"testing"
	"MySQLNotifier/util/log"
	"MySQLNotifier/config"
)

func TestMySQL(t *testing.T) {
	log.New()
	defer log.Get().Close()

	err := New()
	config.ConfigNotifier()
	defer Get().Close()

	if err != nil {
		t.Errorf("mysql init fail:%s", err)
	}

	Get().IsBinLogOpen()
}
