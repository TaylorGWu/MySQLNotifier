package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"sync"
	log "MySQLNotifier/util/log"
)

type MySQLUtil struct {
	DB *sql.DB
	Error error
	InitOnce sync.Once
}

var dbUtil *MySQLUtil

func New() {
	dbUtil.InitOnce.Do(dbUtilInit)
}

func dbUtilInit() {
	dbUtil.DB, dbUtil.Error = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test?parseTime=true")
	if dbUtil.Error != nil {
		log.Get().Errorf("mysql open fail:%s", dbUtil.Error)
		return
	}
}

func (dbUtil *MySQLUtil) Close() {
	dbUtil.DB.Close()
}