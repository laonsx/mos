package logs

import "testing"

func Test_initLogs_outPut_1(t *testing.T) {
	//初始化
	InitLogger("./","app.log",1,true)
	//打印信息
	Print("...,,,,,...")
	Debug("你好 wo china")
	Info("............")
	Warn("你好 my china")
	Error(".............")
	Fatal("你好 china")
}