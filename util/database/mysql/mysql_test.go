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
		t.Errorf("mysql init fail:%s\n", err)
	}

	_, err = Get().IsBinLogOpen()
	if err != nil {
		t.Errorf("mysql IsBinLogOpen fail:%s\n", err)
	}

	_, _, err = Get().ShowMasterStatus()
	if err != nil {
		t.Errorf("mysql ShowMasterStatus fail:%s\n", err)
	}
}
