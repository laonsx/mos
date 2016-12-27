package routers

import (
	"controllers"
	"github.com/julienschmidt/httprouter"
)

var Routers=httprouter.New()

func init() {
	Routers.Handle("GET","/main",controllers.Hello)
}
