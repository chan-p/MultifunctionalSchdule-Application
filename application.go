package main

import (
  //フレームワーク関連パッケージ
  "github.com/labstack/echo"
  "github.com/labstack/echo/engine/standard"
  "github.com/labstack/echo/middleware"

  //データベース関連パッケージ
  "database/sql"
  _ "github.com/go-sql-driver/mysql"

  //ディレクトリ
  "./tool"
  "./model"
)

func db_connect() *sql.DB {
  //db,err := sql.Open("mysql","root:tomonori@tcp(localhost:3306)/social_app")
  db,err := sql.Open("mysql","root:tomonori@tcp(52.196.55.156:3306)/social_app")

  if err != nil {
    panic(err.Error())
  }
  return db
}

func main(){

  e := echo.New()
  db := db_connect()
  defer db.Close()
  //ミドルウェアの使用機能
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  //モデル
  e.Get("/json",tool.Res_json())
   //登録済みイベントの表示
  e.Get("/calender",model.Echo_event(db))
   //イベントの登録
  e.Get("/calender/regist",model.Echo_event_regist(db))
   //イベントの更新
  e.Get("/calender/update",model.Echo_event_update(db))
   //イベントの削除
  e.Get("/calender/delete",model.Echo_event_delete(db))
   //ユーザーの登録
  e.Get("/user/regist",model.Echo_user_regist(db))
   //タスクの表示
  e.Get("/task",model.Echo_task(db))
   //タスクの登録
  e.Get("/task/regist",model.Echo_task_regist(db))

  e.Post("/pull",tool.Auto_pull())
  //サーバー構築 ポート1323
  e.Run(standard.New(":1323"))
}
