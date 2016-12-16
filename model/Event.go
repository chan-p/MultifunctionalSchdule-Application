package model

// import "time"

//タスク情報の構造体
type Event struct {
  Id          int    `json:id`
  User_Id     int    `json:user_id`
  Summary     string `json:summary`
  Dtstart     string `json:dtstart`
  Dtend       string `json:dtend`
  Description string `json:description`
  Year        int    `json:year`
  Month       int    `json:month`
  Day         int    `json:day`
}

