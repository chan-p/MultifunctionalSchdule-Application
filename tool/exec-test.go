package tool

import (
    "net/http"
    _"os/exec"
    _ "time"
    _ "fmt"
    "github.com/labstack/echo"
    "github.com/codeskyblue/go-sh"
)

func Test_cmd() echo.HandlerFunc {
    return func(c echo.Context) error{
        sh.Command("git","status").Run()
        //cmd := exec.Command("git","log")
        //cmd.Start()
        //cmd.Wait()
        //a, _ := exec.Command("sh", "-c", "git log --oneline | wc -l").Output()
        //fmt.Println(string(out))
        return c.String(http.StatusOK,"hey")
    }
}
