package model

// import "time"

//タスク情報の構造体
type Task struct {
  Id        int       `json:id`
  User_Id   int       `json:user_id`
  Title     string    `json:title`
  Sub_task  string    `json:sub_task`
  Dtend     string    `json:dtend`
  Status    int       `json:status`
  Year      int       `json:year`
  Month     int       `json:month`
  Day       int       `json:day`
}


