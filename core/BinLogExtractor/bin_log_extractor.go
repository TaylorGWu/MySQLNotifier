package BinLogExtractor

import (
	"MySQLNotifier/util/database/mysql"
	"MySQLNotifier/util/log"
	"MySQLNotifier/util/string_tool"
	"MySQLNotifier/core/BinLogParser"
	"strconv"
	"time"
	"fmt"
	"MySQLNotifier/common"
	"encoding/json"
)

type BinLogExtractor struct {
	MySQLUtil *mysql.MySQLUtil
	CurrentBinLogFile string	// record current binlog_file' name
	CurrentBinLogPostion int	// record current binlog_file' position
	BinLogChannel chan string // read for BinLogParser
}

var Extractor BinLogExtractor

func New() (err error){
	mysql.New()
	Extractor.MySQLUtil = mysql.Get()
	Extractor.CurrentBinLogFile, Extractor.CurrentBinLogPostion, err = Extractor.MySQLUtil.ShowMasterStatus()
	Extractor.BinLogChannel = make(chan string, 1000)	// init one direction channel for write

	if err != nil {
		log.Get().Errorf("UpdateCurrentBinLogStatus fail:%s\n", err)
		return
	}
	return
}

func Get() (*BinLogExtractor){
	return &Extractor
}

func GetBinLogChannel() (*chan string){
	return &(Extractor.BinLogChannel)
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

		fmt.Println("-----begin-----")
		fmt.Printf("%#v", extractor)
		for i := 0; i < length; i++ {
			binLog := common.BinLog{
				EventType: records[i]["Event_type"],
				Info: records[i]["Info"],
			}

			binLogBytes, _ := json.Marshal(binLog)
			binLogStr := string(binLogBytes)
			Extractor.BinLogChannel <- binLogStr
		}
		fmt.Println("-----end-----")
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

	go func() {
		/*
		for {
			info := <-extractor.BinLogChannel
			fmt.Printf("read:%s\n", info)
		}*/
		BinLogParser.Parser.Run()
	}()

	for {
		Extractor.UpdateCurrentBinLogStatus()
		time.Sleep(10*time.Second)
	}
}