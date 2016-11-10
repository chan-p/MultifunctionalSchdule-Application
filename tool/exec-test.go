package tool

import (
    "net/http"
    "os/exec"
    _ "fmt"
    "github.com/labstack/echo"
    "github.com/codeskyblue/go-sh"
)

func Auto_pull() echo.HandlerFunc {
    return func(c echo.Context) error{
        sh.Command("git","status").Run()
        sh.Command("sh","test.sh").Run()
        cmd := exec.Command("sh", "test.sh")
        cmd.Start()
        return c.String(http.StatusOK,"")
    }
}
