package BinLogExtractor

import (
	"MySQLNotifier/util/database/mysql"
	"MySQLNotifier/util/log"
)

type BinLogExtractor struct {
	MySQLUtil *mysql.MySQLUtil
	CurrentBinLogFile string	// record current binlog_file' name
	CurrentBinLogPostion int	// record current binlog_file' position
}

var Extractor BinLogExtractor

func New() {
	mysql.New()
	Extractor.MySQLUtil = mysql.Get()
}

func Get() (*BinLogExtractor){
	return &Extractor
}

func (extractor *BinLogExtractor) UpdateCurrentBinLogStatus() (err error){
	extractor.CurrentBinLogFile, extractor.CurrentBinLogPostion, err =
		Extractor.MySQLUtil.ShowMasterStatus()

	if err != nil {
		log.Get().Errorf("UpdateCurrentBinLogStatus fail:%s\n", err)
		return
	}
	return
}

func (extractor *BinLogExtractor) Extract() {
}