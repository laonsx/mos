package loger

import (
	"log"
	"sync"
	"fmt"
	"time"
	"os"
)

const (
	_=iota
	LOGDEBUG//! 调试
	LOGINFO//! 信息
	LOGWARN//! 警告
	LOGERROR//! 错误
	LOGFATAL//! 致命错误
)

var loger *log.Logger //日志组件
var std *log.Logger	//终端日志组件
var logMinLevel int//最小日志等级
var terminalOutput bool//是否同时终端输出
var currentDay int//当前天
var logPath string//当前日志文件路径
var fileName string//日志文件名

var outputLock sync.Mutex

func Output(logType int,format string,v ...interface{}) {
	if logType<logMinLevel{
		return //小于级别自动过滤
	}
	outputLock.Lock()
	defer outputLock.Unlock()
	prefixStr:="unknow"
	switch logType { //! 根据消息类型自由搭配颜色
	case LOGDEBUG:
		prefixStr = "[Debug]"
	case LOGINFO:
		prefixStr = "[Info]"
	case LOGWARN:
		prefixStr = "[Warn]"
	case LOGERROR:
		prefixStr = "[Error]"
	case LOGFATAL:
		prefixStr = "[Fatal]"
	}
	s:=fmt.Sprintf(prefixStr+format,v...)
	changDay()//跨天生成新文件
	loger.Output(3,s)
	if true==terminalOutput{
		out:=fmt.Sprintf(format,v...)
		std.Output(3,out)
	}
}

func Print(format string, v ...interface{}) { //! Print
	fmt.Printf(format, v...)
}

func Debug(format string, v ...interface{}) { //! 调试
	Output(LOGDEBUG, format, v...)
}

func Info(format string, v ...interface{}) { //! 信息
	Output(LOGINFO, format, v...)
}

func Warn(format string, v ...interface{}) { //! 警告
	Output(LOGWARN, format, v...)
}

func Error(format string, v ...interface{}) { //! 错误
	Output(LOGERROR, format, v...)
}

func Fatal(format string, v ...interface{}) { //! 致命错误,使用会造成服务器退出,慎用!!
	Output(LOGFATAL, format, v...)
	os.Exit(1)
}

// 跨天生成新文件
func changDay() {
	now:=time.Now()
	nowDay:=now.Day()
	timeformat:=fileName+"-20060102.log"
	fileName:=logPath+"/"+now.Format(timeformat)
	_,err:=os.Stat(fileName)
	if nowDay==currentDay&&nil==err&&loger!=nil{
		return
	}
	currentDay=nowDay
	logfile, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0)
	if err!=nil{
		log.Println("open log file fail!",err.Error())
	}
	loger=log.New(logfile,"",log.Ltime|log.Lmicroseconds|log.Lshortfile)
}

//！ 创建日志管理器  日志路径  日志文件名 日志警示最小等级  终端是否同步输出
func InitLogger(Path string, Name string, MinLevel int, Output bool) {
	os.Mkdir(Path,7777)//! 创建log目录，如果存在则忽略
	logPath=Path
	logMinLevel=MinLevel
	terminalOutput=Output
	fileName=Name
	std=log.New(os.Stderr, "", log.Ltime|log.Lmicroseconds|log.Lshortfile)
}

func init() {
	InitLogger("./.logs/","app.log",1,true)
}