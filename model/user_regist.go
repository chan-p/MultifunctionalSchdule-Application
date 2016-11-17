package model

import (
  //フレームワーク関係のパッケージ
  "net/http"
  "github.com/labstack/echo"

  //データベース関連のパッケージ
  "database/sql"
   _ "github.com/go-sql-driver/mysql"

  //標準パッケージ
  "fmt"
 )

//イベント情報の構造体
type user_status struct {
  email string
  pass string
}

//クエリからの取得情報でのイベント情報初期化
func initation_user(c echo.Context) user_status{
  return user_status{c.QueryParam("email"),c.QueryParam("pass")}
}

func (user user_status) regist_user(db *sql.DB) string{
  query := "insert into User (email,pass) values ('"+user.email+"','"+user.pass+"')"
  _,err := db.Query(query)
  if err != nil {
    fmt.Println(err)
    return "{'status':'false','data':{}}"
  }
  return "{'status':'true','data':{}}"
}

func Echo_user_regist(db *sql.DB) echo.HandlerFunc {
  return func(c echo.Context) error {
    user := initation_user(c)
    status := user.regist_user(db)
    return c.JSON(http.StatusOK,status)
  }
}
