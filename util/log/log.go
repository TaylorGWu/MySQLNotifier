package log

import (
	log "github.com/cihub/seelog"
	"MySQLNotifier/config"
)

var logger log.LoggerInterface

func New() {
	config.ConfigNotifier()
	globalConfig := config.Get()
	logger, err := log.LoggerFromConfigAsFile(globalConfig.LogConfigPath)

	if err != nil {
		panic("init seelog fail!")
	}

	// replace log recorder
	log.ReplaceLogger(logger)
}

func Del() {
	logger.Flush()
}