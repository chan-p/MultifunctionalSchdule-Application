package tool

import (
    "net/http"
    _ "fmt"
    "github.com/labstack/echo"
    "github.com/codeskyblue/go-sh"
)

func Auto_pull() echo.HandlerFunc {
    return func(c echo.Context) error{
        sh.Command("git","pull").Run()
        return c.String(http.StatusOK,"")
    }
}
