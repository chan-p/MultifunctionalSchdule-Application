package model

type Task_Tag struct {
  Id      int    `json:id sql:AUTO_INCREMENT`
  Task_Id int `json:task_id`
  Tag_Id  int `json:tag_id`
}
