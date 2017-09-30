package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	//"MySQLNotifier/common/constant"
)

type NotifierConfig struct {
	LogConfigPath string `yaml:"log_config_path"`
	MySQL         MySQL  `yaml:"mysql"`
	Redis         Redis
}

var notifierConfig NotifierConfig

type MySQL struct {
	Address  string
	User     string
	Password string
	DBName   string `yaml:"dbname"`
}

type Redis struct {
	Address  string
	Password string
	DB       int `yaml:"db"`
}

type BinLogStatus struct {
	IsOpen bool	// is MySQL BinLog Open
	CurrentBinLogFile string // which binlog is recorded currently
	CurrentBinLogPosition int // current binlog writing position
}

var binLogStatus BinLogStatus

func GetBinLogStatus() (*BinLogStatus) {
	return &binLogStatus
}

/*
func (binLogStatus *BinLogStatus) SetIsOpen(status string) {
	if constant.BinLogOpenStatus == status {
		binLogStatus.IsOpen = true
	} else  if constant.BinLogCloseStatus == status {
		binLogStatus.IsOpen = false
	} else {
		log.Get().Errorf("")
	}
}*/

func ConfigNotifier() {
	yamlFile, err := ioutil.ReadFile("/home/taylor/code/Go/src/MySQLNotifier/config/notifier_config.yaml")
	if err != nil {
		panic("read notifier_config.yaml fail!")
	}

	err = yaml.Unmarshal(yamlFile, &notifierConfig)
	if err != nil {
		panic("unmarshal yamlfile into NotifierConfig fail!")
	}
}

func Get() NotifierConfig {
	return notifierConfig
}
