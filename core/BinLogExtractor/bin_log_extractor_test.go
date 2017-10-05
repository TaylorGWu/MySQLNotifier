package BinLogExtractor

import (
	"testing"
	"MySQLNotifier/util/log"
)

func TestBinLogExtractor(t *testing.T) {
	log.New()
	New()
	Get().Run()
}
