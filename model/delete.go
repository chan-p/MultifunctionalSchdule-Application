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
type event_id struct {
  id string
  user_id string
}

func initation_id(c echo.Context) event_id  {
  return event_id{c.QueryParam("id"),c.QueryParam("user_id")}
}

func (event event_id) delete_event_at_db(db *sql.DB) string {
  query := "DELETE FROM Event WHERE id="+event.id+" and user_id="+event.user_id
  _,err := db.Query(query)
  if err != nil {
    fmt.Println(err)
    return "{'status':'false','data':{}}"
  }
  return "{'status':'true','data':{}}"
}

func Echo_delete(db *sql.DB) echo.HandlerFunc {
  return func(c echo.Context) error {
    event := initation_id(c)
    status := event.delete_event_at_db(db)
    return c.JSON(http.StatusOK,status)
  }
}

