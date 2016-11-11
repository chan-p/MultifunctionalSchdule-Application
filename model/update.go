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
type event_info struct {
  id string
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
func initationa(c echo.Context) event_info{
  return event_info{c.QueryParam("id"),c.QueryParam("user_id"),c.QueryParam("summary"),c.QueryParam("dtstart"),c.QueryParam("dtend"),c.QueryParam("description"),c.QueryParam("year"),c.QueryParam("month"),c.QueryParam("day")}
}

//日付データをyear,month,dayにパース
func  parse_time (event event_info) event_info{
  dtstart := event.dtstart
  date_time := strings.Split(dtstart," ")
  date := strings.Split(date_time[0],"-")
  event.year = date[0]
  event.month = date[1]
  event.day = date[2]
  fmt.Println(event)
  return event
}

//dbのイベント情報の更新
func (event event_info) update_event_to_db(db *sql.DB) string{
  event = parse_time(event)
 
  query := "update Event set summary='"+event.summary+"',dtstart='"+event.dtstart+"',dtend='"+event.dtend+"',description='"+event.description+"',year='"+event.year+"',month='"+event.month+"',day='"+event.day+"' where id="+event.id+" and user_id="+event.user_id
  _,err := db.Query(query)
  if err != nil {
    fmt.Println(err)
    return "{'status':'false'}"
  }
  return "{'status':'true'}"
}

func Echo_update(db *sql.DB) echo.HandlerFunc {
  return func(c echo.Context) error {
    event := initationa(c)
    status := event.update_event_to_db(db)
    return c.String(http.StatusOK,status)
  }
}
