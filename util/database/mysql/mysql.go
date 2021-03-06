package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	log "MySQLNotifier/util/log"
	"fmt"
	"MySQLNotifier/config"
)

type MySQLUtil struct {
	DB *sql.DB
}

var dbUtil MySQLUtil

func New() (err error) {
	return dbUtilInit()
}

func dbUtilInit() (err error) {
	dbSource := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true",
		config.Get().MySQL.User,
		config.Get().MySQL.Password,
		config.Get().MySQL.Address,
		config.Get().MySQL.DBName,
	)

	dbUtil.DB, err = sql.Open("mysql", dbSource)
	if err != nil {
		log.Get().Errorf("mysql open fail:%s\n", err)
		return err
	}
	return
}

func Get() (*MySQLUtil) {
	return  &dbUtil
}

func (dbUtil *MySQLUtil) Close() {
	dbUtil.DB.Close()
}

func (dbUtil *MySQLUtil) IsBinLogOpen() (status string, err error) {
	sql := "show variables like 'log_bin'"
	rows, err := dbUtil.DB.Query(sql)

	defer rows.Close()
	if err != nil {
		log.Get().Errorf("mysql query IsBinLog:%s\n", err)
		return
	}

	var variableName, value string
	for rows.Next() {
		if err = rows.Scan(&variableName ,&value); err != nil {
			log.Get().Errorf("mysql query IsBinLog extract value fail:%s", err)
			return
		}
	}
	status = value
	return
}

func (dbUtil *MySQLUtil) ShowMasterStatus() (file string, position int , err error){
	sql := "show master status"
	rows, err := dbUtil.DB.Query(sql)

	defer  rows.Close()
	if err != nil {
		log.Get().Errorf("mysql query ShowMasterStatus fail:%s\n", err)
		return
	}

	var binlogDoDB, binlogIgnoreDB string
	for rows.Next() {
		if err = rows.Scan(&file, &position, &binlogDoDB, &binlogIgnoreDB); err != nil {
			log.Get().Errorf("mysql query ShowMasterStatus extract value fail:%s", err)
			return
		}
	}
	return
}

func (dnUtil *MySQLUtil) GetLatelyBinLog(binLogFile string, position int) (records [](map[string]string), length int, err error){
	sql := fmt.Sprintf("show binlog events in '%s' from %d", binLogFile, position)
	fmt.Println(sql)

	rows, err := dbUtil.DB.Query(sql)
	if err != nil {
		log.Get().Errorf("mysql query GetLatelyBinLog fail:%s\n", err)
		return
	}

	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))

	for i := range values {
		scanArgs[i] = &values[i]
	}

	records = make([](map[string]string), 100, 1000)
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			log.Get().Errorf("mysql query GetLatelyBinLog extract value fail:%s\n", err)
			continue
		}

		record := make(map[string]string)
		for i, col := range values {
			record[columns[i]] = string(col.([]byte))
		}
		records[length] = record
		length += 1
	}
	return
}