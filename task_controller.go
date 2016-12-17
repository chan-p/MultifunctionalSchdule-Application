package main

import (
  // フレームワーク関連パッケージ
  "github.com/labstack/echo"
  _ "github.com/labstack/echo/engine/standard"
  _ "github.com/labstack/echo/middleware"

  // ORM
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"

  "net/http"
  "fmt"
  "strconv"
  _ "time"
  "strings"
  _ "regexp"
  _ "reflect"

  "./model"
)

func taskInit(db *gorm.DB, c echo.Context) model.Task{
  id      , _ := strconv.Atoi(c.QueryParam("id"))
  user_id , _ := strconv.Atoi(c.QueryParam("user_id"))
  // dtend   , _ := time.Parse("2016-12-1", c.QueryParam("dtend"))
  year    , _ := strconv.Atoi(c.QueryParam("year"))
  month   , _ := strconv.Atoi(c.QueryParam("month"))
  day     , _ := strconv.Atoi(c.QueryParam("day"))
  return model.Task{id, user_id, c.QueryParam("title"), c.QueryParam("sub_task"), c.QueryParam("dtend"), year, month, day}
}

func  parseDate(data model.Task) model.Task {
  date := data.Dtend
  date_list := strings.Split(date, "-")

  data.Year , _ = strconv.Atoi(date_list[0])
  data.Month, _ = strconv.Atoi(date_list[1])
  data.Day  , _ = strconv.Atoi(date_list[2])
  return data
}

func TaskRegist(db *gorm.DB) echo.HandlerFunc{
  return func(c echo.Context) error {
    task := taskInit(db, c)
    task = parseDate(task)
    db.Create(&task)
    return c.JSON(http.StatusOK, Res{Status:true})
  }
}

//func TaskDelete(db *gorm.DB) {
//}

//func TaskUpdate(db *gorm.DB) {
//}

//func TaskShowMonth(db *gorm.DB) {
//}

func TaskShowAll(db *gorm.DB) echo.HandlerFunc{
  return func(c echo.Context) error {
    task := taskInit(db, c)
    tasks := []model.Task{}
    db.Find(&tasks, "user_id = ? and year = ? and month = ?", task.User_Id, task.Year, task.Month)
    response := Res{}
    response.Status = true
    day := JsonDay{}
    for k := 0; k < 31; k = k+1 {
      response.Data = append(response.Data, day)
    }
    literal := []string{}
    datetime := ""
    for _, v := range tasks{
      fmt.Println(v)
      literal   = strings.Split(v.Dtend, "T")
      datetime  = literal[0]
      v.Dtend   = datetime
      response.Data[v.Day-1].Day = append(response.Data[v.Day-1].Day,v)
    }
    return c.JSON(http.StatusOK, response)
  }
}

