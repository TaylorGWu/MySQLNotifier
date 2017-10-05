package BinLogExtractor

import (
	"MySQLNotifier/util/database/mysql"
	"MySQLNotifier/util/log"
	"MySQLNotifier/util/string_tool"
	"strconv"
	"time"
)

type BinLogExtractor struct {
	MySQLUtil *mysql.MySQLUtil
	CurrentBinLogFile string	// record current binlog_file' name
	CurrentBinLogPostion int	// record current binlog_file' position
}

var Extractor BinLogExtractor

func New() (err error){
	mysql.New()
	Extractor.MySQLUtil = mysql.Get()
	Extractor.CurrentBinLogFile, Extractor.CurrentBinLogPostion, err = Extractor.MySQLUtil.ShowMasterStatus()
	if err != nil {
		log.Get().Errorf("UpdateCurrentBinLogStatus fail:%s\n", err)
		return
	}
	return
}

func Get() (*BinLogExtractor){
	return &Extractor
}

func (extractor *BinLogExtractor) UpdateCurrentBinLogStatus() (err error){
	currentBinLogFile, _, err :=
		Extractor.MySQLUtil.ShowMasterStatus()

	// in case of after log flusing, new binlog file has been created
	if string_tool.CompareString(extractor.CurrentBinLogFile, currentBinLogFile) < 0 {
		extractor.CurrentBinLogFile = currentBinLogFile
		extractor.CurrentBinLogPostion = 0
	}

	records, length, err := extractor.MySQLUtil.GetLatelyBinLog(extractor.CurrentBinLogFile, extractor.CurrentBinLogPostion)
	// in case of no new binlog record
	if 0 != length {
		extractor.CurrentBinLogFile = records[length-1]["Log_name"]
		extractor.CurrentBinLogPostion, _ = strconv.Atoi(records[length-1]["End_log_pos"])

		/*
		fmt.Println("-----begin-----")
		fmt.Printf("%#v", extractor)
		for i := 0; i < length; i++ {
			fmt.Printf("%d:%s\n", i, records[i]["Info"])
		}
		fmt.Println("-----end-----")
		*/
	}
	return
}

func (extractor *BinLogExtractor) Extract() {
}

func (extractor *BinLogExtractor) Run() {
	err := New()
	if err != nil {
		log.Get().Errorf("extractor:Run fail:%s\n", err)
		panic("extractor:Run fail")
	}

	for {
		Extractor.UpdateCurrentBinLogStatus()
		time.Sleep(10*time.Second)
	}
}