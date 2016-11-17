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
  _ "strings"
 )

//タスク情報の構造体
type task_data struct {
  user_id   string
  title     string
  sub_task  string
  year      string
  month     string
  day       string
}

//クエリからの取得情報でのイベント情報初期化
func initat(c echo.Context) task_data{
  return task_data{c.QueryParam("id"),c.QueryParam("title"),c.QueryParam("sub_task"),c.QueryParam("year"),c.QueryParam("month"),c.QueryParam("day")}
}
func (task task_data) regist_task(db *sql.DB) json {
  query := "insert into Task (user_id,title,sub_task,year,month,day) values ('"+task.user_id+"','"+task.title+"','"+task.sub_task+"','"+task.year+"','"+task.month+"','"+task.day+")"
  _,err := db.Query(query)

  if err != nil {
    fmt.Println(err)
    fal := json_task{0,"0","0",0,0,0}
    res := json{}
    res.Status = false
    res.Data = append(res.Data,fal)
    return res
  }
  fal := json_task{0,"0","0",0,0,0}
  res := json{}
  res.Status = true
  res.Data = append(res.Data,fal)
  return res
}

func Echo_task_regist(db *sql.DB) echo.HandlerFunc {
  return func(c echo.Context) error {
    task := initat(c)
    status := task.regist_task(db)
    return c.JSON(http.StatusOK,status)
  }
}
