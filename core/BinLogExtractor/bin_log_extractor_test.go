package BinLogExtractor

import (
	"testing"
	"MySQLNotifier/util/log"
	"fmt"
)

func TestBinLogExtractor(t *testing.T) {
	log.New()
	New()
	Get().UpdateCurrentBinLogStatus()
	fmt.Printf("%#v\n", Get())
}
