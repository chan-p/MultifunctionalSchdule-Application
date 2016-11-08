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
        fmt.Println("sleep中: ", time.Now().Format("15:04:05"))
        fmt.Println("sleep終了: ", time.Now().Format("15:04:05"))
        exec.Command("GOOS=linux GOARCH=amd64 go build ../application.go").Start()
        cmd.Start()
        exec.Command(".././application").Start()
        cmd.Wait()
        return c.String(http.StatusOK, "Hello World")
    }
}
