package main

import (
  "github.com/jinzhu/gorm"
  _ "github.com/go-sql-driver/mysql"
  "../model"
)

func main(){
  db, _  := gorm.Open("mysql","root:tomonori@tcp(52.196.55.156:3306)/social_app?parseTime=true")
  // db.CreateTable(&model.Tag{})
  db.CreateTable(&model.Task_Tag{})
}
