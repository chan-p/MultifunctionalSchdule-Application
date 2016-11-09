package tool

import (
    "net/http"
    _ "os/exec"
    _ "time"
    _ "fmt"
    "github.com/labstack/echo"
    "github.com/codeskyblue/go-sh"
)

func Auto_pull() echo.HandlerFunc {
    return func(c echo.Context) error{
        sh.Command("git","pull").Run()
        //str := "pkill -f application"
        //err := exec.Command("sh","-c",str).Start()
        //fmt.Println(err)
        //sh.Command("pkill","-f","application")
        //fmt.Println("aa")
        return c.String(http.StatusOK,"hey")
    }
}
