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

func eventInit(db *gorm.DB, c echo.Context) model.Event {
  id      , _ := strconv.Atoi(c.QueryParam("id"))
  user_id , _ := strconv.Atoi(c.QueryParam("user_id"))
  year    , _ := strconv.Atoi(c.QueryParam("year"))
  month   , _ := strconv.Atoi(c.QueryParam("month"))
  day     , _ := strconv.Atoi(c.QueryParam("day"))
  return model.Event{id,user_id,c.QueryParam("summary"),c.QueryParam("dtstart"),c.QueryParam("dtend"),c.QueryParam("description"),year,month,day}
}

func  parseDateTime(data model.Event) model.Event {
  date := data.Dtstart
  fmt.Println(date)
  date = strings.Split(date, " ")[0]
  date_list := strings.Split(date, "-")

  data.Year , _ = strconv.Atoi(date_list[0])
  data.Month, _ = strconv.Atoi(date_list[1])
  data.Day  , _ = strconv.Atoi(date_list[2])
  return data
}

func EventRegist(db *gorm.DB) echo.HandlerFunc{
  return func(c echo.Context) error {
    event := eventInit(db, c)
    event = parseDateTime(event)
    db.Create(&event)
    return c.JSON(http.StatusOK, Res{Status:true})
  }
}

func EventDelete(db *gorm.DB) echo.HandlerFunc {
  return func(c echo.Context) error {
    event := eventInit(db, c)
    db.First(&event)
    db.Delete(&event)
    return c.JSON(http.StatusOK, Res{Status:true})
  }
}

func EventUpdate(db *gorm.DB) echo.HandlerFunc {
  return func(c echo.Context) error {
    event := eventInit(db, c)
    old_event := model.Event{Id:event.Id,User_Id:event.User_Id}
    event = parseDateTime(event)
    db.Model(&old_event).Update(&event)
    return c.JSON(http.StatusOK, Res{Status:true})
  }
}

func EventShowAll(db *gorm.DB) echo.HandlerFunc {
  return func(c echo.Context) error {
    event := eventInit(db, c)
    events := []model.Event{}
    db.Find(&events, "user_id = ? and year = ? and month = ?", event.User_Id, event.Year, event.Month)
    response := ResEvent{}
    response.Status = true
    day := JsonDayEvent{}
    for k := 0; k < 31; k = k+1 {
      response.Data = append(response.Data, day)
    }
    literal  := []string{}
    datetime := ""
    for _, v := range events{
      literal   = strings.Split(events[0].Dtstart, "T")
      datetime  = literal[0] + " " + literal[1][:len(literal[1])-1]
      v.Dtstart = datetime
      literal   = strings.Split(events[0].Dtend, "T")
      datetime  = literal[0] + " " + literal[1][:len(literal[1])-1]
      v.Dtend   = datetime
      response.Data[v.Day-1].Day = append(response.Data[v.Day-1].Day,v)
    }
    return c.JSON(http.StatusOK,response)
  }
}
