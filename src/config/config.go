package config

import (
	"os"
	"loger"
	"encoding/json"
	"io/ioutil"
)
type Config struct {
	Database_IP   string `json:"dbip"`   //! 数据库IP
	Database_Port int    `json:"dbport"` //! 数据库端口

	LoginServer_IP   string `json:"loginserver_ip"`    //! 登录服务器IP
	LoginServer_Port int    `json:"loginserver_port"`  //! 登录服务器端口
	LoginServerID    int    `json:"loginserver_id"`    //! 登录服务器ID
	LoginServerLimit int    `json:"loginserver_limit"` //! 登陆服务器人数限制

	GameServer_IP   string `json:"gameserver_ip"`    //! 游戏服务器IP
	GameServer_Port int    `json:"gameserver_port"`  //! 游戏服务器端口
	GameServerID    int    `json:"gameserver_id"`    //! 游戏服务器ID
	GameServerLimit int    `json:"gameserver_limit"` //! 游戏服务器人数限制
	GameServerName  string `json:"gameserver_name"`  //! 游戏服务器名字
	GameServer_New bool `json:"gameserver_isnew"` //! 是否为新服
}

var Global_Config Config

//配置文件路径
const CONFIGPATH  ="./config/config.json"

//载入配置文件
func init() {
	f, err := os.Open(CONFIGPATH)
	if err != nil {
		if isExist := os.IsExist(err); isExist {
			loger.Fatal("Open config file fail. Error: %v", err.Error())
			return
		}
		//! 文件不存在则创建默认配置文件
		createDefaultConfigFile()
		f, _ = os.Open(CONFIGPATH)
	}
	defer f.Close()
	data,err:=ioutil.ReadAll(f)
	if err != nil {
		loger.Error("Read config file fail. Error: %v", err.Error())
		return
	}
	err=json.Unmarshal(data,&Global_Config)
	if err != nil {
		loger.Error("Unmarshal config file fail. Error: %v", err.Error())
		return
	}
	loger.Debug("Load config success: \r\n %s", string(data))
}

//! 创建默认配置文件
func createDefaultConfigFile() {
	f,err:=os.Create(CONFIGPATH)
	if err!=nil{
		loger.Fatal("Create default config file fail. Error: %v", err.Error())
	}
	c:=Config{
		Database_IP:      "127.0.0.1",
		Database_Port:    27017,
		LoginServer_IP:   "127.0.0.1",
		LoginServer_Port: 9999,
		LoginServerID:    1,
		LoginServerLimit: 5000,
	}
	j,err:=json.Marshal(c)
	if err!=nil{
		loger.Error("Marshal json fail. Error: %v", err.Error())
	}
	write_num,err:=f.Write(j)
	defer f.Close()
	if err != nil {
		loger.Error("Write default config file fail, write %d bytes. Error: %v", write_num, err.Error())
	}
}