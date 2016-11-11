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
  "strings"
 )

//イベント情報の構造体
type event_data struct {
  user_id string
  summary string
  dtstart string
  dtend string
  description string
  year string
  month string
  day string
}

//クエリからの取得情報でのイベント情報初期化
func initation(c echo.Context) event_data{
  return event_data{c.QueryParam("user_id"),c.QueryParam("summary"),c.QueryParam("dtstart"),c.QueryParam("dtend"),c.QueryParam("description"),"0","0","0"}
}

//日付データをyear,month,dayにパース
func  parse_timedata(event event_data) event_data{
  dtstart := event.dtstart
  date_time := strings.Split(dtstart," ")
  date := strings.Split(date_time[0],"-")
  event.year = date[0]
  event.month = date[1]
  event.day = date[2]
  fmt.Println(event)
  return event
}

func (event event_data) regist_event(db *sql.DB) string{
  event = parse_timedata(event)
  query := "insert into Event (user_id,summary,dtstart,dtend,description,year,month,day) values ('"+event.user_id+"','"+event.summary+"','"+event.dtstart+"','"+event.dtend+"','"+event.description+"','"+event.year+"','"+event.month+"','"+event.day+"')"
  _,err := db.Query(query)
  if err != nil {
    fmt.Println(err)
    return "{'status':'false','data':{}}"
  }
  return "{'status':'true','data':{}}"
}

func Echo_regist(db *sql.DB) echo.HandlerFunc {
  return func(c echo.Context) error {
    event := initation(c)
    status := event.regist_event(db)
    return c.JSON(http.StatusOK,status)
  }
}
