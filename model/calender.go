package model

import (
  "net/http"
  "github.com/labstack/echo"
  "database/sql"
  "fmt"
  _ "github.com/go-sql-driver/mysql"
)
//ユーザーの構造体
type user_data struct {
  id string
  year string
  month string
}

//イベントの構造体(json形式)
type event struct {
  Id string `json:Id@`
  Summary string `json:"Summary"`
  Dtstart string `json:"dtstart"`
  Dtend string  `json:"dtend"`
  Description string `json:"dtstart"`
}

//クエリから情報取得
//ユーザー情報の初期化
func user_initation(c echo.Context) user_data{
  return user_data{c.QueryParam("id"),c.QueryParam("year"),c.QueryParam("month")}
}

//データベースからユーザーの登録したイベント情報を抽出
func (user user_data) extract_eventdata_from_db(db *sql.DB) []string {
  query := "select id,summary,dtstart,dtend,description from Event where user_id="+user.id+" and year="+user.year+" and month="+ user.month
  rows, err := db.Query(query)
  var value []string

  fmt.Println(err)
  if err != nil {
    value = append(value,"false")
    return value
  }
  colum,err := rows.Columns()

  values := make([]sql.RawBytes,len(colum))
  scanArgs := make([]interface{},len(values))

  for i := range values {
    scanArgs[i] = &values[i]
  }
  for rows.Next(){
    err = rows.Scan(scanArgs...)
    for _,col := range values {
      if col == nil {
        value = append(value,"false")
      } else {
        value = append(value,string(col))
      }
    }
  }
  return value
}

//ユーザーのイベント情報を返す
func (user user_data) get_event(db *sql.DB) string{
  data := user.extract_eventdata_from_db(db)
  if len(data) == 1 {
    return "false"
  }
  st := "{'status':'true','data':{\n"
  for i := 0;i < len(data);i = i + 4 {
    st += "[id:"+data[0+i]+",Summary:"+data[1+i]+",dtstart:"+data[2+i]+",dtend:"+data[2+i]+",description:"+data[4+i]+"]"
    st += "\n"
  }
  st += "}}"
  return st
}

func Echo_event(db *sql.DB) echo.HandlerFunc {
  return func(c echo.Context) error {
    //ユーザー情報を取得
    user := user_initation(c)
    //イベント情報を取得
    json := user.get_event(db)
    json = "{'status': 'true','data': {'1': [{'id': 1,'summary': '4分の3発表','dtstart': '2016-11-05 09:00:00','dtend': '2016-11-05 11:00:00','description': '場所：スニル'}],'2': [{'id': 1,'summary': '4分の3発表','dtstart': '2016-11-05 09:00:00','dtend': '2016-11-05 11:00:00','description': '場所：スニル'}, {'id': 3,'summary': '4分の3発表','dtstart': '2016-11-05 09:00:00','dtend': '2016-11-05 11:00:00','description': '場所：スニル'}]}}"
    return c.JSON(http.StatusOK,json)
  }
}
