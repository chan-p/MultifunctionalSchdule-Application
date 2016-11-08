package main

import (
  //echoパッケージ
  "github.com/labstack/echo"
  "github.com/labstack/echo/engine/standard"
  "github.com/labstack/echo/middleware"
  //ファイルを分けたときパッケージとして別のディレクトリに入れてimportすることで使える
  //"./handler"
  //"./json"
)

func main(){
  //Echoインスタンス生成
  e := echo.New()
  //標準出力にログを出力
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())
  //ルーティング
  e.Get("/hello",MainPage())
  e.Get("/json",Response())
  //サーバー起動
  e.Run(standard.New(":1323"))
}
