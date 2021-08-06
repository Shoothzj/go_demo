package demo_beego_log

import (
	"github.com/astaxie/beego/logs"
	"testing"
)

func TestBeegoLog(t *testing.T) {
	logger := logs.NewLogger()
	logger.Emergency("this is emergency")
	logger.Alert("this is alert")
	logger.Critical("this is critical")
	logger.Error("this is error")
	logger.Warning("this is warning")
	logger.Notice("this is notice")
	logger.Informational("this is informational")
	logger.Debug("this is debug")
}
