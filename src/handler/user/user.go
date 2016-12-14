package user

import (
		"fmt"
		"net/http"
		"github.com/labstack/echo"
		"github.com/tkc/GoogleAppEngineDeploy-Sandbox/src/db"
)

type (
		Handler struct {
		}
)

func (h Handler) Get(c echo.Context) error {
		connection := db.Connect()
		fmt.Print(connection)
		return c.String(http.StatusOK, "Hello, world")
}
