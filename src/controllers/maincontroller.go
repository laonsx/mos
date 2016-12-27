package controllers

import (
	"net/http"
	"loger"
	"github.com/julienschmidt/httprouter"
)

func Hello(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	loger.Info("connection...")
	w.Write([]byte("hello world !"))
}