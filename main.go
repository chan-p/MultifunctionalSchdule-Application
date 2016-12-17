package main

import (
  //フレームワーク関連パッケージ
  "github.com/labstack/echo"
  "github.com/labstack/echo/engine/standard"
  "github.com/labstack/echo/middleware"

  //データベース関連パッケージ
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"


  //ディレクトリ
  "./model"
)

type Res struct {
  Status bool `json:status`
  Data   []JsonDay
}

type JsonDay struct {
  Day    []model.Task
}

type ResEvent struct {
  Status bool `json:status`
  Data   []JsonDayEvent
}

type JsonDayEvent struct {
  Day    []model.Event
}

func db_connect() *sql.DB {
  //db,err := sql.Open("mysql","root:tomonori@tcp(localhost:3306)/social_app")
  db,err := sql.Open("mysql","root:tomonori@tcp(52.196.55.156:3306)/social_app?parseTime=true")

  if err != nil {
    panic(err.Error())
  }
  return db
}

func gorm_connect() *gorm.DB {
  //db,err := sql.Open("mysql","root:tomonori@tcp(localhost:3306)/social_app")
  db,err := gorm.Open("mysql","root:tomonori@tcp(52.196.55.156:3306)/social_app?parseTime=true")

  if err != nil {
    panic(err.Error())
  }
  return db
}

func main(){

  e := echo.New()
  db := db_connect()
  gorm := gorm_connect()
  defer db.Close()
  //ミドルウェアの使用機能
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  // ルーティング
  // イベントの表示
  e.Get("/calender"       ,EventShowAll(gorm))
  // イベントの登録
  e.Get("/calender/regist",EventRegist(gorm))
  // イベントの更新
  e.Get("/calender/update",EventUpdate(gorm))
  // イベントの削除
  e.Get("/calender/delete",EventDelete(gorm))
  // ユーザーの登録
  //e.Post("/user/regist",model.Echo_user_regist(db))
  // タスクの表示
  e.Get("/task"           ,TaskShowAll(gorm))
  // タスクの登録
  e.Get("/task/regist"    ,TaskRegist(gorm))
  // タスクの削除
  e.Get("/task/delete"    ,TaskDelete(gorm))
  // タスクの更新
  e.Get("/task/update"    ,TaskUpdate(gorm))
  //サーバー構築 ポート1323
  e.Run(standard.New(":1323"))
}
