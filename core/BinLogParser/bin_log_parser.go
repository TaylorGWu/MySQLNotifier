package BinLogParser

import (
	"MySQLNotifier/util/database/redis"
	"MySQLNotifier/util/log"
	"MySQLNotifier/core/BinLogExtractor"
	"encoding/json"
	"MySQLNotifier/common"
	"fmt"
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

func Get() (*BinLogParser) {
	return &Parser
}

func (parser *BinLogParser) Run() {
	err := New()
	if err != nil {
		log.Get().Errorf("parser:Run fail:%s\n", err)
		panic("parser:Run fail")
	}

	go func() {
		binLogChannel := BinLogExtractor.GetBinLogChannel()
		for {
			binLogStr := <-(*binLogChannel)
			var binLog common.BinLog
			// todo json unmarshal err
			json.Unmarshal([]byte(binLogStr), binLog)
			fmt.Printf("%#v", binLog)
		}
	}()
}