package log

import (
	log "github.com/cihub/seelog"
)

var logger log.LoggerInterface

func New() {
	logger, err := log.LoggerFromConfigAsFile("path")

	if err != nil {
		panic("init seelog fail!")
	}

	// replace log recorder
	log.ReplaceLogger(logger)
}

func Del() {
	logger.Flush()
}