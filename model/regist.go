package model

import (
  "net/http"
  "github.com/labstack/echo"
  "database/sql"
  _ "fmt"
  _ "github.com/go-sql-driver/mysql"
)

type event_data struct {
  id string
  summary string
  dtstart string
  dtend string
  description string
}

func inita(c echo.Context) event_data{
  return event_data{c.QueryParam("user_id"),c.QueryParam("summary"),c.QueryParam("dtstart"),c.QueryParam("dtend"),c.QueryParam("description")}
}

func (event event_data) regist_event(db *sql.DB) string{
  query := "insert into Event (user_id,summary,dtsart,dtend,description) values ('"+event.id+"','"+event.summary+"','"+event.dtstart+"','"+event.dtend+"','"+event.description+"')"
  _,err := db.Query(query)
  if err != nil {
    return "false"
  }
  return "true"
}

func Echo_regist(db *sql.DB) echo.HandlerFunc {
  return func(c echo.Context) error {
    event := inita(c)
    status := event.regist_event(db)
    return c.String(http.StatusOK,status)
  }
}
