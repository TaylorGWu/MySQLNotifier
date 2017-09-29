package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
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

func ConfigNotifier() {
	yamlFile, err := ioutil.ReadFile("../../config/notifier_config.yaml")
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
