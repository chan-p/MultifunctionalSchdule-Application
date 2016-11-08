package tool

import (
    "net/http"
    "os/exec"
    "time"
    "fmt"
    "github.com/labstack/echo"
)

func Res_cmd() echo.HandlerFunc {
    return func(c echo.Context) error {
        cmd := exec.Command("sleep", "5s")
        fmt.Println("sleep中中中: ", time.Now().Format("15:04:05"))
        fmt.Println("sleepppp終了: ", time.Now().Format("15:04:05"))
        cmd.Start()
        exec.Command(".././application").Start()
        cmd.Wait()
        return c.String(http.StatusOK, "Heoollo World")
    }
}
