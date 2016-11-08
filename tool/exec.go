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
        cmd := exec.Command("git pull")
        cmd.Start()
        fmt.Println("sasdadasdadaleep中中中: ", time.Now().Format("15:04:05"))
        cmd.Wait()
        fmt.Println("slpppp終了: ", time.Now().Format("15:04:05"))
        aa := exec.Command(".././application")
        aa.Start()
        aa.Wait()
        return c.String(http.StatusOK, "Heoollo World")
    }
}
