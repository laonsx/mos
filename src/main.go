package main

import (
	_ "config"
	"loger"
	"net/http"
	"routers"
	"tools"
	"db"
	"config"
)

/*type MyMux struct {

}

func (this MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		sayhelloName(w, r)
		return
	}
	if r.URL.Path == "/ha" {
		helloServer(w, r)
		return
	}

	http.NotFound(w, r)
	return
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello myroute!")
}

func helloServer(w http.ResponseWriter,r *http.Request) {
	auth:=r.Header.Get("Authorization")
	if auth==""{
		w.Header().Set("WWW-Authenticate", `Basic realm="Dotcoo User Login"`)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	loger.Debug(auth)
	auths:=strings.SplitN(auth,"",2)
	if len(auths)!=2{
		loger.Error("error....")
		return
	}
	authMethod:=auths[0]
	authB64:=auths[1]
	loger.Info(authMethod,authB64)
	switch authMethod {
	case "Basic":
		authstr,err:=base64.StdEncoding.DecodeString(authB64)
		if err!=nil{
			loger.Error(err.Error())
			io.WriteString(w,"unauthorized!\n")
			return
		}
		loger.Info(string(authstr))
		userPwd:=strings.SplitN(string(authstr),":",2)
		if len(userPwd)!=2{
			loger.Error(err.Error())
			return
		}
		username:=userPwd[0]
		password:=userPwd[1]
		loger.Info(username,password)
	default:
		loger.Error("error")
		return
	}
	io.WriteString(w,"hello world!\n")
}*/

func main() {
	//mux := &MyMux{}
	//http.ListenAndServe(":9090", mux)
	//db.Insert()
	conf:=config.Global_Config
	dbs,err:=db.New(conf.Database_Root,conf.Database_PW,conf.Database_IP,conf.Database_Port,conf.Database_Tb)
	if err==nil{
		i,err:=dbs.Table("userinfo").Key("username","created","age").Value("haha","2013-12-09","30" ).Insert()
		if err==nil{
			loger.Info("insert success ",i)
		}
	}
	dbs.Table("userinfo").Where("uid=59").Delete()
	dbs.Table("userinfo").Where("uid=60").Key("username","departname","age").Value("zys","zhongguo","8").Update()
	data:=dbs.Table("userinfo").Where("").OrderBy("age desc").FindAll()
	db.Print(data)
	loger.Info("start server on :8080")
	tools.StartConsole()
	http.ListenAndServe("127.0.0.1:8080",routers.Routers)

}