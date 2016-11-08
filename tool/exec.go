package tool

import (
    "net/http"
    "os/exec"
    "time"
    "github.com/labstack/echo"
)

func cmd() echo.HandlerFunc {
    return func(c echo.Context) error {
        exec.Command("sleep", "5s").Start()
        fmt.Println("sleep中: ", time.Now().Format("15:04:05"))
        cmd.Wait()
        fmt.Println("sleep終了: ", time.Now().Format("15:04:05"))
        return c.String(http.StatusOK, "Hello World")
    }
}
