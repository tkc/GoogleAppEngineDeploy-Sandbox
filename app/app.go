package hellow

import (
		"net/http"
		"github.com/tkc/GoogleAppEngineDeploy-Sandbox/src/handler/user"
		"github.com/labstack/echo"
		"github.com/labstack/echo/engine/standard"
)

func init() {

		uh := user.Handler{}
		e := echo.New()
		e.GET("/", uh.Get)
		s := standard.New("")
		s.SetHandler(e)
		http.Handle("/", s)

}