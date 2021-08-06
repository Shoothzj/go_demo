package service

import (
	"github.com/gogf/gf/net/ghttp"
	"net/http"
)

func Hello(r *ghttp.Request) {

}

func Delete(r *ghttp.Request) {
	r.Response.ClearBuffer()
	r.Response.WriteHeader(http.StatusNoContent)
	r.Exit()
}