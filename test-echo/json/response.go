package json

import(
  "net/http"
  "github.com/labstack/echo"
)

func Response() echo.HandlerFunc{
  return func (c echo.Context) error {
    //返したいjsonの形式
    jsonMap := map[string]string{
      "foo": "bar",
      "hoge": "fuga",
    }
    return c.JSON(http.StatusOK, jsonMap)
  }
}


