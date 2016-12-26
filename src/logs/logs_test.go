package logs

import "testing"

func Test_initLogs_outPut_1(t *testing.T) {
	InitLogger("./","app.log",1,true)
	Print("...,,,,,...")
	Debug("你好 wo china")
	Info("............")
	Warn("你好 my china")
	Error(".............")
	Fatal("你好 china")
}